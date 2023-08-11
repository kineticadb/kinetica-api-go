package kinetica

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/hamba/avro/v2"
	"github.com/mitchellh/mapstructure"
	"github.com/samber/lo"
	"go.opentelemetry.io/otel/trace"
)

func mapToStruct(m map[string]interface{}) reflect.Value {
	var structFields []reflect.StructField

	for k, v := range m {
		sf := reflect.StructField{
			Name: strings.Title(k),
			Type: reflect.TypeOf(v),
		}
		structFields = append(structFields, sf)
	}

	// Creates the struct type
	structType := reflect.StructOf(structFields)

	// Creates a new struct
	return reflect.New(structType)
}

func mergeStructs(structs ...interface{}) reflect.Type {
	var structFields []reflect.StructField
	var structFieldNames []string

	for _, item := range structs {
		rt := reflect.TypeOf(item)
		for i := 0; i < rt.NumField(); i++ {
			field := rt.Field(i)
			if !stringInSlice(field.Name, structFieldNames) {
				structFields = append(structFields, field)
				structFieldNames = append(structFieldNames, field.Name)
			}
		}
	}

	return reflect.StructOf(structFields)
}

func stringInSlice(value string, slice []string) bool {
	for _, elem := range slice {
		if elem == value {
			return true
		}
	}
	return false
}

func NewDefaultExecuteSqlOptions() *ExecuteSqlOptions {
	return &ExecuteSqlOptions{
		Encoding:              "binary",
		ParallelExecution:     true,
		CostBasedOptimization: false,
		PlanCache:             true,
		RuleBasedOptimization: true,
		ResultsCaching:        true,
		PagingTable:           "",
		PagingTableTtl:        5,
		DistributedJoins:      true,
		DistributedOperations: true,
		SsqOptimization:       true,
		LateMaterialization:   false,
		Ttl:                   0,
		UpdateOnExistingPk:    false,
		PreserveDictEncoding:  true,
		ValidateChangeColumn:  true,
		PrepareMode:           false,
	}
}

func (kinetica *Kinetica) ExecuteSqlRaw(
	ctx context.Context,
	statement string, offset int64, limit int64, requestSchemaStr string,
	data []byte) (*ExecuteSqlResponse, error) {
	var (
		childCtx  context.Context
		childSpan trace.Span
	)

	childCtx, childSpan = kinetica.tracer.Start(ctx, "kinetica.ExecuteSqlRaw()")
	defer childSpan.End()

	return kinetica.ExecuteSqlRawWithOpts(childCtx, statement, offset, limit, requestSchemaStr,
		data, NewDefaultExecuteSqlOptions())
}

func (kinetica *Kinetica) ExecuteSqlRawWithOpts(
	ctx context.Context,
	statement string, offset int64, limit int64, requestSchemaStr string,
	data []byte, options *ExecuteSqlOptions) (*ExecuteSqlResponse, error) {
	var (
		childCtx  context.Context
		childSpan trace.Span
	)

	childCtx, childSpan = kinetica.tracer.Start(ctx, "kinetica.ExecuteSqlRawWithOpts()")
	defer childSpan.End()

	mapOptions := kinetica.buildExecuteSqlOptionsMap(childCtx, options)
	response := ExecuteSqlResponse{}
	request := ExecuteSqlRequest{Statement: statement, Offset: offset, Limit: limit,
		RequestSchema: requestSchemaStr, Data: data, Encoding: options.Encoding, Options: *mapOptions}
	err := kinetica.submitRawRequest(
		childCtx, "/execute/sql",
		&Schemas.executeSqlRequest, &Schemas.executeSqlResponse,
		&request, &response)

	return &response, err
}

func (kinetica *Kinetica) buildExecuteSqlOptionsMap(ctx context.Context, options *ExecuteSqlOptions) *map[string]string {
	var (
		childSpan trace.Span
	)

	_, childSpan = kinetica.tracer.Start(ctx, "kinetica.buildExecuteSqlOptionsMap()")
	defer childSpan.End()

	mapOptions := make(map[string]string)
	mapOptions["parallel_execution"] = strconv.FormatBool(options.ParallelExecution)
	mapOptions["cost_based_optimization"] = strconv.FormatBool(options.CostBasedOptimization)
	mapOptions["plan_cache"] = strconv.FormatBool(options.PlanCache)
	mapOptions["rule_based_optimization"] = strconv.FormatBool(options.RuleBasedOptimization)
	mapOptions["results_caching"] = strconv.FormatBool(options.ResultsCaching)

	if options.PagingTable != "" {
		mapOptions["paging_table"] = options.PagingTable
	}

	mapOptions["paging_table_ttl"] = strconv.FormatInt(options.PagingTableTtl, 10)
	mapOptions["distributed_joins"] = strconv.FormatBool(options.DistributedJoins)
	mapOptions["distributed_operations"] = strconv.FormatBool(options.DistributedOperations)
	mapOptions["ssq_optimization"] = strconv.FormatBool(options.SsqOptimization)
	mapOptions["late_materialization"] = strconv.FormatBool(options.LateMaterialization)
	mapOptions["ttl"] = strconv.FormatInt(options.Ttl, 10)
	mapOptions["update_on_existing_pk"] = strconv.FormatBool(options.UpdateOnExistingPk)
	mapOptions["preserve_dict_encoding"] = strconv.FormatBool(options.PreserveDictEncoding)
	mapOptions["validate_change_column"] = strconv.FormatBool(options.ValidateChangeColumn)
	mapOptions["prepare_mode"] = strconv.FormatBool(options.PrepareMode)

	return &mapOptions
}

type ExecuteSqlMapResult struct {
	*ExecuteSqlResponse
	ResultsMap *[]map[string]interface{}
}

func (kinetica *Kinetica) ExecuteSqlMap(
	ctx context.Context, sql string, offset int64, limit int64) (*ExecuteSqlMapResult, error) {
	var (
		childCtx  context.Context
		childSpan trace.Span
	)

	childCtx, childSpan = kinetica.tracer.Start(ctx, "kinetica.ExecuteSqlMap()")
	defer childSpan.End()

	return kinetica.ExecuteSqlMapWithOpts(childCtx, sql, offset, limit, NewDefaultExecuteSqlOptions())
}

func (kinetica *Kinetica) ExecuteSqlMapWithOpts(
	ctx context.Context, statement string, offset int64, limit int64,
	options *ExecuteSqlOptions) (*ExecuteSqlMapResult, error) {
	var (
		childCtx  context.Context
		childSpan trace.Span
	)

	childCtx, childSpan = kinetica.tracer.Start(ctx, "kinetica.ExecuteSqlMapWithOpts()")
	defer childSpan.End()

	raw, err := kinetica.ExecuteSqlRawWithOpts(childCtx, statement, offset, limit, "", nil, options)
	if err != nil {
		return nil, err
	}
	recordSchema, recordErr := avro.Parse(raw.ResponseSchema)
	if recordErr != nil {
		return nil, recordErr
	}
	// start := time.Now()
	if options.Encoding == "binary" {
		tmpResultList := make(map[string]interface{})
		fmt.Println(recordSchema)
		avro.Unmarshal(recordSchema, raw.BinaryEncodedResponse, &tmpResultList)
		fmt.Println(tmpResultList)
		col1 := tmpResultList["column_1"].([]interface{})
		columnsHeaders := tmpResultList["column_headers"].([]interface{})
		columnsTypes := tmpResultList["column_datatypes"].([]interface{})
		fmt.Println(columnsTypes...)
		resultList := make([]map[string]interface{}, len(col1))
		// convert to line oriented format
		for colIndex, header := range columnsHeaders {
			headerStr := header.(string)
			tmpColumnName := fmt.Sprint("column_", (colIndex + 1))
			columnValues := tmpResultList[tmpColumnName].([]interface{})
			fmt.Println(columnValues...)
			for vi, value := range columnValues {
				if resultList[vi] == nil {
					resultList[vi] = make(map[string]interface{})
				}
				resultList[vi][headerStr] = value
			}
		}
		return &ExecuteSqlMapResult{raw, &resultList}, nil
		// duration := time.Since(start)
		// fmt.Println("GetRecordsMap", duration.Milliseconds(), " ns")
	} else {
		// TODO
		panic("JSON decoding is not yet implemented")
	}
}

type ExecuteSqlStructResult struct {
	*ExecuteSqlResponse
	ResultsStruct *[]interface{}
}

func (kinetica *Kinetica) ExecuteSqlStruct(
	ctx context.Context,
	statement string, offset int64, limit int64, newInstance func() interface{}) (*ExecuteSqlStructResult, error) {
	var (
		childCtx  context.Context
		childSpan trace.Span
	)

	childCtx, childSpan = kinetica.tracer.Start(ctx, "kinetica.ExecuteSqlStruct()")
	defer childSpan.End()

	return kinetica.ExecuteSqlStructWithOpts(childCtx, statement, offset, limit, NewDefaultExecuteSqlOptions(), newInstance)
}

func (kinetica *Kinetica) ExecuteSqlStructWithOpts(
	ctx context.Context,
	statement string, offset int64, limit int64, options *ExecuteSqlOptions,
	newInstance func() interface{}) (*ExecuteSqlStructResult, error) {
	var (
		childCtx  context.Context
		childSpan trace.Span
	)

	childCtx, childSpan = kinetica.tracer.Start(ctx, "kinetica.ExecuteSqlStructWithOpts()")
	defer childSpan.End()

	result, err := kinetica.ExecuteSqlMapWithOpts(childCtx, statement, offset, limit, options)
	if err != nil {
		return nil, err
	}
	// start := time.Now()
	resultList := make([]interface{}, len(*result.ResultsMap))
	for i, valueMap := range *result.ResultsMap {
		newInst := newInstance()
		resultList[i] = newInst
		err = mapstructure.Decode(valueMap, &resultList[i])
		if err != nil {
			return nil, err
		}
	}
	// duration := time.Since(start)
	// fmt.Println("GetRecordsStruct", duration.Milliseconds(), " ns")

	return &ExecuteSqlStructResult{result.ExecuteSqlResponse, &resultList}, nil
}

// ExecuteSQLStructWithOpts - Generic method which is parameterized over user defined structs to receive results
//
//	@param ctx
//	@param kinetica
//	@param statement
//	@param offset
//	@param limit
//	@param options
//	@return []T
//	@return error
func ExecuteSQLStructWithOpts[T any](
	ctx context.Context, kinetica Kinetica,
	statement string, offset int64, limit int64, options *ExecuteSqlOptions) ([]T, error) {
	var (
		childCtx   context.Context
		childSpan  trace.Span
		resultList []T
	)

	childCtx, childSpan = kinetica.tracer.Start(ctx, "Generic kinetica.ExecuteSqlStructWithOpts()")
	defer childSpan.End()

	// Force the limit to 0 to create the paging table
	result, err := kinetica.ExecuteSqlMapWithOpts(childCtx, statement, offset, 0, options)
	if err != nil {
		return nil, err
	}

	pagingTable := result.ExecuteSqlResponse.PagingTable
	numRecords := result.TotalNumberOfRecords

	resultList = make([]T, numRecords)

	switch {
	case numRecords > 10000:
		recordRange := make([]int64, numRecords)
		for i := int64(0); i < numRecords; i++ {
			recordRange[i] = i
		}

		recordRanges := lo.Chunk(recordRange, 10000)

		lo.ForEach(recordRanges, func(val []int64, i int) {
			offset := lo.Min(val)
			limit := lo.Max(val)
			result, err := GetRecordsStructWithOpts[T](context.TODO(), kinetica, pagingTable, offset, limit, NewDefaultGetRecordsOptions(), func() T { return *new(T) })
			if err != nil {
				kinetica.kineticaLogger.Error(err.Error())
			} else {
				resultList = append(resultList, result...)
			}
		})

	default:
		result, err := GetRecordsStructWithOpts[T](context.TODO(), kinetica, pagingTable, 0, numRecords, NewDefaultGetRecordsOptions(), func() T { return *new(T) })
		if err != nil {
			kinetica.kineticaLogger.Error(err.Error())
		} else {
			resultList = append(resultList, result...)
		}
	}

	return resultList, nil
}

func (kinetica *Kinetica) ShowTableDDL(ctx context.Context, table string) (*string, error) {
	var (
		childCtx  context.Context
		childSpan trace.Span
	)

	childCtx, childSpan = kinetica.tracer.Start(ctx, "kinetica.ShowTableDDL()")
	defer childSpan.End()

	result, err := kinetica.ExecuteSqlMap(childCtx, "SHOW "+table, 0, 10)
	if err != nil {
		return nil, err
	}
	if len(*result.ResultsMap) == 0 {
		return nil, errors.New("No DDL returned by query. Result set size is 0")
	}
	// return nil, nil
	ddl := (*result.ResultsMap)[0]["DDL"].(string)
	return &ddl, nil
}

func (kinetica *Kinetica) ShowStoredProcedureDDL(ctx context.Context, proc string) (*string, error) {
	var (
		childCtx  context.Context
		childSpan trace.Span
	)

	childCtx, childSpan = kinetica.tracer.Start(ctx, "kinetica.ShowStoredProcedureDDL()")
	defer childSpan.End()

	result, err := kinetica.ExecuteSqlMap(childCtx, "SHOW PROCEDURE "+proc, 0, 10)
	if err != nil {
		return nil, err
	}
	if len(*result.ResultsMap) == 0 {
		return nil, errors.New("No DDL returned by query. Result set size is 0")
	}
	// return nil, nil
	ddl := strings.TrimSpace((*result.ResultsMap)[0]["PROCEDURE_DEFINITION"].(string))
	return &ddl, nil
}

type ShowExplainVerboseSqlPlan struct {
	Id int `json:"ID,string"`
	/*
		{2 @ rank:1, tom:0
		filter-plan-in  |--alias--|-num_chunks-|---count----|-set-name----
		filter-plan-in  | TableAlias_0_ |          1 |                   17 | sys_sql_temp.40375_Aggregate_2_692dc858-7a90-11eb-b9ee-0242ac110002
		filter-plan-in  | TableAlias_1_ |          1 |                   34 | sys_sql_temp.40375_Aggregate_4_692dc8b2-7a90-11eb-ba54-0242ac110002
		filter-plan-step|--i-|----time----|----count----|---in_count--|--out_count--|-J-|-----filter-type-|---set-indices---|---------stencil-types----------|---expression
		filter-plan-step|  0 |   0.002362 |         578 |             |             |   |           start |                 |  B0 B1                         |
		filter-plan-step|  1 |   0.001226 |         122 |         578 |         122 |   |   equi-join:1:1 |  1,0            |  E1 E1                         | (TableAlias_0_.vendor_id == TableAlias_1_.vendor_id)
		filter-plan-out | count-time = 8e-07 |  count =         122 | set-name = filter_planner_view_65
	*/
	AdditionalInfo     string  `json:"ADDITIONAL_INFO"` // this is a large text block
	Columns            string  `json:"COLUMNS"`
	Options            string  `json:"OPTIONS"`
	RunTime            float64 `json:"RUN_TIME,string"`
	TableDefinitions   string  `json:"TABLE_DEFINITIONS"`
	Dependencies       string  `json:"DEPENDENCIES"`
	Endpoint           string  `json:"ENDPOINT"`
	JsonRequest        string  `json:"JSON_REQUEST"`
	LastUseTables      string  `json:"LAST_USE_TABLES"`
	Expressions        string  `json:"EXPRESSIONS"`
	ResultDistribution string  `json:"RESULT_DISTRIBUTION"` // e.g. "NA / replicated" or "shardkey10,shardkey20; / shardkey1,shardkey2;"
	ResultRows         int64   `json:"RESULT_ROWS,string"`
	Parent             *ShowExplainVerboseSqlResult
}

type ShowExplainVerboseSqlResult struct {
	Plans *[]ShowExplainVerboseSqlPlan `json:"PLAN"`
}

func (kinetica *Kinetica) ShowExplainVerboseSqlStatement(ctx context.Context, statement string) (*ShowExplainVerboseSqlResult, error) {
	var (
		childCtx  context.Context
		childSpan trace.Span
	)

	childCtx, childSpan = kinetica.tracer.Start(ctx, "kinetica.ShowExplainVerboseSqlStatement()")
	defer childSpan.End()

	return kinetica.parseExplainVerboseAnalyseSqlStatement(childCtx, "EXPLAIN VERBOSE FORMAT JSON "+statement)
}

func (kinetica *Kinetica) ShowExplainVerboseAnalyseSqlStatement(ctx context.Context, statement string) (*ShowExplainVerboseSqlResult, error) {
	var (
		childCtx  context.Context
		childSpan trace.Span
	)

	childCtx, childSpan = kinetica.tracer.Start(ctx, "kinetica.ShowExplainVerboseAnalyseSqlStatement()")
	defer childSpan.End()

	return kinetica.parseExplainVerboseAnalyseSqlStatement(childCtx, "EXPLAIN VERBOSE ANALYZE FORMAT JSON "+statement)
}

func (kinetica *Kinetica) parseExplainVerboseAnalyseSqlStatement(
	ctx context.Context,
	explainStatement string) (*ShowExplainVerboseSqlResult, error) {
	var (
		childCtx  context.Context
		childSpan trace.Span
	)

	childCtx, childSpan = kinetica.tracer.Start(ctx, "kinetica.parseExplainVerboseAnalyseSqlStatement()")
	defer childSpan.End()

	// explain verbose analyze format json
	result, err := kinetica.ExecuteSqlMap(childCtx, explainStatement, 0, 10)
	if err != nil {
		return nil, err
	}
	if len(*result.ResultsMap) == 0 {
		return nil, errors.New("No explain returned by query. Result set size is 0")
	}
	jsonStr := strings.TrimSpace((*result.ResultsMap)[0]["EXPLAIN"].(string))
	resultStruct := ShowExplainVerboseSqlResult{}
	err = json.Unmarshal([]byte(jsonStr), &resultStruct)
	if err != nil {
		return nil, err
	}
	// Assign Parent, so mthat we can resolve dependencies later.
	plans := *resultStruct.Plans
	for i, _ := range plans {
		// fmt.Printf("FIRST %p\n", &plan)
		plan := &plans[i]
		plan.Parent = &resultStruct
	}
	// return nil, nil
	// ddl := strings.TrimSpace((*result.ResultsMap)[0]["PROCEDURE_DEFINITION"].(string))
	return &resultStruct, nil
}

func (plan *ShowExplainVerboseSqlPlan) JsonRequestMap() (*map[string]interface{}, error) {
	resultMap := make(map[string]interface{})
	if plan.JsonRequest == "" {
		return &resultMap, nil
	}
	err := json.Unmarshal([]byte(plan.JsonRequest), &resultMap)
	if err != nil {
		return nil, err
	}
	return &resultMap, nil
}

func (plan *ShowExplainVerboseSqlPlan) FindDependentPlans() (*[]ShowExplainVerboseSqlPlan, error) {
	result := []ShowExplainVerboseSqlPlan{}
	if plan.Dependencies == "" {
		return &result, nil
	}
	split := strings.Split(plan.Dependencies, ",")
	for _, dependencyStr := range split {
		dependency, err := strconv.Atoi(dependencyStr)
		if err != nil {
			return nil, err
		}
		plans := plan.Parent.Plans
		for _, plan := range *plans {
			if dependency == plan.Id {
				result = append(result, plan)
			}
		}
	}
	return &result, nil
}
