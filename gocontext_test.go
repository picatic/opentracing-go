package opentracing

import (
	"testing"

	"golang.org/x/net/context"
)

func TestContextWithSpan(t *testing.T) {
	span := &noopSpan{}
	ctx := ContextWithSpan(context.Background(), span)
	span2 := SpanFromContext(ctx)
	if span != span2 {
		t.Errorf("Not the same span returned from context, expected=%+v, actual=%+v", span, span2)
	}

	ctx = context.Background()
	span2 = SpanFromContext(ctx)
	if span2 != nil {
		t.Errorf("Expected nil span, found %+v", span2)
	}

	ctx = ContextWithSpan(ctx, span)
	span2 = SpanFromContext(ctx)
	if span != span2 {
		t.Errorf("Not the same span returned from context, expected=%+v, actual=%+v", span, span2)
	}
}

func TestStartSpanFromContext(t *testing.T) {
	testTracer := testTracer{}
	testTracerCtx :=  ContextWithTracer(context.Background(), testTracer)

	// Test the case where there *is* a Span in the Context.
	{
		parentSpan := &testSpan{}
		parentCtx := ContextWithSpan(testTracerCtx, parentSpan)
		childSpan, childCtx := StartSpanFromContext(parentCtx, "child")
		if !childSpan.(testSpan).HasParent {
			t.Errorf("Failed to find parent: %v", childSpan)
		}
		if childSpan != SpanFromContext(childCtx) {
			t.Errorf("Unable to find child span in context: %v", childCtx)
		}
	}

	// Test the case where there *is not* a Span in the Context.
	{
		childSpan, childCtx := StartSpanFromContext(testTracerCtx, "child")
		if childSpan.(testSpan).HasParent {
			t.Errorf("Should not have found parent: %v", childSpan)
		}
		if childSpan != SpanFromContext(childCtx) {
			t.Errorf("Unable to find child span in context: %v", childCtx)
		}
	}
}
func TestContextWithTracer(t *testing.T) {
	tracer := &NoopTracer{}
	ctx := ContextWithTracer(context.Background(), tracer)
	tracer2 := TracerFromContext(ctx)
	if tracer != tracer2 {
		t.Errorf("Not the same tracer returned from context, expected=%+v, actual=%+v", tracer, tracer2)
	}

	ctx = context.Background()
	tracer2 = TracerFromContext(ctx)
	if tracer2 != nil {
		t.Errorf("Expected nil tracer, found %+v", tracer2)
	}

	ctx = ContextWithTracer(ctx, tracer)
	tracer2 = TracerFromContext(ctx)
	if tracer != tracer2 {
		t.Errorf("Not the same tracer returned from context, expected=%+v, actual=%+v", tracer, tracer2)
	}
}
