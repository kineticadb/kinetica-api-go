package kinetica

import "github.com/hamba/avro/v2"

type AvroSchemas struct {
	adminAddHostRequest            avro.Schema
	adminAddHostResponse           avro.Schema
	adminAddRanksRequest           avro.Schema
	adminAddRanksResponse          avro.Schema
	adminBackupBeginRequest        avro.Schema
	adminBackupBeginResponse       avro.Schema
	adminBackupEndRequest          avro.Schema
	adminBackupEndResponse         avro.Schema
	adminRemoveHostRequest         avro.Schema
	adminRemoveHostResponse        avro.Schema
	adminRemoveRanksRequest        avro.Schema
	adminRemoveRanksResponse       avro.Schema
	adminShutdownRequest           avro.Schema
	adminShutdownResponse          avro.Schema
	alterGraphRequest              avro.Schema
	alterGraphResponse             avro.Schema
	alterUserRequest               avro.Schema
	alterUserResponse              avro.Schema
	createSchemaRequest            avro.Schema
	createSchemaResponse           avro.Schema
	dropSchemaRequest              avro.Schema
	dropSchemaResponse             avro.Schema
	createResourceGroupRequest     avro.Schema
	createResourceGroupResponse    avro.Schema
	deleteResourceGroupRequest     avro.Schema
	deleteResourceGroupResponse    avro.Schema
	getRecordsRequest              avro.Schema
	getRecordsResponse             avro.Schema
	insertRecordsRequest           avro.Schema
	insertRecordsResponse          avro.Schema
	executeSqlRequest              avro.Schema
	executeSqlResponse             avro.Schema
	showTableRequest               avro.Schema
	showTableResponse              avro.Schema
	showSchemaRequest              avro.Schema
	showSchemaResponse             avro.Schema
	showSystemPropertiesRequest    avro.Schema
	showSystemPropertiesResponse   avro.Schema
	showSystemTimingRequest        avro.Schema
	showSystemTimingResponse       avro.Schema
	showSystemStatusRequest        avro.Schema
	showSystemStatusResponse       avro.Schema
	showSqlProcRequest             avro.Schema
	showSqlProcResponse            avro.Schema
	showResourceStatisticsRequest  avro.Schema
	showResourceStatisticsResponse avro.Schema
	showSecurityRequest            avro.Schema
	showSecurityResponse           avro.Schema
	showResourceGroupsRequest      avro.Schema
	showResourceGroupsResponse     avro.Schema
	aggregateStatisticsRequest     avro.Schema
	aggregateStatisticsResponse    avro.Schema
	aggregateGroupByRequest        avro.Schema
	aggregateGroupByResponse       avro.Schema
	createJobRequest               avro.Schema
	createJobResponse              avro.Schema
	getJobRequest                  avro.Schema
	getJobResponse                 avro.Schema
}

var Schemas = AvroSchemas{
	adminAddHostRequest:            *parseSchema("avsc/admin_add_host_request.json"),
	adminAddHostResponse:           *parseSchema("avsc/admin_add_host_response.json"),
	adminAddRanksRequest:           *parseSchema("avsc/admin_add_ranks_request.json"),
	adminAddRanksResponse:          *parseSchema("avsc/admin_add_ranks_response.json"),
	adminBackupBeginRequest:        *parseSchema("avsc/admin_backup_begin_request.json"),
	adminBackupBeginResponse:       *parseSchema("avsc/admin_backup_begin_response.json"),
	adminBackupEndRequest:          *parseSchema("avsc/admin_backup_end_request.json"),
	adminBackupEndResponse:         *parseSchema("avsc/admin_backup_end_response.json"),
	adminRemoveHostRequest:         *parseSchema("avsc/admin_remove_host_request.json"),
	adminRemoveHostResponse:        *parseSchema("avsc/admin_remove_host_response.json"),
	adminRemoveRanksRequest:        *parseSchema("avsc/admin_remove_ranks_request.json"),
	adminRemoveRanksResponse:       *parseSchema("avsc/admin_remove_ranks_response.json"),
	adminShutdownRequest:           *parseSchema("avsc/admin_shutdown_request.json"),
	adminShutdownResponse:          *parseSchema("avsc/admin_shutdown_response.json"),
	alterGraphRequest:              *parseSchema("avsc/alter_graph_request.json"),
	alterGraphResponse:             *parseSchema("avsc/alter_graph_response.json"),
	alterUserRequest:               *parseSchema("avsc/alter_user_request.json"),
	alterUserResponse:              *parseSchema("avsc/alter_user_response.json"),
	getRecordsRequest:              *parseSchema("avsc/get_records_request.json"),
	getRecordsResponse:             *parseSchema("avsc/get_records_response.json"),
	insertRecordsRequest:           *parseSchema("avsc/insert_records_request.json"),
	insertRecordsResponse:          *parseSchema("avsc/insert_records_response.json"),
	executeSqlRequest:              *parseSchema("avsc/execute_sql_request.json"),
	executeSqlResponse:             *parseSchema("avsc/execute_sql_response.json"),
	showTableRequest:               *parseSchema("avsc/show_table_request.json"),
	showTableResponse:              *parseSchema("avsc/show_table_response.json"),
	showSchemaRequest:              *parseSchema("avsc/show_schema_request.json"),
	showSchemaResponse:             *parseSchema("avsc/show_schema_response.json"),
	showSystemPropertiesRequest:    *parseSchema("avsc/show_system_properties_request.json"),
	showSystemPropertiesResponse:   *parseSchema("avsc/show_system_properties_response.json"),
	showSystemTimingRequest:        *parseSchema("avsc/show_system_timing_request.json"),
	showSystemTimingResponse:       *parseSchema("avsc/show_system_timing_response.json"),
	showSystemStatusRequest:        *parseSchema("avsc/show_system_status_request.json"),
	showSystemStatusResponse:       *parseSchema("avsc/show_system_status_response.json"),
	showSqlProcRequest:             *parseSchema("avsc/show_sql_proc_request.json"),
	showSqlProcResponse:            *parseSchema("avsc/show_sql_proc_response.json"),
	showResourceStatisticsRequest:  *parseSchema("avsc/show_resource_statistics_request.json"),
	showResourceStatisticsResponse: *parseSchema("avsc/show_resource_statistics_response.json"),
	showSecurityRequest:            *parseSchema("avsc/show_security_request.json"),
	showSecurityResponse:           *parseSchema("avsc/show_security_response.json"),
	showResourceGroupsRequest:      *parseSchema("avsc/show_resource_groups_request.json"),
	showResourceGroupsResponse:     *parseSchema("avsc/show_resource_groups_response.json"),
	aggregateStatisticsRequest:     *parseSchema("avsc/aggregate_statistics_request.json"),
	aggregateStatisticsResponse:    *parseSchema("avsc/aggregate_statistics_response.json"),
	aggregateGroupByRequest:        *parseSchema("avsc/aggregate_group_by_request.json"),
	aggregateGroupByResponse:       *parseSchema("avsc/aggregate_group_by_response.json"),
	createJobRequest:               *parseSchema("avsc/create_job_request.json"),
	createJobResponse:              *parseSchema("avsc/create_job_response.json"),
	getJobRequest:                  *parseSchema("avsc/get_job_request.json"),
	getJobResponse:                 *parseSchema("avsc/get_job_response.json"),
	createSchemaRequest:            *parseSchema("avsc/create_schema_request.json"),
	createSchemaResponse:           *parseSchema("avsc/create_schema_response.json"),
	dropSchemaRequest:              *parseSchema("avsc/drop_schema_request.json"),
	dropSchemaResponse:             *parseSchema("avsc/drop_schema_response.json"),
	createResourceGroupRequest:     *parseSchema("avsc/create_resource_group_request.json"),
	createResourceGroupResponse:    *parseSchema("avsc/create_resource_group_response.json"),
	deleteResourceGroupRequest:     *parseSchema("avsc/delete_resource_group_request.json"),
	deleteResourceGroupResponse:    *parseSchema("avsc/delete_resource_group_response.json"),
}

// GetRecordsOptions --------------------------------------------
type GetRecordsOptions struct {
	Expression      string
	FastIndexLookup bool
	SortBy          string
	SortOrder       string
	Encoding        string
}

type GetRecordsRequest struct {
	TableName string            `avro:"table_name"`
	Offset    int64             `avro:"offset"`
	Limit     int64             `avro:"limit"`
	Encoding  string            `avro:"encoding"`
	Options   map[string]string `avro:"options"`
}

type GetRecordsResponse struct {
	TableName            string            `avro:"table_name"`
	TypeName             string            `avro:"type_name"`
	TypeSchema           string            `avro:"type_schema"`
	RecordsBinary        [][]byte          `avro:"records_binary"`
	RecordsJson          []string          `avro:"records_json"`
	TotalNumberOfRecords int64             `avro:"total_number_of_records"`
	HasMoreRecords       bool              `avro:"has_more_records"`
	Info                 map[string]string `avro:"info"`
}

// InsertRecords ---------------------------
type InsertRecordsOptions struct {
	UpdateOnExistingPk     bool
	IgnoreExistingPk       bool
	ReturnRecordIDs        bool
	TruncateStrings        bool
	ReturnIndividualErrors bool
	AllowPartialBatch      bool
}

type InsertRecordsRequest struct {
	TableName    string            `avro:"table_name"`
	List         [][]byte          `avro:"list"`
	ListString   [][]string        `avro:"list_str"`
	ListEncoding string            `avro:"list_encoding"`
	Options      map[string]string `avro:"options"`
}

type InsertRecordsResponse struct {
	RecordIDs     []string          `avro:"record_ids"`
	CountInserted int               `avro:"count_inserted"`
	CountUpdated  int               `avro:"count_updated"`
	Info          map[string]string `avro:"info"`
}

// ShowTableOptions --------------------------------------------
type ShowTableOptions struct {
	ForceSynchronous   bool
	GetSizes           bool
	ShowChildren       bool
	NoErrorIfNotExists bool
	GetColumnInfo      bool
}

type ShowTableRequest struct {
	TableName string            `avro:"table_name"`
	Options   map[string]string `avro:"options"`
}

type ShowTableResponse struct {
	TableName         string                `avro:"table_name"`
	TableNames        []string              `avro:"table_names"`
	TableDescriptions [][]string            `avro:"table_descriptions"`
	TypeIds           []string              `avro:"type_ids"`
	TypeSchemas       []string              `avro:"type_schemas"`
	TypeLabels        []string              `avro:"type_labels"`
	Properties        []map[string][]string `avro:"properties"`
	AdditionalInfo    []map[string]string   `avro:"additional_info"`
	Sizes             []int64               `avro:"sizes"`
	FullSizes         []int64               `avro:"full_sizes"`
	JoinSizes         []float64             `avro:"join_sizes"`
	TotalSize         int64                 `avro:"total_size"`
	TotalFullSize     int64                 `avro:"total_full_size"`
	Info              map[string]string     `avro:"info"`
}

// ShowSchemaOptions --------------------------------------------
type ShowSchemaOptions struct {
	NoErrorIfNotExists bool
}

type ShowSchemaRequest struct {
	SchemaName string            `avro:"schema_name"`
	Options    map[string]string `avro:"options"`
}

type ShowSchemaResponse struct {
	SchemaName     string              `avro:"schema_name"`
	SchemaNames    []string            `avro:"schema_names"`
	AdditionalInfo []map[string]string `avro:"additional_info"`
	Info           map[string]string   `avro:"info"`
}

// ExecuteSqlOptions --------------------------------------------
type ExecuteSqlOptions struct {
	Encoding              string
	ParallelExecution     bool
	CostBasedOptimization bool
	PlanCache             bool
	RuleBasedOptimization bool
	ResultsCaching        bool
	PagingTable           string
	PagingTableTtl        int64
	DistributedJoins      bool
	DistributedOperations bool
	SsqOptimization       bool
	LateMaterialization   bool
	Ttl                   int64
	UpdateOnExistingPk    bool
	PreserveDictEncoding  bool
	ValidateChangeColumn  bool
	PrepareMode           bool
}

type ExecuteSqlRequest struct {
	Statement     string            `avro:"statement"`
	Offset        int64             `avro:"offset"`
	Limit         int64             `avro:"limit"`
	Encoding      string            `avro:"encoding"`
	RequestSchema string            `avro:"request_schema_str"`
	Data          []byte            `avro:"data"`
	Options       map[string]string `avro:"options"`
}

type ExecuteSqlResponse struct {
	CountAffected         int64             `avro:"count_affected"`
	ResponseSchema        string            `avro:"response_schema_str"`
	BinaryEncodedResponse []byte            `avro:"binary_encoded_response"`
	JsonEncodedResponse   string            `avro:"json_encoded_response"`
	TotalNumberOfRecords  int64             `avro:"total_number_of_records"`
	HasMoreRecords        bool              `avro:"has_more_records"`
	PagingTable           string            `avro:"paging_table"`
	Info                  map[string]string `avro:"info"`
}

// ShowSystemPropertiesOptions --------------------------------------------
type ShowSystemPropertiesOptions struct {
	Properties string
}

type ShowSystemPropertiesRequest struct {
	Options map[string]string `avro:"options"`
}

type ShowSystemPropertiesResponse struct {
	PropertyMap map[string]string `avro:"property_map"`
	Info        map[string]string `avro:"info"`
}

// ShowSystemTimingOptions --------------------------------------------
type ShowSystemTimingOptions struct {
}

type ShowSystemTimingRequest struct {
	Options map[string]string `avro:"options"`
}

type ShowSystemTimingResponse struct {
	Endpoints []string          `avro:"endpoints"`
	TimeInMs  []float32         `avro:"time_in_ms"`
	JobIds    []string          `avro:"jobIds"`
	Info      map[string]string `avro:"info"`
}

// ShowSystemStatusOptions --------------------------------------------
type ShowSystemStatusOptions struct {
}

type ShowSystemStatusRequest struct {
	Options map[string]string `avro:"options"`
}

type ShowSystemStatusResponse struct {
	StatusMap map[string]string `avro:"status_map"`
	Info      map[string]string `avro:"info"`
}

// ShowSqlProcOptions --------------------------------------------
type ShowSqlProcOptions struct {
	NoErrorIfNotExists bool
}

type ShowSqlProcRequest struct {
	ProcedureName string            `avro:"procedure_name"`
	Options       map[string]string `avro:"options"`
}

type ShowSqlProcResponse struct {
	ProcedureNames       []string            `avro:"procedure_names"`
	ProcedureDefinitions []string            `avro:"procedure_definitions"`
	AdditionalInfo       []map[string]string `avro:"additional_info"`
	Info                 map[string]string   `avro:"info"`
}

// ShowResourceStatisticsOptions --------------------------------------------
type ShowResourceStatisticsOptions struct {
	TableNames string
}

type ShowResourceStatisticsRequest struct {
	Options map[string]string `avro:"options"`
}

type ShowResourceStatisticsResponse struct {
	StatisticsMap map[string]string `avro:"statistics_map"`
	Info          map[string]string `avro:"info"`
}

// ShowSecurityOptions --------------------------------------------
type ShowSecurityOptions struct {
}

type ShowSecurityRequest struct {
	Names   []string          `avro:"names"`
	Options map[string]string `avro:"options"`
}

type ShowSecurityResponse struct {
	Types          map[string]string              `avro:"types"`
	Roles          map[string][]string            `avro:"roles"`
	Permissions    map[string][]map[string]string `avro:"permissions"`
	ResourceGroups map[string][]string            `avro:"resource_groups"`
	Info           map[string]string              `avro:"info"`
}

// ShowResourceGroupsOptions --------------------------------------------
type ShowResourceGroupsOptions struct {
	ShowDefaultValues bool
	ShowDefaultGroup  bool
}

type ShowResourceGroupsRequest struct {
	Names   []string          `avro:"names"`
	Options map[string]string `avro:"options"`
}

type ShowResourceGroupsResponse struct {
	Groups []map[string]string `avro:"groups"`
	Info   map[string]string   `avro:"info"`
}

// AdminAddHostsRequest --------------------------------------------
type AdminAddHostRequest struct {
	HostAddress string            `avro:"host_address"`
	Options     map[string]string `avro:"options"`
}

type AdminAddHostRequestOptions struct {
	DryRun               bool
	AcceptsFailover      bool
	PublicAddress        string
	HostManagerPublicURL string
	RamLimit             uint64
	GPUs                 []int
}

type AdminAddHostResponse struct {
	addedHost string            `     avro:"added_host"`
	info      map[string]string `avro:"info"`
}

// AdminAddRanksRequestOptions --------------------------------------------
type AdminAddRanksRequestOptions struct {
	DryRun bool `avro:"dry_run"`
}
type AdminAddRanksRequest struct {
	Hosts        []string            `avro:"hosts"`
	ConfigParams []map[string]string `avro:"config_params"`
	Options      map[string]string   `avro:"options"`
}

type AdminAddRanksResponse struct {
	AddedRanks []string          `avro:"added_ranks"`
	Info       map[string]string `avro:"info"`
}

// AdminBackupBegin --------------------------------------------
type AdminBackupBeginRequestOptions struct{}

type AdminBackupBeginRequest struct {
	Options map[string]string `avro:"options"`
}

type AdminBackupBeginResponse struct {
	Info map[string]string `avro:"info"`
}

// AdminBackupEnd --------------------------------------------
type AdminBackupEndRequestOptions struct{}

type AdminBackupEndRequest struct {
	Options map[string]string `avro:"options"`
}

type AdminBackupEndResponse struct {
	Info map[string]string `avro:"info"`
}

// AdminRemoveHostRequest --------------------------------------------
type AdminRemoveHostRequest struct {
	Host    string            `avro:"host"`
	Options map[string]string `avro:"options"`
}

type AdminRemoveHostRequestOptions struct {
	DryRun bool
}

type AdminRemoveHostResponse struct {
	info map[string]string `avro:"info"`
}

// AdminRemoveRanksRequestOptions --------------------------------------------
type AdminRemoveRanksRequestOptions struct {
	RebalanceShardedData   bool
	RebalanceUnshardedData bool
	Aggressiveness         int
}
type AdminRemoveRanksRequest struct {
	Ranks   []string          `avro:"ranks"`
	Options map[string]string `avro:"options"`
}

type AdminRemoveRanksResponse struct {
	RemovedRanks []string          `avro:"removed_ranks"`
	Info         map[string]string `avro:"info"`
}

// AdminShutdown --------------------------------------------------------
type AdminShutdownRequestOptions struct {
}

type AdminShutdownRequest struct {
	ExitType      string            `avro:"exit_type"`
	Authorization string            `avro:"authorization`
	Options       map[string]string `avro:"options"`
}

type AdminShutdownResponse struct {
	ExitStatus string            `avro:"exit_status"`
	Info       map[string]string `avro:"info"`
}

// AggregateStatisticsOptions --------------------------------------------
type AggregateStatisticsOptions struct {
	AdditionalColumnNames string
	WeightColumnName      string
}

type AggregateStatisticsRequest struct {
	TableName  string `avro:"table_name"`
	ColumnName string `avro:"column_name"`
	// Find supported values for Stats here: https://www.kinetica.com/docs/api/rest/aggregate_statistics_rest.html
	Stats   string            `avro:"stats"`
	Options map[string]string `avro:"options"`
}

type AggregateStatisticsResponse struct {
	Stats map[string]float64 `avro:"stats"`
	Info  map[string]string  `avro:"info"`
}

// AggregateGroupByOptions --------------------------------------------
type AggregateGroupByOptions struct {
	Encoding                   string
	FilterExpression           string
	HavingClause               string
	SortOrder                  string
	SortBy                     string
	ResultTable                string
	ResultTablePersist         bool
	ResultTableForceReplicated bool
	ResultTableGeneratePk      bool
	Ttl                        int64
	ChunkSize                  int64
	CreateIndexes              string
	ViewId                     string
	PivotColumn                string
	PivotValues                string
	GroupingSets               string
	Rollup                     string
	Cube                       string
}

type AggregateGroupByRequest struct {
	TableName   string            `avro:"table_name"`
	ColumnNames []string          `avro:"column_names"`
	Offset      int64             `avro:"offset"`
	Limit       int64             `avro:"limit"`
	Encoding    string            `avro:"encoding"`
	Options     map[string]string `avro:"options"`
}

type AggregateGroupByResponse struct {
	ResponseSchema        string            `avro:"response_schema_str"`
	BinaryEncodedResponse []byte            `avro:"binary_encoded_response"`
	JsonEncodedResponse   string            `avro:"json_encoded_response"`
	TotalNumberOfRecords  int64             `avro:"total_number_of_records"`
	HasMoreRecords        bool              `avro:"has_more_records"`
	PagingTable           string            `avro:"paging_table"`
	Info                  map[string]string `avro:"info"`
}

// CreateJobOptions --------------------------------------------
type CreateJobOptions struct {
	RemoveJobOnComplete string
	JobTag              string
}

// RequestEncodingEnum - the encoding of the request payload for the job.
type RequestEncodingEnum string

const (
	RequestEncodingBinary RequestEncodingEnum = "binary"
	RequestEncodingJSON   RequestEncodingEnum = "json"
	RequestEncodingSnappy RequestEncodingEnum = "snappy"
)

func (ree RequestEncodingEnum) String() string {
	return string(ree)
}

type CreateJobRequest struct {
	// Indicates which endpoint to execute, e.g. '/alter/table'.
	Endpoint string `avro:"endpoint"`
	// The encoding of the request payload for the job.
	RequestEncoding RequestEncodingEnum `avro:"request_encoding"`
	// Binary-encoded payload for the job to be run asynchronously.
	// The payload must contain the relevant input parameters for the endpoint indicated in @{input endpoint}.
	// Please see the documentation for the appropriate endpoint to see what values must (or can) be specified.
	// If this parameter is used, then @{input request_encoding} must be {binary}@{choice of input request_encoding}
	// or {snappy}@{choice of input request_encoding}.
	Data []byte `avro:"data"`
	// JSON-encoded payload for the job to be run asynchronously.
	// The payload must contain the relevant input parameters for the endpoint indicated in @{input endpoint}.
	// Please see the documentation for the appropriate endpoint to see what values must (or can) be specified.
	// If this parameter is used, then @{input request_encoding} must be {json}@{choice of input request_encoding}.
	DataStr string            `avro:"data_str"`
	Options map[string]string `avro:"options"`
}

type CreateJobResponse struct {
	// An identifier for the job created by this call.
	JobID int64 `avro:"job_id"`
	// Additional information.
	Info map[string]string `avro:"info"`
}

// GetJobOptions --------------------------------------------
type GetJobOptions struct {
	JobTag string
}

type GetJobRequest struct {
	// An identifier for the job created by this call.
	JobID   int64             `avro:"job_id"`
	Options map[string]string `avro:"options"`
}

// JobStatusEnum - Status of the submitted job
type JobStatusEnum string

const (
	// JobStatusRunning -   The job is currently executing.
	JobStatusRunning JobStatusEnum = "RUNNING"
	// JobStatusDone -      The job execution has successfully completed and the response is included in the output
	//                      parameter job_response or output parameter job_response_str field
	JobStatusDone JobStatusEnum = "DONE"
	// JobStatusError -     The job was attempted, but an error was encountered. The output parameter status_map
	//                      contains the details of the error in error_message
	JobStatusError JobStatusEnum = "ERROR"
	// JobStatusCancelled - Job cancellation was requested while the execution was in progress.
	JobStatusCancelled JobStatusEnum = "CANCELLED"
)

func (jse JobStatusEnum) String() string {
	return string(jse)
}

type GetJobResponse struct {
	// Indicates which endpoint to execute, e.g. '/alter/table'.
	Endpoint string `avro:"endpoint"`
	// Status of the submitted job
	// RUNNING	    The job is currently executing.
	// DONE	        The job execution has successfully completed and the response is included in the output parameter
	//              job_response or output parameter job_response_str field
	// ERROR	    The job was attempted, but an error was encountered. The output parameter status_map
	//              contains the details of the error in error_message
	// CANCELLED	Job cancellation was requested while the execution was in progress.
	JobStatus JobStatusEnum `avro:"job_status"`
	// True if the end point is still executing.
	Running bool `avro:"running"`
	// Approximate percentage of the job completed.
	Progress int `avro:"progress"`
	// True if the job execution completed and no errors were encountered.
	Successful bool `avro:"successful"`
	// The encoding of the job result (contained in output parameter job_response or output parameter job_response_str.
	ResponseEncoding string `avro:"response_encoding"`
	// The binary-encoded response of the job. This field is populated only when the job has completed and
	// output parameter response_encoding is binary
	JobResponse []byte `avro:"job_response"`
	// The json-encoded response of the job. This field is populated only when the job has completed and
	// output parameter response_encoding is json
	JobResponseStr string `avro:"job_response_str"`
	// Map of various status strings for the executed job.
	StatusMap map[string]string `avro:"status_map"`
	// Additional information.
	Info map[string]string `avro:"info"`
}

// AlterGraphRequestOptions --------------------------------------------
type AlterGraphRequestOptions struct {
	// ServerId - Indicates which graph server(s) to send the request to.
	// Default is to send to get information about all the servers.
	ServerID string `avro:"server_id"`
	// BypassClientCheck - Set for non-user requests.
	BypassClientCheck bool `avro:"bypass_client_check"`
}

// AlterGraphActionEnum - Actions that can be used on AlterGraphRequest
type AlterGraphActionEnum string

const (
	// AlterGraphActionAddTableMonitor - Add a table monitor to a graph.
	// The table name is specified as the action argument.
	AlterGraphActionAddTableMonitor AlterGraphActionEnum = "add_table_monitor"
	// AlterGraphActionResetClient - The job execution has successfully completed and the response is included in the
	// output parameter job_response or output parameter job_response_str field
	AlterGraphActionResetClient AlterGraphActionEnum = "reset_client"
	// AlterGraphActionResetServer - Reset all current operations on the server side. This is also sent on (re)start.
	AlterGraphActionResetServer AlterGraphActionEnum = "reset_server"
	// AlterGraphActionCancelTask - Cancel a specific task on the graph server.
	AlterGraphActionCancelTask AlterGraphActionEnum = "cancel_task"
	// AlterGraphActionAlterLogger - Change the server side log level; e.g., 'GraphServer.GraphSolver=DEBUG'
	AlterGraphActionAlterLogger AlterGraphActionEnum = "alter_logger"
	// AlterGraphActionDeleteAll - Delete all graphs, and remove any persistence info.
	AlterGraphActionDeleteAll AlterGraphActionEnum = "delete_all"
	// AlterGraphActionStatus - Current status of the graph client (db side).
	AlterGraphActionStatus AlterGraphActionEnum = "status"
	// AlterGraphActionCollectGraphs - Get the create command for all persisted graphs.
	AlterGraphActionCollectGraphs AlterGraphActionEnum = "collect_graphs"
	// AlterGraphActionRestoreGraphs - Re-creates all graphs from persist info on rank0.
	AlterGraphActionRestoreGraphs AlterGraphActionEnum = "restore_graphs"
)

func (agae AlterGraphActionEnum) String() string {
	return string(agae)
}

type AlterGraphRequest struct {
	// GraphName - Graph on which the operation should be applied.
	// If empty then it will apply to all graphs.
	// This request can be sent from the graph server to the graph client,
	// or from the client to the server depending on the type of operation.
	GraphName string `avro:"graph_name"`
	// Action - Operation to be applied
	Action AlterGraphActionEnum `avro:"action"`
	// ActionArg - Action specific argument.
	ActionArg string            `avro:"action_arg"`
	Options   map[string]string `avro:"options"`
}

type AlterGraphResponse struct {
	// Action - Operation to be applied
	Action AlterGraphActionEnum `avro:"action"`
	// ActionArg - Action specific argument.
	ActionArg string `avro:"action_arg"`
	// Info -
	Info map[string]string `avro:"info"`
}

// AlterUserRequest --------------------------------------------
// RequestEncodingEnum - the encoding of the request payload for the job.
type AlterUserActionEnum string

const (
	AlterUserActionSetPassword      AlterUserActionEnum = "set_password"
	AlterUserActionSetResourceGroup AlterUserActionEnum = "set_resource_group"
	AlterUserActionSetDefaultSchema AlterUserActionEnum = "set_default_schema"
)

func (auae AlterUserActionEnum) String() string {
	return string(auae)
}

type AlterUserRequest struct {
	// Name of the user to be altered. Must be an existing user.
	Name string `avro:"name"`
	// Modification operation to be applied to the user.
	Action AlterUserActionEnum `avro:"action"`
	// The value of the modification, depending on input parameter action.
	Value string `avro:"value"`
	// The value of the modification, depending on input parameter action.
	Options map[string]string `avro:"options"`
}

type AlterUserResponse struct{}

type AlterUserRequestOptions struct{}

// CreateSchemaRequest --------------------------------------------
type CreateSchemaRequest struct {
	Name    string            `avro:"schema_name"`
	Options map[string]string `avro:"options"`
}

type CreateSchemaResponse struct {
	SchemaName string            `avro:"schema_name"`
	Info       map[string]string `avro:"info"`
}

type DropSchemaRequest struct {
	Name    string            `avro:"schema_name"`
	Options map[string]string `avro:"options"`
}
type DropSchemaRequestOptions struct {
	NoErrorIfNotExists bool `avro:"no_error_if_not_exists"`
	Cascade            bool `avro:"cascade"`
}
type DropSchemaResponse struct {
	Name string            `avro:"schema_name"`
	Info map[string]string `avro:"info"`
}
type CreateSchemaRequestOptions struct {
	// NoErrorIfExists - If true, prevents an error from occurring if the schema already exists.
	// The default value is false. The supported values are:
	//
	// true
	// false
	NoErrorIfExists bool `avro:"no_error_if_exists"`
}

// CreateResourceGroupRequest --------------------------------------------

type CreateResourceGroupRankingEnum string

const (
	CreateResourceGroupRankingFirst  CreateResourceGroupRankingEnum = "first"
	CreateResourceGroupRankingLast   CreateResourceGroupRankingEnum = "last"
	CreateResourceGroupRankingBefore CreateResourceGroupRankingEnum = "before"
	CreateResourceGroupRankingAfter  CreateResourceGroupRankingEnum = "after"
)

func (crgre CreateResourceGroupRankingEnum) String() string {
	return string(crgre)
}

type CreateResourceGroupRequest struct {
	// Name -Name of the group to be created.
	// Must contain only letters, digits, and underscores, and cannot begin with a digit.
	// Must not match existing resource group name.
	Name string `avro:"name"`
	// TierAttributes - Optional map containing tier names and their respective attribute group limits.
	// The only valid attribute limit that can be set is max_memory (in bytes) for the VRAM & RAM tiers.
	// For instance, to set max VRAM capacity to 1GB and max RAM capacity to 10GB,
	//
	// use: {'VRAM':{'max_memory':'1000000000'}, 'RAM':{'max_memory':'10000000000'}}.
	//
	// The default value is an empty map ( {} ).
	TierAttributes map[string]map[string]string `avro:"tier_attributes"`
	// Ranking - Indicates the relative ranking among existing resource groups where this
	// new resource group will be placed. When using before or after,
	// specify which resource group this one will be inserted before or after in input parameter
	// adjoining_resource_group.
	Ranking CreateResourceGroupRankingEnum `avro:"ranking"`
	// AdjoiningResourceGroup - Indicates the relative ranking among existing resource groups where this
	// new resource group will be placed.
	// When using before or after, specify which resource group this one will be inserted before or after
	// in input parameter adjoining_resource_group. The supported values are:
	//
	// first
	// last
	// before
	// after
	AdjoiningResourceGroup string            `avro:"adjoining_resource_group"`
	Options                map[string]string `avro:"options"`
}

type CreateResourceGroupRequestOptions struct {
	// MaxCpuConcurrency - Maximum number of simultaneous threads that will be used to execute a request
	// for this group.
	// The minimum allowed value is 4.
	MaxCpuConcurrency string `avro:"max_cpu_concurrency"`
	// MaxData - Maximum amount of cumulative ram usage regardless of tier status for this group.
	// The minimum allowed value is -1.
	MaxData string `avro:"max_data"`
	// MaxSchedulingPriority - Maximum priority of a scheduled task for this group.
	// The minimum allowed value is 1.
	// The maximum allowed value is 100.
	MaxSchedulingPriority string `avro:"max_scheduling_priority"`
	// MaxTierPriority - Maximum priority of a tiered object for this group.
	// The minimum allowed value is 1.
	// The maximum allowed value is 10.
	MaxTierPriority string `avro:"max_tier_priority"`
}

type CreateResourceGroupResponse struct {
	SchemaName string            `avro:"schema_name"`
	Info       map[string]string `avro:"info"`
}
type DeleteResourceGroupRequest struct {
	Name    string            `avro:"name"`
	Options map[string]string `avro:"options"`
}
type DeleteResourceGroupRequestOptions struct {
	Cascade bool `avro:"cascade_delete"`
}

type DeleteResourceGroupResponse struct {
	Name string            `avro:"name"`
	Info map[string]string `avro:"info"`
}
