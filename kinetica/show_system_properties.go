package kinetica

import (
	"context"

	"go.opentelemetry.io/otel/trace"
)

func NewDefaultShowSystemPropertiesOptions() *ShowSystemPropertiesOptions {
	return &ShowSystemPropertiesOptions{
		Properties: "",
	}
}

func (kinetica *Kinetica) buildShowSystemPropertiesOptionsMap(ctx context.Context, options *ShowSystemPropertiesOptions) *map[string]string {
	var (
		// childCtx context.Context
		childSpan trace.Span
	)

	_, childSpan = kinetica.tracer.Start(ctx, "kinetica.buildShowSystemPropertiesOptionsMap()")
	defer childSpan.End()

	mapOptions := make(map[string]string)
	if options.Properties != "" {
		mapOptions["properties"] = options.Properties
	}

	return &mapOptions
}

func (kinetica *Kinetica) ShowSystemPropertiesRaw(ctx context.Context) (*ShowSystemPropertiesResponse, error) {
	var (
		childCtx  context.Context
		childSpan trace.Span
	)

	childCtx, childSpan = kinetica.tracer.Start(ctx, "kinetica.ShowSystemPropertiesRaw()")
	defer childSpan.End()

	return kinetica.ShowSystemPropertiesRawWithOpts(childCtx, NewDefaultShowSystemPropertiesOptions())
}

func (kinetica *Kinetica) ShowSystemPropertiesRawWithOpts(
	ctx context.Context, options *ShowSystemPropertiesOptions) (*ShowSystemPropertiesResponse, error) {
	var (
		childCtx  context.Context
		childSpan trace.Span
	)

	childCtx, childSpan = kinetica.tracer.Start(ctx, "kinetica.ShowSystemPropertiesRawWithOpts()")
	defer childSpan.End()

	mapOptions := kinetica.buildShowSystemPropertiesOptionsMap(childCtx, options)
	response := ShowSystemPropertiesResponse{}
	request := ShowSystemPropertiesRequest{Options: *mapOptions}
	err := kinetica.submitRawRequest(
		childCtx, "/show/system/properties",
		&Schemas.showSystemPropertiesRequest, &Schemas.showSystemPropertiesResponse,
		&request, &response)

	return &response, err
}
