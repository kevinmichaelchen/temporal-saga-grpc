package temporal

import (
	"go.temporal.io/sdk/interceptor"
)

func ClientInterceptors() ([]interceptor.ClientInterceptor, error) {
	//i, err := opentelemetry.NewTracingInterceptor(opentelemetry.TracerOptions{
	//	TextMapPropagator: otel.GetTextMapPropagator(),
	//})
	//if err != nil {
	//	return nil, fmt.Errorf("failed to create OTEL tracing interceptor: %w", err)
	//}
	//return []interceptor.ClientInterceptor{
	//	i,
	//}, nil
	return nil, nil
}

func WorkerInterceptors() ([]interceptor.WorkerInterceptor, error) {
	//i, err := opentelemetry.NewTracingInterceptor(opentelemetry.TracerOptions{
	//	TextMapPropagator: otel.GetTextMapPropagator(),
	//})
	//if err != nil {
	//	return nil, fmt.Errorf("failed to create OTEL tracing interceptor: %w", err)
	//}
	//return []interceptor.WorkerInterceptor{
	//	i,
	//}, nil
	return nil, nil
}
