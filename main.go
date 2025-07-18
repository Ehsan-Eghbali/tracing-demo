package main

import (
	"fmt"
	"log"
	"net/http"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.12.0"
)

func main() {
	InitTracer()

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		tracer := otel.Tracer("my-service")
		ctx, span := tracer.Start(ctx, "handle-hello")
		defer span.End()

		fmt.Fprintln(w, "Hello from Jaeger-traced Go service!")
	})

	fmt.Println("Listening on :8085")
	log.Fatal(http.ListenAndServe(":8085", nil))
}

func InitTracer() {
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(
		jaeger.WithEndpoint("http://localhost:14268/api/traces"),
	))
	if err != nil {
		log.Fatalf("failed to initialize exporter: %v", err)
	}

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exp),
		sdktrace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String("go-jaeger-demo"),
		)),
	)
	otel.SetTracerProvider(tp)
}
