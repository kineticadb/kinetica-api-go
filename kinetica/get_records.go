package kinetica

import (
	"context"
	"log"
	"strconv"
	"sync"

	"github.com/hamba/avro/v2"
	"github.com/mitchellh/mapstructure"
	"go.opentelemetry.io/otel/trace"
)

func NewDefaultGetRecordsOptions() *GetRecordsOptions {
	return &GetRecordsOptions{
		Expression: "",
		// FastIndexLookup: false,
		SortBy:    "",
		SortOrder: "ascending",
		Encoding:  "binary",
	}
}

func (gpudb *Kinetica) GetRecordsRaw(
	ctx context.Context,
	table string, offset int64, limit int64) (*GetRecordsResponse, error) {
	var (
		childCtx  context.Context
		childSpan trace.Span
	)

	childCtx, childSpan = gpudb.tracer.Start(ctx, "gpudb.GetRecordsRaw()")
	defer childSpan.End()

	return gpudb.GetRecordsRawWithOpts(childCtx, table, offset, limit, NewDefaultGetRecordsOptions())
}

func (gpudb *Kinetica) GetRecordsRawWithOpts(
	ctx context.Context,
	table string, offset int64, limit int64, options *GetRecordsOptions) (*GetRecordsResponse, error) {
	var (
		childCtx  context.Context
		childSpan trace.Span
	)

	childCtx, childSpan = gpudb.tracer.Start(ctx, "gpudb.GetRecordsRawWithOpts()")
	defer childSpan.End()

	mapOptions := gpudb.buildGetRecordsOptionsMap(childCtx, options)
	response := GetRecordsResponse{}
	request := GetRecordsRequest{TableName: table, Offset: offset, Limit: limit, Encoding: options.Encoding, Options: *mapOptions}
	err := gpudb.submitRawRequest(
		childCtx, "/get/records",
		&Schemas.getRecordsRequest, &Schemas.getRecordsResponse,
		&request, &response)

	return &response, err
}

func (gpudb *Kinetica) buildGetRecordsOptionsMap(ctx context.Context, options *GetRecordsOptions) *map[string]string {
	var (
		childSpan trace.Span
	)

	_, childSpan = gpudb.tracer.Start(ctx, "gpudb.buildGetRecordsOptionsMap()")
	defer childSpan.End()

	mapOptions := make(map[string]string)
	if options.Expression != "" {
		mapOptions["expression"] = options.Expression
	}

	if options.Expression != "" {
		mapOptions["fast_index_lookup"] = strconv.FormatBool(options.FastIndexLookup)
	}

	if options.SortBy != "" {
		mapOptions["sort_by"] = options.SortBy
		if options.SortBy != "" {
			mapOptions["sort_order"] = options.SortOrder
		}
	}

	return &mapOptions
}

func (gpudb *Kinetica) GetRecordsMap(
	ctx context.Context,
	table string, offset int64, limit int64) (*[]map[string]interface{}, error) {
	var (
		childCtx  context.Context
		childSpan trace.Span
	)

	childCtx, childSpan = gpudb.tracer.Start(ctx, "gpudb.GetRecordsMap()")
	defer childSpan.End()

	return gpudb.GetRecordsMapWithOpts(childCtx, table, offset, limit, NewDefaultGetRecordsOptions())
}

func (gpudb *Kinetica) GetRecordsMapWithOpts(
	ctx context.Context,
	table string, offset int64, limit int64, options *GetRecordsOptions) (*[]map[string]interface{}, error) {
	var (
		childCtx  context.Context
		childSpan trace.Span
	)

	childCtx, childSpan = gpudb.tracer.Start(ctx, "gpudb.GetRecordsMapWithOpts()")
	defer childSpan.End()

	raw, err := gpudb.GetRecordsRawWithOpts(childCtx, table, offset, limit, options)
	if err != nil {
		return nil, err
	}
	// start := time.Now()
	recordSchema, recordErr := avro.Parse(raw.TypeSchema)
	if recordErr != nil {
		return nil, recordErr
	}

	if options.Encoding == "binary" {
		resultList := make([]map[string]interface{}, len(raw.RecordsBinary))
		for i, value := range raw.RecordsBinary {
			avro.Unmarshal(recordSchema, value, &resultList[i])
		}

		return &resultList, nil
		// duration := time.Since(start)
		// fmt.Println("GetRecordsMap", duration.Milliseconds(), " ns")
	}

	// TODO
	panic("JSON decoding is not yet implemented")
}

func (gpudb *Kinetica) GetRecordsStruct(
	ctx context.Context,
	table string, offset int64, limit int64, newInstance func() interface{}) (*[]interface{}, error) {
	var (
		childCtx  context.Context
		childSpan trace.Span
	)

	childCtx, childSpan = gpudb.tracer.Start(ctx, "gpudb.GetRecordsStruct()")
	defer childSpan.End()

	return gpudb.GetRecordsStructWithOpts(childCtx, table, offset, limit, NewDefaultGetRecordsOptions(), newInstance)
}

func (gpudb *Kinetica) GetRecordsStructWithOpts(
	ctx context.Context,
	table string, offset int64, limit int64, options *GetRecordsOptions, newInstance func() interface{}) (*[]interface{}, error) {
	var (
		childCtx  context.Context
		childSpan trace.Span
	)

	childCtx, childSpan = gpudb.tracer.Start(ctx, "gpudb.GetRecordsStructWithOpts()")
	defer childSpan.End()

	recordsMap, err := gpudb.GetRecordsMapWithOpts(childCtx, table, offset, limit, options)
	if err != nil {
		return nil, err
	}
	// start := time.Now()
	resultList := make([]interface{}, len(*recordsMap))

	for i, valueMap := range *recordsMap {
		newInst := newInstance()
		resultList[i] = newInst
		err = mapstructure.Decode(valueMap, &resultList[i])

		if err != nil {
			return nil, err
		}
	}
	// duration := time.Since(start)
	// fmt.Println("GetRecordsStruct", duration.Milliseconds(), " ns")

	return &resultList, nil
}

// GetRecordsStructWithOpts - GetRecords generic variant to work with any mapstructure tagged struct
//
//	@param ctx
//	@param gpudb
//	@param table
//	@param offset
//	@param limit
//	@param options
//	@return []T
//	@return error
func GetRecordsStructWithOpts[T any](
	ctx context.Context,
	gpudb Kinetica, table string, offset int64, limit int64, options *GetRecordsOptions, newInstance func() T) ([]T, error) {
	var (
		childCtx   context.Context
		childSpan  trace.Span
		resultList []T
	)

	childCtx, childSpan = gpudb.tracer.Start(ctx, "gpudb.GetRecordsStructWithOpts()")
	defer childSpan.End()

	recordsMap, err := gpudb.GetRecordsMapWithOpts(childCtx, table, offset, limit, options)
	if err != nil {
		return nil, err
	}

	resultList = make([]T, len(*recordsMap))

	for i, valueMap := range *recordsMap {
		newInst := newInstance()
		resultList[i] = newInst
		err = mapstructure.Decode(valueMap, &resultList[i])

		if err != nil {
			return nil, err
		}
	}
	return resultList, nil
}

func (gpudb *Kinetica) GetRecordsStructSendChannel(
	ctx context.Context,
	table string, offset int64, limit int64, resultChannel chan interface{},
	newInstance func() interface{}) (*GetRecordsResponse, error) {
	var (
		childCtx  context.Context
		childSpan trace.Span
	)

	childCtx, childSpan = gpudb.tracer.Start(ctx, "gpudb.GetRecordsStructSendChannel()")
	defer childSpan.End()

	return gpudb.GetRecordsStructSendChannelWithOpts(
		childCtx, table, offset, limit, NewDefaultGetRecordsOptions(), resultChannel, newInstance)
}

func (gpudb *Kinetica) GetRecordsStructSendChannelWithOpts(
	ctx context.Context, table string, offset int64, limit int64, options *GetRecordsOptions,
	resultChannel chan interface{}, newInstance func() interface{}) (*GetRecordsResponse, error) {
	var (
		allNodesWaitGroup sync.WaitGroup
		childCtx          context.Context
		childSpan         trace.Span
	)

	childCtx, childSpan = gpudb.tracer.Start(ctx, "gpudb.GetRecordsStructSendChannelWithOpts()")
	defer childSpan.End()

	raw, err := gpudb.GetRecordsRawWithOpts(childCtx, table, offset, limit, options)
	if err != nil {
		return nil, err
	}

	recordSchema, _ := avro.Parse(raw.TypeSchema)
	// start := time.Now()

	for _, value := range raw.RecordsBinary {
		allNodesWaitGroup.Add(1)

		go func() {
			newMap := make(map[string]interface{})
			avro.Unmarshal(recordSchema, value, &newMap)

			newInst := newInstance()
			err = mapstructure.Decode(newMap, &newInst)

			if err != nil {
				log.Fatalln(err)
			}
			resultChannel <- newInst

			allNodesWaitGroup.Done()
		}()
	}

	go func() {
		allNodesWaitGroup.Wait()
		close(resultChannel)
		// duration := time.Since(start)
		// fmt.Println("GetRecordsStruct", duration.Milliseconds(), " ms")
	}()

	return raw, nil
}
