package opentracing

import "golang.org/x/net/context"

type contextKey int

const (
	tracerKey  contextKey = iota
	activeSpanKey
)


// ContextWithSpan returns a new `context.Context` that holds a reference to
// the given `Span`.
func ContextWithTracer(ctx context.Context, tracer Tracer) context.Context {
	return context.WithValue(ctx, tracerKey, tracer)
}

// TracerFromContext returns the `Tracer` previously associated with `ctx`, or
// `nil` if no such `Tracer` could be found.
func TracerFromContext(ctx context.Context) Tracer {
	val := ctx.Value(tracerKey)
	if Tracer, ok := val.(Tracer); ok {
		return Tracer
	}
	return nil
}



// ContextWithSpan returns a new `context.Context` that holds a reference to
// the given `Span`.
func ContextWithSpan(ctx context.Context, span Span) context.Context {
	return context.WithValue(ctx, activeSpanKey, span)
}

// SpanFromContext returns the `Span` previously associated with `ctx`, or
// `nil` if no such `Span` could be found.
func SpanFromContext(ctx context.Context) Span {
	val := ctx.Value(activeSpanKey)
	if span, ok := val.(Span); ok {
		return span
	}
	return nil
}

// StartSpanFromContext starts and returns a Span with `operationName`, using
// any Span found within `ctx` as a parent. If no such parent could be found,
// StartSpanFromContext creates a root (parentless) Span.
//
// The second return value is a context.Context object built around the
// returned Span.
//
// Example usage:
//
//    SomeFunction(ctx context.Context, ...) {
//        sp, ctx := opentracing.StartSpanFromContext(ctx, "SomeFunction")
//        defer sp.Finish()
//        ...
//    }
func StartSpanFromContext(ctx context.Context, operationName string) (Span, context.Context) {
	parent := SpanFromContext(ctx)
	tracer := TracerFromContext(ctx)
	span := tracer.StartSpanWithOptions(StartSpanOptions{
		OperationName: operationName,
		Parent:        parent,
	})
	return span, ContextWithSpan(ctx, span)
}
