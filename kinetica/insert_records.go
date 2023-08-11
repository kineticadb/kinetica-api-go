package kinetica

import (
	"context"
	"strconv"

	"github.com/hamba/avro/v2"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/multierr"
)

// NewDefaultInsertRecordsOptions
//
//	@return *InsertRecordsOptions
func NewDefaultInsertRecordsOptions() *InsertRecordsOptions {
	return &InsertRecordsOptions{
		UpdateOnExistingPk:     false,
		IgnoreExistingPk:       false,
		ReturnRecordIDs:        false,
		TruncateStrings:        false,
		ReturnIndividualErrors: false,
		AllowPartialBatch:      false,
	}
}

// InsertRecordsRaw //
//
//	@receiver kinetica
//	@param ctx
//	@param table
//	@param data
//	@return *InsertRecordsResponse
//	@return error
func (kinetica *Kinetica) InsertRecordsRaw(
	ctx context.Context,
	table string, data []interface{}) (*InsertRecordsResponse, error) {
	var (
		childCtx  context.Context
		childSpan trace.Span
	)

	childCtx, childSpan = kinetica.tracer.Start(ctx, "kinetica.InsertRecordsRaw()")
	defer childSpan.End()

	return kinetica.InsertRecordsRawWithOpts(childCtx, table, data, NewDefaultInsertRecordsOptions())
}

// InsertRecordsRawWithOpts //
//
//	@receiver kinetica
//	@param ctx
//	@param table
//	@param data
//	@param options
//	@return *InsertRecordsResponse
//	@return error
func (kinetica *Kinetica) InsertRecordsRawWithOpts(
	ctx context.Context,
	table string, data []interface{}, options *InsertRecordsOptions) (*InsertRecordsResponse, error) {
	var (
		childCtx  context.Context
		childSpan trace.Span
		errors    []error
		// insertRecordsMutex *sync.Mutex
	)

	childCtx, childSpan = kinetica.tracer.Start(ctx, "kinetica.InsertRecordsRawWithOpts()")
	defer childSpan.End()
	// insertRecordsMutex = &sync.Mutex{}

	showTableResult, err := kinetica.ShowTableRawWithOpts(context.TODO(), table, &ShowTableOptions{
		ForceSynchronous:   true,
		GetSizes:           false,
		ShowChildren:       false, // needs to be false for tables
		NoErrorIfNotExists: false,
		GetColumnInfo:      true,
	})
	if err != nil {
		return nil, err
	}

	typeSchema := showTableResult.TypeSchemas[0]
	recordSchema, recordErr := avro.Parse(typeSchema)
	if recordErr != nil {
		return nil, err
	}

	// Convert the records into a byte array according to the Avro schema
	buffer := make([][]byte, len(data))
	for i := 0; i < len(data); i++ {
		buf, err := avro.Marshal(recordSchema, data[i])
		if err != nil {
			errors = append(errors, err)
		}
		buffer[i] = buf
	}

	// insertRecordsMutex.Lock()
	mapOptions := kinetica.buildInsertRecordsOptionsMap(childCtx, options)

	response := InsertRecordsResponse{}
	request := InsertRecordsRequest{
		TableName:    table,
		List:         buffer,
		ListString:   [][]string{},
		ListEncoding: "binary",
		Options:      *mapOptions,
	}
	// insertRecordsMutex.Unlock()

	err = kinetica.submitRawRequest(
		childCtx, "/insert/records",
		&Schemas.insertRecordsRequest, &Schemas.insertRecordsResponse,
		&request, &response)
	if err != nil {
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return nil, multierr.Combine(errors...)
	}
	return &response, nil
}

func (kinetica *Kinetica) buildInsertRecordsOptionsMap(ctx context.Context, options *InsertRecordsOptions) *map[string]string {
	var (
		// childCtx context.Context
		childSpan trace.Span
	)

	_, childSpan = kinetica.tracer.Start(ctx, "kinetica.buildInsertRecordsOptionsMap()")
	defer childSpan.End()

	mapOptions := make(map[string]string)

	mapOptions["update_on_existing_pk"] = strconv.FormatBool(options.UpdateOnExistingPk)
	mapOptions["ignore_existing_pk"] = strconv.FormatBool(options.IgnoreExistingPk)
	mapOptions["return_record_ids"] = strconv.FormatBool(options.ReturnRecordIDs)
	mapOptions["truncate_strings"] = strconv.FormatBool(options.TruncateStrings)
	mapOptions["return_individual_errors"] = strconv.FormatBool(options.ReturnIndividualErrors)
	mapOptions["allow_partial_batch"] = strconv.FormatBool(options.AllowPartialBatch)

	return &mapOptions
}

// InsertRecordsMap //
//
//	@receiver kinetica
//	@param ctx
//	@param table
//	@param data
//	@return *InsertRecordsResponse
//	@return error
func (kinetica *Kinetica) InsertRecordsMap(
	ctx context.Context,
	table string, data []interface{}) (*InsertRecordsResponse, error) {
	var (
		childCtx  context.Context
		childSpan trace.Span
	)

	childCtx, childSpan = kinetica.tracer.Start(ctx, "kinetica.GetRecordsMap()")
	defer childSpan.End()

	return kinetica.InsertRecordsMapWithOpts(childCtx, table, data, NewDefaultInsertRecordsOptions())
}

// InsertRecordsMapWithOpts
//
//	@receiver kinetica
//	@param ctx
//	@param table
//	@param data
//	@param options
//	@return *InsertRecordsResponse
//	@return error
func (kinetica *Kinetica) InsertRecordsMapWithOpts(
	ctx context.Context,
	table string, data []interface{}, options *InsertRecordsOptions) (*InsertRecordsResponse, error) {
	var (
		childCtx  context.Context
		childSpan trace.Span
	)

	childCtx, childSpan = kinetica.tracer.Start(ctx, "kinetica.InsertRecordsMapWithOpts()")
	defer childSpan.End()

	raw, err := kinetica.InsertRecordsRawWithOpts(childCtx, table, data, options)
	if err != nil {
		return nil, err
	}
	// start := time.Now()
	return raw, nil
	// TODO
}

// InsertRecordsStruct
//
//	@receiver kinetica
//	@param ctx
//	@param table
//	@param data
//	@return *InsertRecordsResponse
//	@return error
func (kinetica *Kinetica) InsertRecordsStruct(
	ctx context.Context,
	table string, data []interface{}) (*InsertRecordsResponse, error) {
	var (
		childCtx  context.Context
		childSpan trace.Span
	)

	childCtx, childSpan = kinetica.tracer.Start(ctx, "kinetica.InsertRecordsStruct()")
	defer childSpan.End()

	return kinetica.InsertRecordsStructWithOpts(childCtx, table, data, NewDefaultInsertRecordsOptions())
}

// InsertRecordsStructWithOpts
//
//	@receiver kinetica
//	@param ctx
//	@param table
//	@param data
//	@param options
//	@return *InsertRecordsResponse
//	@return error
func (kinetica *Kinetica) InsertRecordsStructWithOpts(
	ctx context.Context,
	table string, data []interface{}, options *InsertRecordsOptions) (*InsertRecordsResponse, error) {
	var (
		childCtx  context.Context
		childSpan trace.Span
	)

	childCtx, childSpan = kinetica.tracer.Start(ctx, "kinetica.InsertRecordsStructWithOpts()")
	defer childSpan.End()

	response, err := kinetica.InsertRecordsMapWithOpts(childCtx, table, data, options)
	if err != nil {
		return nil, err
	}

	return response, nil
}
