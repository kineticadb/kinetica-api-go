package kinetica

import (
	"context"

	"go.opentelemetry.io/otel/trace"
)

func NewDefaultShowSystemStatusOptions() *ShowSystemStatusOptions {
	return &ShowSystemStatusOptions{}
}

func (kinetica *Kinetica) buildShowSystemStatusOptionsMap(
	ctx context.Context, options *ShowSystemStatusOptions) *map[string]string {
	var (
		// childCtx context.Context
		childSpan trace.Span
	)

	_, childSpan = kinetica.tracer.Start(ctx, "kinetica.buildShowSystemStatusOptionsMap()")
	defer childSpan.End()

	mapOptions := make(map[string]string)

	return &mapOptions
}

func (kinetica *Kinetica) ShowSystemStatusRaw(ctx context.Context) (*ShowSystemStatusResponse, error) {
	var (
		childCtx  context.Context
		childSpan trace.Span
	)

	childCtx, childSpan = kinetica.tracer.Start(ctx, "kinetica.ShowSystemStatusRaw()")
	defer childSpan.End()

	return kinetica.ShowSystemStatusRawWithOpts(childCtx, NewDefaultShowSystemStatusOptions())
}

func (kinetica *Kinetica) ShowSystemStatusRawWithOpts(
	ctx context.Context, options *ShowSystemStatusOptions) (*ShowSystemStatusResponse, error) {
	var (
		childCtx  context.Context
		childSpan trace.Span
	)

	childCtx, childSpan = kinetica.tracer.Start(ctx, "kinetica.ShowSystemStatusRawWithOpts()")
	defer childSpan.End()

	mapOptions := kinetica.buildShowSystemStatusOptionsMap(childCtx, options)
	response := ShowSystemStatusResponse{}
	request := ShowSystemStatusRequest{Options: *mapOptions}
	err := kinetica.submitRawRequest(
		childCtx, "/show/system/status",
		&Schemas.showSystemStatusRequest, &Schemas.showSystemStatusResponse,
		&request, &response)

	return &response, err
}
