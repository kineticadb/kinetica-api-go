package main

import (
	"context"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/kineticadb/kinetica-api-go/example"
	"github.com/kineticadb/kinetica-api-go/kinetica"
	"github.com/gocarina/gocsv"
	"go.uber.org/multierr"
	"go.uber.org/zap"
)

var (
	pagingTable string
)

func main() {
	endpoint := "http://127.0.0.1:9191" //os.Args[1]
	username := ""                      //os.Args[2]
	password := ""                      //os.Args[3]

	// Logger, err := zap.NewProduction()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	ctx := context.TODO()
	options := kinetica.KineticaOptions{Username: username, Password: password}
	// fmt.Println("Options", options)
	dbInst := kinetica.NewWithOptions(ctx, endpoint, &options)

	var anyValue []any
	intSlice := []int{1, 2, 3, 4}
	anyValue = append(anyValue, intSlice)

	fmt.Println(anyValue...)

	// runShowExplainVerboseAnalyseSqlStatement(dbInst)
	// runShowExplainVerboseSqlStatement(dbInst)
	// runShowStoredProcedureDDL(dbInst)
	// runShowTableDDL(dbInst)
	// runShowTableResourcesBySchema1(dbInst)
	// runAggregateGroupBy1(dbInst)
	// runAggregateGroupBy2(dbInst)
	// runAggregateGroupBy3(dbInst)
	// runAggregateStatistics1(dbInst)
	// runAggregateStatistics2(dbInst)
	// runShowResourceGroups1(dbInst)
	// runShowSecurity1(dbInst)
	// runShowResourceStatistics1(dbInst)
	// runShowResourceStatistics2(dbInst)
	// runShowResourceStatistics3(dbInst)
	// runShowSqlProc1(dbInst)
	// runShowSqlProc2(dbInst)
	// runShowSystemStatus1(dbInst)
	// runShowSystemTiming1(dbInst)
	// runShowSystemProperties1(dbInst)
	// runExecuteSql1(dbInst)
	pagingTable = runExecuteSql2(dbInst)
	// runExecuteSql3(dbInst)
	// runExecuteSql4(dbInst)
	// runExecuteSql5(dbInst)
	// runExecuteSql6(dbInst)
	// runExecuteSql7(dbInst)
	// runExecuteSql8(dbInst)
	// runExecuteSql9(dbInst)
	// runShowSchema1(dbInst)
	// runShowTable1(dbInst)
	// runGetRecords1(dbInst)
	// runGetRecords2(dbInst)
	// runGetRecords3(dbInst)
	runGetRecords4(dbInst)
	runGetRecords5(dbInst, pagingTable)
	// runInsertRecords(Logger, dbInst)
	// runCreateResourceGroup(dbInst, "lucid_test")
	// runDeleteResourceGroup(dbInst, "lucid_test")
	// runCreateSchema(dbInst, "lucid_test")
	// runDropSchema(dbInst, "lucid_test")

}

func runShowExplainVerboseAnalyseSqlStatement(dbInst *kinetica.Kinetica) {
	start := time.Now()
	// string field
	result, err := dbInst.ShowExplainVerboseAnalyseSqlStatement(context.TODO(), `select * from 
(select vendor_id, payment_type, sum(fare_amount) as sum_fare1 from demo.nyctaxi group by vendor_id, payment_type) a
inner join
(select vendor_id, rate_code_id, sum(trip_distance) as sum_trip1 from demo.nyctaxi group by vendor_id, rate_code_id) b
on a.vendor_id = b.vendor_id
`)
	if err != nil {
		panic(err)
	}
	duration := time.Since(start)
	//fmt.Println(*result)
	for _, plan := range *result.Plans {
		fmt.Println(plan)
		requestMap, _ := plan.JsonRequestMap()
		fmt.Println("JSON", requestMap)
	}

	for _, plan := range *result.Plans {
		fmt.Println(plan)
		depPlans, _ := plan.FindDependentPlans()
		fmt.Println("DependentPlans", depPlans)
	}
	fmt.Println("ShowExplainVerboseAnalyseSqlStatement", duration.Milliseconds(), " ms")
}

func runShowExplainVerboseSqlStatement(dbInst *kinetica.Kinetica) {
	start := time.Now()
	// string field
	result, err := dbInst.ShowExplainVerboseSqlStatement(context.TODO(), "select count(*) from demo.nyctaxi")
	if err != nil {
		panic(err)
	}
	duration := time.Since(start)
	fmt.Println(*result)
	fmt.Println("ShowExplainVerboseSqlStatement", duration.Milliseconds(), " ms")
}

func runShowStoredProcedureDDL(dbInst *kinetica.Kinetica) {
	start := time.Now()
	// string field
	result, err := dbInst.ShowStoredProcedureDDL(context.TODO(), "OPENSKY.opensky_processing")
	if err != nil {
		panic(err)
	}
	duration := time.Since(start)
	fmt.Println(*result)
	fmt.Println("ShowStoredProcedureDDL", duration.Milliseconds(), " ms")
}

func runShowTableDDL(dbInst *kinetica.Kinetica) {
	start := time.Now()
	// string field
	result, err := dbInst.ShowTableDDL(context.TODO(), "demo.nyctaxi_shard")
	if err != nil {
		panic(err)
	}
	duration := time.Since(start)
	fmt.Println(*result)
	fmt.Println("ShowTableDDL", duration.Milliseconds(), " ms")
}

func runShowTable1(dbInst *kinetica.Kinetica) {
	start := time.Now()
	//result, err := dbInst.ShowTableRaw("MASTER.nyctaxi")
	result, err := dbInst.ShowTableRawWithOpts(context.TODO(), "otel.trace_span", &kinetica.ShowTableOptions{
		ForceSynchronous:   true,
		GetSizes:           true,
		ShowChildren:       false, // needs to be false for tables
		NoErrorIfNotExists: false,
		GetColumnInfo:      true,
	})
	if err != nil {
		panic(err)
	}
	duration := time.Since(start)
	fmt.Println(result)
	fmt.Println("ShowTable", duration.Milliseconds(), " ms")
}

func runShowSystemStatus1(dbInst *kinetica.Kinetica) {
	start := time.Now()
	result, err := dbInst.ShowSystemStatusRaw(context.TODO())
	if err != nil {
		panic(err)
	}
	duration := time.Since(start)
	fmt.Println(*result)
	fmt.Println("ShowSystemStatusRaw", duration.Milliseconds(), " ms")
}

func runShowSystemProperties1(dbInst *kinetica.Kinetica) {
	start := time.Now()
	result, err := dbInst.ShowSystemPropertiesRaw(context.TODO())
	if err != nil {
		panic(err)
	}
	duration := time.Since(start)
	fmt.Println(*result)
	fmt.Println("ShowSystemPropertiesRaw", duration.Milliseconds(), " ms")
}

// Insert SQL test - Works
func runExecuteSql4(dbInst *kinetica.Kinetica) {
	start := time.Now()
	sql := `INSERT INTO otel.trace_span
	(id, resource_id, scope_id, event_id, link_id, trace_id, span_id, parent_span_id, trace_state, name, span_kind, start_time_unix_nano, end_time_unix_nano, dropped_attributes_count, dropped_events_count, dropped_links_count, message, status_code)
	VALUES('3f02a130-726f-4fda-a115-18a12e7c5884', 'ba3bd6c7-afdf-43aa-9287-375b956720db', 'b4006169-85d3-4874-bcff-49f2af13d9a0', '6d820f15-99f2-4e05-803c-e80c8be8748a', '56e36b10-a184-4fc1-8aa6-d3105e954972', '5b8aa5a2d2c872e8321cf37308d69df2', '5fb397be34d26b51', '051581bf3cb55c13', '', 'Hello-Greetings', 0, 1651258378114000000, 1651272778114000000, 0, 0, 0, '', 0)`
	result, err := dbInst.ExecuteSqlRaw(context.TODO(), sql, 0, 0, "", nil)
	if err != nil {
		panic(err)
	}
	duration := time.Since(start)
	fmt.Println("Records inserted =", result.CountAffected)
	fmt.Println("ExecuteSqlRaw -Insert", duration.Milliseconds(), " ms")
}

func runExecuteSql5(dbInst *kinetica.Kinetica) {
	start := time.Now()

	result, err := kinetica.ExecuteSQLStructWithOpts[example.Weather](context.TODO(), *dbInst, "select * from ki_home.weather", 0, 10, kinetica.NewDefaultExecuteSqlOptions())
	if err != nil {
		panic(err)
	}
	duration := time.Since(start)
	fmt.Printf("%+v\n", result)

	fmt.Println("Generic ExecuteSqlStructWithOpts", duration.Milliseconds(), " ms")
}

func runExecuteSql6(dbInst *kinetica.Kinetica) {
	start := time.Now()

	result, err := kinetica.ExecuteSQLStructWithOpts[example.WeatherSummary](context.TODO(), *dbInst, "select max(x), max(avg_temp) from ki_home.weather;", 0, 10, kinetica.NewDefaultExecuteSqlOptions())
	if err != nil {
		panic(err)
	}
	duration := time.Since(start)
	fmt.Printf("%+v\n", result)

	fmt.Println("Generic ExecuteSqlStructWithOpts aggregate functions", duration.Milliseconds(), " ms")
}

// This fails
func runExecuteSql3(dbInst *kinetica.Kinetica) {
	start := time.Now()

	result, err := dbInst.ExecuteSqlStruct(context.TODO(), "select * from ki_home.weather", 0, 10, func() interface{} { return example.Weather{} })
	if err != nil {
		panic(err)
	}
	duration := time.Since(start)
	fmt.Printf("%+v\n", *result.ResultsStruct)

	for i := 0; i < len(*result.ResultsStruct); i++ {
		weather := (*result.ResultsStruct)[i].(example.Weather)
		fmt.Println(weather)
	}
	fmt.Println("ExecuteSqlStruct", duration.Milliseconds(), " ms")
}

func runExecuteSql2(dbInst *kinetica.Kinetica) string {
	start := time.Now()
	result, err := dbInst.ExecuteSqlMap(context.TODO(), "select max(x), max(avg_temp) from ki_home.weather", 0, 0)
	if err != nil {
		panic(err)
	}
	duration := time.Since(start)
	fmt.Println(*result.ResultsMap)
	fmt.Println("ExecuteSqlMap", duration.Milliseconds(), " ms")
	return result.PagingTable
}

// This works
func runExecuteSql1(dbInst *kinetica.Kinetica) {
	start := time.Now()
	result, err := dbInst.ExecuteSqlRaw(context.TODO(), "select max(trace_id) as max_trace_id from ki_home.sample", 0, 0, "", nil)
	if err != nil {
		panic(err)
	}
	duration := time.Since(start)
	fmt.Println(result)
	fmt.Println("ExecuteSqlRaw", duration.Milliseconds(), " ms")
}

// runExecuteSql7 - /has/schema
func runExecuteSql7(dbInst *kinetica.Kinetica) {
	start := time.Now()
	result, err := dbInst.ExecuteSqlRaw(context.TODO(), "execute endpoint '/has/schema' JSON '{\"schema_name\":\"otel1\"}';", 0, 0, "", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
	duration := time.Since(start)
	fmt.Println(result)
	fmt.Println("ExecuteSqlRaw - /has/schema - ", duration.Milliseconds(), " ms")
}

// runExecuteSql8 - /create/schema for existing schema
//
//	@param dbInst
func runExecuteSql8(dbInst *kinetica.Kinetica) {
	start := time.Now()
	result, err := dbInst.ExecuteSqlRaw(context.TODO(), "execute endpoint '/create/schema' JSON '{\"schema_name\":\"otel\"}';", 0, 0, "", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
	duration := time.Since(start)
	fmt.Println(result)
	fmt.Println(result.Info)
	fmt.Println("ExecuteSqlRaw - /create/schema - ", duration.Milliseconds(), " ms")
}

func runExecuteSql9(dbInst *kinetica.Kinetica) {
	start := time.Now()
	result, err := dbInst.ExecuteSqlRaw(context.TODO(), "execute endpoint '/has/table' JSON '{\"table_name\":\"otel.metric_gauge_1\"}';", 0, 0, "", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
	duration := time.Since(start)
	fmt.Println(result)
	fmt.Println(result.Info)
	fmt.Println("ExecuteSqlRaw - /has/table - ", duration.Milliseconds(), " ms")
}

func runGetRecords1(dbInst *kinetica.Kinetica) {
	start := time.Now()
	result, err := dbInst.GetRecordsStruct(context.TODO(), "ki_home.weather", 0, 10, func() interface{} { return example.Weather{} })
	if err != nil {
		panic(err)
	}
	duration := time.Since(start)
	fmt.Println(result)
	fmt.Println("GetRecordsStruct", duration.Milliseconds(), " ms")
}

func runGetRecords2(dbInst *kinetica.Kinetica) {
	messages := make(chan interface{}, 1000)

	start := time.Now()
	dbInst.GetRecordsStructSendChannel(context.TODO(), "MASTER.nyctaxi", 0, 1000, messages, func() interface{} { return example.NycTaxi{} })
	for elem := range messages {
		noop(elem)
	}
	duration := time.Since(start)
	fmt.Println("GetRecordsStruct", duration.Milliseconds(), " ms")
}

func runGetRecords3(dbInst *kinetica.Kinetica) {

	start := time.Now()
	response, err := dbInst.GetRecordsRaw(context.TODO(), "otel.trace_span", 0, 50)
	if err != nil {
		panic(err)
	}
	duration := time.Since(start)
	fmt.Println("GetRecordsRaw", duration.Milliseconds(), " ms")
	fmt.Println(response)
}

func runGetRecords4(dbInst *kinetica.Kinetica) {
	start := time.Now()
	result, err := kinetica.GetRecordsStructWithOpts[example.Weather](context.TODO(), *dbInst, "ki_home.weather", 0, 10, kinetica.NewDefaultGetRecordsOptions(), func() example.Weather { return *new(example.Weather) })
	if err != nil {
		panic(err)
	}
	duration := time.Since(start)
	fmt.Println(result)
	fmt.Println("GetRecordsStruct", duration.Milliseconds(), " ms")
}

func runGetRecords5(dbInst *kinetica.Kinetica, pagingTable string) {
	start := time.Now()
	result, err := kinetica.GetRecordsStructWithOpts[example.WeatherSummary](context.TODO(), *dbInst, pagingTable, 0, 10, kinetica.NewDefaultGetRecordsOptions(), func() example.WeatherSummary { return *new(example.WeatherSummary) })
	if err != nil {
		fmt.Println(err.Error())
	}
	duration := time.Since(start)
	fmt.Println(result)
	fmt.Println("GetRecordsStruct", duration.Milliseconds(), " ms")
}

// ChunkBySize - Splits a slice into multiple slices of the given size
//
//	@param items
//	@param chunkSize
//	@return [][]T
func ChunkBySize[T any](items []T, chunkSize int) [][]T {
	var _chunks = make([][]T, 0, (len(items)/chunkSize)+1)
	for chunkSize < len(items) {
		items, _chunks = items[chunkSize:], append(_chunks, items[0:chunkSize:chunkSize])
	}
	return append(_chunks, items)
}

func runInsertRecords(logger *zap.Logger, dbInst *kinetica.Kinetica) {
	in, err := os.Open("./example/trace_span.csv")
	if err != nil {
		panic(err)
	}
	defer in.Close()

	spans := []example.Span{}

	if err := gocsv.UnmarshalFile(in, &spans); err != nil {
		panic(err)
	}

	// conversion from []example.Span to []interface{} is necessary
	// so that InsertRecords can deal with any structure and not
	// just a specific one
	spanRecords := make([]interface{}, len(spans))
	spanAttributeRecords := make([]interface{}, len(spans))

	for i := 0; i < len(spans); i++ {
		spanRecords[i] = spans[i]
		attribValue := example.AttributeValue{
			IntValue:    i*100 + 1,
			StringValue: "",
			BoolValue:   0,
			DoubleValue: 0,
			BytesValue:  []byte{},
		}
		spanAttributeRecords[i] = example.SpanAttribute{SpanID: spans[i].ID, Key: fmt.Sprintf("Key%d", i), AttributeValue: attribValue}
	}

	logger.Info("Span records : ", zap.Int("Count", len(spanRecords)))
	logger.Info("Span attribute records : ", zap.Int("Count", len(spanAttributeRecords)))

	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func(data []any, logger *zap.Logger, wg *sync.WaitGroup) {
		logger.Info("Inserting Span records")
		err = doChunkedInsert(context.TODO(), logger, dbInst, "otel.trace_span", data)
		if err != nil {
			fmt.Println(err)
		}
		wg.Done()
	}(spanRecords, logger, wg)

	wg.Add(1)
	go func(data []any, logger *zap.Logger, wg *sync.WaitGroup) {
		logger.Info("Inserting Span attribute records")
		err = doChunkedInsert(context.TODO(), logger, dbInst, "otel.trace_span_attribute", data)
		if err != nil {
			fmt.Println(err)
		}
		wg.Done()
	}(spanAttributeRecords, logger, wg)

	wg.Wait()
	// response, err := dbInst.InsertRecordsRaw(context.TODO(), "otel.trace_span_attribute", spanAttributeRecords)
	// if err != nil {
	// 	panic(err)
	// }

	fmt.Println("All done")
}

func runCreateJob(dbInst *kinetica.Kinetica) {
	start := time.Now()
	// result, err := dbInst.CreateJobRaw(
	// 	"OPENSKY")
	// if err != nil {
	// 	panic(err)
	// }
	duration := time.Since(start)
	// fmt.Println(result)
	fmt.Println("ShowSchema", duration.Milliseconds(), " ms")
}

func noop(elem interface{}) {

}

func doChunkedInsert(ctx context.Context, logger *zap.Logger, dbInst *kinetica.Kinetica, tableName string, records []any) error {

	logger.Info("Writing to - ", zap.String("Table", tableName), zap.Int("Record count", len(records)))

	recordChunks := ChunkBySize(records, 10000)

	errsChan := make(chan error, len(recordChunks))
	respChan := make(chan int, len(recordChunks))

	wg := &sync.WaitGroup{}
	// var mutex = &sync.Mutex{}

	for _, recordChunk := range recordChunks {
		wg.Add(1)
		go func(data []any, wg *sync.WaitGroup) {

			// mutex.Lock()
			resp, err := dbInst.InsertRecordsRaw(context.TODO(), tableName, data)
			errsChan <- err
			respChan <- resp.CountInserted
			// mutex.Unlock()

			wg.Done()
		}(recordChunk, wg)
	}
	wg.Wait()
	close(errsChan)
	close(respChan)
	var errs error
	for err := range errsChan {
		errs = multierr.Append(errs, err)
	}

	for resp := range respChan {
		fmt.Println("Count Inserted = ", resp)
	}
	return errs
}
