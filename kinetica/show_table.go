package kinetica

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"sync"

	"go.opentelemetry.io/otel/trace"
)

type ShowTableDistributionType string

const (
	Replicated       ShowTableDistributionType = "Replicated"
	Sharded                                    = "Sharded"
	ShardedRandomly                            = "Random"
	ExternalResource                           = "External Resource"
	LogicalView                                = "Logical View"
)

type ShowTableChar struct {
	TypeName string
	Length   int
}

var CharTypes = []ShowTableChar{
	{"char1", 1},
	{"char2", 2},
	{"char4", 4},
	{"char8", 8},
	{"char16", 16},
	{"char32", 32},
	{"char64", 64},
	{"char128", 128},
	{"char256", 256},
}

type ShowTableColumn struct {
	ColumnName     string
	IsChar         bool
	IsString       bool
	DataType       string
	CharLength     int
	ShardKeyFlag   bool
	PrimaryKeyFlag bool
	DictFlag       bool
	DataFlag       bool
	NullableFlag   bool
	StoreOnlyFlag  bool
}

type ShowTableResource struct {
	Name                    string
	TableType               ShowTableDistributionType
	ResultTable             bool
	TypeId                  string
	TypeSchema              map[string]interface{} // JSON
	TypeLabel               string
	Columns                 []ShowTableColumn
	Size                    int64
	FullSize                int64
	ShowTableAdditionalInfo *ShowTableAdditionalInfo
}

type ShowTableResourcesBySchema struct {
	Tables                     []ShowTableResource
	MaterializedViews          []ShowTableResource
	MaterializedExternalTables []ShowTableResource
	LogicalExternalTable       []ShowTableResource
	LogicalViews               []ShowTableResource
}

type ShowTableAdditionalInfo struct {
	AttributeIndexes         []string // sep by ;
	CollectionNames          string
	GlobalAccessMode         string
	IsDirty                  bool
	RecordBytes              int64
	RefreshMethod            string
	RemainingTableTtl        int64
	StrategyDefinition       string
	PartitionDefinitionsJson string
	RequestAvroJson          string
	RequestAvroType          string
	TableTtl                 int64
	CompressedColumns        map[string]string
	ForeignKeys              string // verify with TPC-DS
	ForeignShardKey          string // verify with TPC-DS
	IsViewPersisted          bool   // seems always be false
	LastRefreshTime          string
	NextRefreshTime          string
	PartitionType            string // NONE, LIST, ...
	IsAutomaticPartition     bool
	PartitionKeys            []string // sep by ,
	SchemaName               string
	TableMonitor             string
	UserChunkSize            int64
	ViewTableName            string
	ColumnMemoryUsageTotal   int64
	ColumnMemoryUsage        map[string]int64
	IndexMemoryUsageTotal    int64
	IndexMemoryUsage         map[string]int64
}

func NewDefaultShowTableOptions() *ShowTableOptions {
	return &ShowTableOptions{
		ForceSynchronous:   true,
		GetSizes:           true,
		ShowChildren:       true,
		NoErrorIfNotExists: false,
		GetColumnInfo:      true,
	}
}

func (gpudb *Kinetica) buildShowTableOptionsMap(ctx context.Context, options *ShowTableOptions) map[string]string {
	var (
		childSpan trace.Span
	)

	_, childSpan = gpudb.tracer.Start(ctx, "gpudb.buildShowTableOptionsMap()")
	defer childSpan.End()

	mapOptions := make(map[string]string)
	mapOptions["force_synchronous"] = strconv.FormatBool(options.ForceSynchronous)
	mapOptions["get_sizes"] = strconv.FormatBool(options.GetSizes)
	mapOptions["show_children"] = strconv.FormatBool(options.ShowChildren)
	mapOptions["no_error_if_not_exists"] = strconv.FormatBool(options.NoErrorIfNotExists)
	mapOptions["get_column_info"] = strconv.FormatBool(options.GetColumnInfo)

	return mapOptions
}

func (gpudb *Kinetica) ShowTableRaw(ctx context.Context, table string) (*ShowTableResponse, error) {
	var (
		childCtx  context.Context
		childSpan trace.Span
	)

	childCtx, childSpan = gpudb.tracer.Start(ctx, "gpudb.ShowTableRaw()")
	defer childSpan.End()

	return gpudb.ShowTableRawWithOpts(childCtx, table, NewDefaultShowTableOptions())
}

func (gpudb *Kinetica) ShowTableRawWithOpts(
	ctx context.Context, table string, options *ShowTableOptions) (*ShowTableResponse, error) {
	var (
		childCtx       context.Context
		childSpan      trace.Span
		showTableMutex *sync.Mutex
	)

	childCtx, childSpan = gpudb.tracer.Start(ctx, "gpudb.ShowTableRawWithOpts()")
	defer childSpan.End()
	showTableMutex = &sync.Mutex{}

	showTableMutex.Lock()
	mapOptions := gpudb.buildShowTableOptionsMap(childCtx, options)
	showTableMutex.Unlock()

	response := ShowTableResponse{}
	request := ShowTableRequest{TableName: table, Options: mapOptions}
	err := gpudb.submitRawRequest(
		childCtx, "/show/table", &Schemas.showTableRequest, &Schemas.showTableResponse, &request, &response)

	return &response, err
}

// Might be required for 7.0 which does not support /show/schemas
func (gpudb *Kinetica) ShowTableSchemas(ctx context.Context) (*[]string, error) {
	var (
		childCtx  context.Context
		childSpan trace.Span
	)

	childCtx, childSpan = gpudb.tracer.Start(ctx, "gpudb.ShowTableSchemas()")
	defer childSpan.End()

	rawResult, err := gpudb.ShowTableRawWithOpts(childCtx, "", &ShowTableOptions{
		ForceSynchronous:   true,
		GetSizes:           false,
		ShowChildren:       true,
		NoErrorIfNotExists: false,
		GetColumnInfo:      false,
	})
	if err != nil {
		return nil, err
	}
	schemas := []string{}
	expectedTypes := []string{"COLLECTION", "SCHEMA"}
	for i, objName := range rawResult.TableNames {
		if ContainsAnyStr(&rawResult.TableDescriptions[i], &expectedTypes) {
			schemas = append(schemas, objName)
		}
	}
	return &schemas, nil
}

func (gpudb *Kinetica) ShowTableSchemasCleaned(ctx context.Context) (*[]string, error) {
	var (
		childCtx  context.Context
		childSpan trace.Span
	)

	childCtx, childSpan = gpudb.tracer.Start(ctx, "gpudb.ShowTableSchemasCleaned()")
	defer childSpan.End()

	allSchemas, err := gpudb.ShowTableSchemas(childCtx)
	if err != nil {
		return nil, err
	}
	blacklistedSchemas := []string{"sys_sql_temp", "sys_temp", "sys_security", "SYSTEM", "kml", "kml_audit", "kml_audit_archive", "__SQL_TEMP", "__TEMP", "filesystem", "__REVEAL_TMP"}
	schemas := []string{}
	for _, aSchema := range *allSchemas {
		if !ContainsStr(&blacklistedSchemas, aSchema) {
			schemas = append(schemas, aSchema)
		}
	}
	return &schemas, err
}

func (gpudb *Kinetica) ShowTableResourcesBySchema(
	ctx context.Context, schema string, hideTempTables bool) (*ShowTableResourcesBySchema, error) {
	var (
		childSpan trace.Span
	)

	_, childSpan = gpudb.tracer.Start(ctx, "gpudb.ShowTableResourcesBySchema()")
	defer childSpan.End()

	rawResult, err := gpudb.ShowTableRawWithOpts(ctx, schema, &ShowTableOptions{
		ForceSynchronous:   true,
		GetSizes:           true,
		ShowChildren:       true,
		NoErrorIfNotExists: false,
		GetColumnInfo:      true,
	})
	if err != nil {
		return nil, err
	}
	result := new(ShowTableResourcesBySchema)
	const tempPrefix = "__TEMP"
	for i, resourceName := range rawResult.TableNames {
		// filter temp tables
		if hideTempTables && strings.HasPrefix(resourceName, tempPrefix) {
			continue
		}
		tableDescriptions := &rawResult.TableDescriptions[i]
		additionalInfo := &rawResult.AdditionalInfo[i]
		replicated := ContainsStr(tableDescriptions, "REPLICATED")
		tableType := resolveTableType(&rawResult.Properties[i], replicated)
		if ContainsStr(tableDescriptions, "MATERIALIZED_EXTERNAL_TABLE") {
			result.MaterializedExternalTables = append(result.MaterializedExternalTables, *constructShowTableResource(resourceName, i, tableType, rawResult))
		} else if ContainsStr(tableDescriptions, "MATERIALIZED_VIEW") {
			result.MaterializedViews = append(result.MaterializedViews, *constructShowTableResource(resourceName, i, tableType, rawResult))
		} else if ContainsStr(tableDescriptions, "LOGICAL_VIEW") || (*additionalInfo)["request_avro_type"] == "is_create_view" {
			result.LogicalViews = append(result.LogicalViews, *constructShowTableResource(resourceName, i, LogicalView, rawResult))
		} else if ContainsStr(tableDescriptions, "LOGICAL_EXTERNAL_TABLE") {
			result.LogicalExternalTable = append(result.LogicalExternalTable, *constructShowTableResource(resourceName, i, ExternalResource, rawResult))
		} else {
			result.Tables = append(result.Tables, *constructShowTableResource(resourceName, i, tableType, rawResult))
		}
	}
	// fmt.Println(rawResult)
	return result, nil
}

func constructShowTableResource(
	resourceName string, i int, tableType ShowTableDistributionType, rawResponse *ShowTableResponse) *ShowTableResource {
	// fmt.Println("Resource:", resourceName, " ", tableType)
	typeSchema := make(map[string]interface{})
	err := json.Unmarshal([]byte(rawResponse.TypeSchemas[i]), &typeSchema)
	if err != nil {
		typeSchema["error"] = err
	}
	return &ShowTableResource{
		Name:                    resourceName,
		TableType:               tableType,
		ResultTable:             ContainsStr(&rawResponse.TableDescriptions[i], "RESULT_TABLE"),
		TypeId:                  rawResponse.TypeIds[i],
		TypeSchema:              typeSchema,
		TypeLabel:               rawResponse.TypeLabels[i],
		Columns:                 *convertPropertiesToColumns(rawResponse.Properties[i], typeSchema),
		Size:                    rawResponse.Sizes[i],
		FullSize:                rawResponse.FullSizes[i],
		ShowTableAdditionalInfo: parseAdditionalInfo(rawResponse.AdditionalInfo[i]),
	}
}

func parseAdditionalInfo(m map[string]string) *ShowTableAdditionalInfo {
	columnInfo := make(map[string]interface{})
	err := json.Unmarshal([]byte(m["column_info"]), &columnInfo)
	if err != nil {
		fmt.Println(err)
	}
	columnMemoryUsage := convertToColumnMemoryUsage(columnInfo)
	indexMemoryUsage := convertToIndexMemoryUsage(columnInfo)
	return &ShowTableAdditionalInfo{
		AttributeIndexes:         convertPropertySliceString(m["attribute_indexes"], ";"),
		CollectionNames:          m["collection_names"],
		GlobalAccessMode:         m["global_access_mode"],
		IsDirty:                  convertPropertyBool(m["is_dirty"]),
		RecordBytes:              convertPropertyInt64(m["record_bytes"]),
		RefreshMethod:            m["refresh_method"],
		RemainingTableTtl:        convertPropertyInt64(m["remaining_table_ttl"]),
		StrategyDefinition:       m["strategy_definition"],
		PartitionDefinitionsJson: m["partition_definitions_json"],
		RequestAvroJson:          m["request_avro_json"],
		RequestAvroType:          m["request_avro_type"],
		TableTtl:                 convertPropertyInt64(m["table_ttl"]),
		CompressedColumns:        convertPropertyCompressedColumns(m["compressed_columns"]), // example compressed_columns -> {payment_type,lz4};{pickup_datetime,lz4};{pickup_latitude,lz4}
		ForeignKeys:              m["foreign_keys"],
		ForeignShardKey:          m["foreign_shard_key"],
		IsViewPersisted:          convertPropertyBool(m["is_view_persisted"]),
		LastRefreshTime:          m["last_refresh_time"],
		NextRefreshTime:          m["next_refresh_time"],
		PartitionType:            m["partition_type"],
		IsAutomaticPartition:     convertPropertyBool(m["is_automatic_partition"]),
		PartitionKeys:            convertPropertySliceString(m["partition_keys"], ","),
		SchemaName:               m["schema_name"],
		TableMonitor:             m["table_monitor"],
		UserChunkSize:            convertPropertyInt64(m["user_chunk_size"]),
		ViewTableName:            m["view_table_name"],
		ColumnMemoryUsageTotal:   sumMapValues(columnMemoryUsage),
		ColumnMemoryUsage:        columnMemoryUsage,
		IndexMemoryUsageTotal:    sumMapValues(indexMemoryUsage),
		IndexMemoryUsage:         indexMemoryUsage,
	}
}

func sumMapValues(info map[string]int64) int64 {
	var result int64 = 0
	for _, v := range info {
		result += v
	}
	return result
}

func convertToIndexMemoryUsage(info map[string]interface{}) map[string]int64 {
	result := make(map[string]int64)
	for k, v := range info {
		if strings.HasPrefix(k, "__index_") {
			newKey := strings.TrimPrefix(k, "__index_")
			subMap := v.(map[string]interface{})
			result[newKey] = int64(subMap["memory_usage"].(float64))
		}
	}
	return result
}

func convertToColumnMemoryUsage(info map[string]interface{}) map[string]int64 {
	result := make(map[string]int64)
	for k, v := range info {
		if !strings.HasPrefix(k, "__index_") {
			subMap := v.(map[string]interface{})
			result[k] = int64(subMap["memory_usage"].(float64))
		}
	}
	return result
}

func convertPropertyCompressedColumns(str string) map[string]string {
	splitted := convertPropertySliceString(str, ";")
	if splitted == nil {
		return nil
	}
	result := make(map[string]string)
	for _, v := range splitted {
		replaced := strings.ReplaceAll(v, "{", "")
		replaced = strings.ReplaceAll(replaced, "}", "")
		s := strings.Split(replaced, ",")
		result[s[0]] = s[1]
	}
	return result
}

func convertPropertySliceString(str string, split string) []string {
	if str == "" {
		return nil
	}
	return strings.Split(str, split)
}

func convertPropertyInt64(str string) int64 {
	if str == "" {
		return 0
	}
	res, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0
	}
	return res
}

func convertPropertyBool(str string) bool {
	if str == "true" {
		return true
	}
	return false
}

func convertPropertiesToColumns(m map[string][]string, schema map[string]interface{}) *[]ShowTableColumn {
	result := []ShowTableColumn{}
	if schema["fields"] != nil {
		fields := schema["fields"].([]interface{})
		for key, value := range m {
			containsChar, charLen := ContainsAnyCharType(&value)
			fieldType := findFieldType(&fields, key)
			column := ShowTableColumn{
				ColumnName:     key,
				IsChar:         containsChar,
				IsString:       fieldType == "string",
				DataType:       fieldType,
				CharLength:     charLen,
				PrimaryKeyFlag: ContainsStr(&value, "primary_key"),
				ShardKeyFlag:   ContainsStr(&value, "shard_key"),
				DictFlag:       ContainsStr(&value, "dict"),
				DataFlag:       ContainsStr(&value, "data"),
				NullableFlag:   ContainsStr(&value, "nullable"),
				StoreOnlyFlag:  ContainsStr(&value, "store_only"),
			}
			result = append(result, column)
		}
	}
	return &result
}

func findFieldType(types *[]interface{}, column string) string {
	for _, field := range *types {
		fieldMap := field.(map[string]interface{})
		name := fieldMap["name"].(string)
		if name == column {
			typeField := fieldMap["type"]
			switch v := typeField.(type) {
			case string:
				return v
			case []interface{}:
				if len(v) > 0 {
					return v[0].(string)
				} else {
					return "UNKNOWN"
				}
			}
		}
	}
	return "UNKNOWN"
}

func hasAnyShardKey(properties *map[string][]string) bool {
	expectedValues := []string{"primary_key", "shard_key"}
	for _, v := range *properties {
		if ContainsAnyStr(&v, &expectedValues) {
			return true
		}
	}
	return false
}

func resolveTableType(properties *map[string][]string, replicated bool) ShowTableDistributionType {
	hasAnyShardKey := hasAnyShardKey(properties)
	if hasAnyShardKey {
		return Sharded
	} else if replicated {
		return Replicated
	}
	return ShardedRandomly
}

func ContainsAnyCharType(s *[]string) (bool, int) {
	if s == nil {
		return false, -1
	}
	for _, a := range *s {
		for _, v := range CharTypes {
			if a == v.TypeName {
				return true, v.Length
			}
		}
	}
	return false, -1
}

func ContainsAnyStr(s *[]string, e *[]string) bool {
	if s == nil {
		return false
	}
	for _, a := range *s {
		for _, v := range *e {
			if a == v {
				return true
			}
		}
	}
	return false
}

func ContainsStr(s *[]string, e string) bool {
	return ContainsAnyStr(s, &[]string{e})
}

/*
Replicated tables
-> TableDescriptions = REPLICATED
-> Additionalinfo: request_avro_type = ""

Randomly sharded table
-> TableDescriptions = nil
-> foreign_key might be interesting to link the tables
-> Additionalinfo: request_avro_type = is_table
-> Shard_keys in Properties under the field name (primary_key or shard_key, if there is primary_key, the shard_key is not listed)
-> put in same list, but flag with randomly sharded keys

Sharded table
-> TableDescriptions = nil
-> foreign_shard_key might be interesting to link the tables
-> foreign_key might be interesting to link the tables
-> Additionalinfo: request_avro_type
-> Shard_keys in Properties under the field name (primary_key or shard_key, if there is primary_key, the shard_key is not listed)

Where are the indexes?

Materialized Ext Table (replicated)
-> TableDescriptions = [MATERIALIZED_EXTERNAL_TABLE, REPLICATED]
-> TypeLabels = projection_result

Materialized Ext Table (sharded)
-> TableDescriptions = [MATERIALIZED_EXTERNAL_TABLE]
-> TypeLabels = projection_result

Materialized View
-> TableDescriptions = [MATERIALIZED_VIEW, RESULT_TABLE]
-> TypeLabels = projection_result

Logical View
-> TableDescriptions = [LOGICAL_VIEW]

Logical External Table
-> TableDescriptions = [LOGICAL_EXTERNAL_TABLE]
-> TypeLabels contains list with paths!
Result tables (not MVs)
if there is only RESULT table use one von the normal above and flag with result table

Indexes
-> AdditionalInfo -> attribute_indexes = field1;field2;field3
- request_avro_type = is_external_table

General useful properties:
- AdditionalInfo
	- collection_names
	- foreign_keys
	- is_dirty
	- is_view_persisted
	- record_bytes
	- user_chunk_size
	- request_avro_type
	- attribute_indexes
	- compressed_columns
	- foreign_shard_keys
	- partition_definitions
	- partition_keys
	- remaining_table_ttl
	- ... all in Struct
*/
