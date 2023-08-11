package example

type NycTaxi struct {
	Vendor_id          string  `mapstructure:"vendor_id"`
	Pickup_datetime    int64   `mapstructure:"pickup_datetime"`
	Dropoff_datetime   int64   `mapstructure:"dropoff_datetime"`
	Passenger_count    int32   `mapstructure:"passenger_count"`
	Trip_distance      float32 `mapstructure:"trip_distance"`
	Pickup_longitude   float32 `mapstructure:"pickup_longitude"`
	Pickup_latitude    float32 `mapstructure:"pickup_latitude"`
	Rate_code_id       int32   `mapstructure:"rate_code_id"`
	Store_and_fwd_flag string  `mapstructure:"store_and_fwd_flag"`
	Dropoff_longitude  float32 `mapstructure:"dropoff_longitude"`
	Dropoff_latitude   float32 `mapstructure:"dropoff_latitude"`
	Payment_type       string  `mapstructure:"payment_type"`
	Fare_amount        float32 `mapstructure:"fare_amount"`
	Surcharge          float32 `mapstructure:"surcharge"`
	Mta_tax            float32 `mapstructure:"mta_tax"`
	Tip_amount         float32 `mapstructure:"tip_amount"`
	Tolls_amount       float32 `mapstructure:"tolls_amount"`
	Total_amount       float32 `mapstructure:"total_amount"`
	Cab_type           int32   `mapstructure:"cab_type"`
}

type NycTaxiMinMax struct {
	MaxLenVenvor int32 `mapstructure:"max_len_vendor"`
	MinLenVenvor int32 `mapstructure:"min_len_vendor"`
}

// Span - This structure can be used to insert records, read from a CSV file
// and also used as a return value in GetREcords, which is why each field
// is associated with different tags, eacj, pertaining to a different case
// The 'mapStructure' tags are used for GetREcords calls, the 'avro' tags
// assist in 'avro' marshalling for InsertRecords calls and the 'csv' tags
// are needed for reading in records into structures from CSV files.
type Span struct {
	ID string `mapstructure:"id" avro:"id" csv:"id"`

	ResourceID string `mapstructure:"resource_id" avro:"resource_id" csv:"resource_id"`

	ScopeID string `mapstructure:"scope_id" avro:"scope_id" csv:"scope_id"`

	EventID string `mapstructure:"event_id" avro:"event_id" csv:"event_id"`

	LinkID string `mapstructure:"link_id" avro:"link_id" csv:"link_id"`

	TraceID string `mapstructure:"trace_id" avro:"trace_id" csv:"trace_id"`

	SpanID string `mapstructure:"span_id" avro:"span_id" csv:"span_id"`

	ParentSpanID string `mapstructure:"parent_span_id" avro:"parent_span_id" csv:"parent_span_id"`

	TraceState string `mapstructure:"trace_state" avro:"trace_state" csv:"trace_state"`

	Name string `mapstructure:"name" avro:"name" csv:"name"`

	SpanKind int `mapstructure:"span_kind" avro:"span_kind" csv:"span_kind"`

	StartTimeUnixNano string `mapstructure:"start_time_unix_nano" avro:"start_time_unix_nano" csv:"start_time_unix_nano"`

	EndTimeUnixNano string `mapstructure:"end_time_unix_nano" avro:"end_time_unix_nano" csv:"end_time_unix_nano"`

	DroppedAttributesCount int `mapstructure:"dropped_attributes_count" avro:"dropped_attributes_count" csv:"dropped_attributes_count"`

	DroppedEventsCount int `mapstructure:"dropped_events_count" avro:"dropped_events_count" csv:"dropped_events_count"`

	DroppedLinksCount int `mapstructure:"dropped_links_count" avro:"dropped_links_count" csv:"dropped_links_count"`

	Message string `mapstructure:"message" avro:"message" csv:"message"`

	StatusCode int `mapstructure:"status_code" avro:"status_code" csv:"status_code"`
}

type SpanAttribute struct {
	SpanID         string `avro:"span_id"`
	Key            string `avro:"key"`
	AttributeValue `mapstructure:",squash"`
}

type AttributeValue struct {
	IntValue    int     `avro:"int_value"`
	StringValue string  `avro:"string_value"`
	BoolValue   int8    `avro:"bool_value"`
	DoubleValue float64 `avro:"double_value"`
	BytesValue  []byte  `avro:"bytes_value"`
}

type SampleExecuteSql struct {
	MaxID interface{} `mapstructure:"max_trace_id" avro:"column_1"`
}

type SampleGetRecords struct {
	MaxID int `mapstructure:"max_trace_id" avro:"trace_id"`
}

type WeatherSummary struct {
	EXPR_0 float64 `mapstructure:"EXPR_0" avro:"EXPR_0"`
	EXPR_1 float64 `mapstructure:"EXPR_1" avro:"EXPR_1"`
}

type Weather struct {
	City          string  `mapstructure:"city" avro:"city"`
	StateProvince string  `mapstructure:"state_province" avro:"state_province"`
	Country       string  `mapstructure:"country" avro:"country"`
	X             float64 `mapstructure:"x" avro:"x"`
	Y             float64 `mapstructure:"y" avro:"y"`
	AvgTemp       float64 `mapstructure:"avg_temp" avro:"avg_temp"`
	TimeZone      string  `mapstructure:"time_zone" avro:"time_zone"`
}
