package opentracing

import "golang.org/x/net/context"

type noopTraceContext struct{}
type noopSpan struct{}
type noopTraceContextSource struct{}
type noopTracer struct {
	noopTraceContextSource
}

var (
	defaultNoopTraceContext       = noopTraceContext{}
	defaultNoopSpan               = noopSpan{}
	defaultNoopTraceContextSource = noopTraceContextSource{}
	defaultNoopTracer             = noopTracer{}
	emptyTags                     = Tags{}
	emptyBytes                    = []byte{}
	emptyStringMap                = map[string]string{}
)

const (
	emptyString = ""
)

// noopTraceContext:

func (n noopTraceContext) SetTraceAttribute(key, val string) TraceContext { return n }
func (n noopTraceContext) TraceAttribute(key string) string               { return emptyString }

// noopSpan:
func (n noopSpan) StartChild(operationName string) Span {
	return defaultNoopSpan
}
func (n noopSpan) SetTag(key string, value interface{}) Span { return n }
func (n noopSpan) Finish()                                   {}
func (n noopSpan) TraceContext() TraceContext                { return defaultNoopTraceContext }
func (n noopSpan) AddToGoContext(ctx context.Context) (Span, context.Context) {
	return n, GoContextWithSpan(ctx, n)
}
func (n noopSpan) LogEvent(event string)                                 {}
func (n noopSpan) LogEventWithPayload(event string, payload interface{}) {}
func (n noopSpan) Log(data LogData)                                      {}

// noopTraceContextSource:
func (n noopTraceContextSource) TraceContextToBinary(tcid TraceContext) ([]byte, []byte) {
	return emptyBytes, emptyBytes
}
func (n noopTraceContextSource) TraceContextToText(tcid TraceContext) (map[string]string, map[string]string) {
	return emptyStringMap, emptyStringMap
}
func (n noopTraceContextSource) TraceContextFromBinary(
	traceContextID []byte,
	traceAttrs []byte,
) (TraceContext, error) {
	return defaultNoopTraceContext, nil
}
func (n noopTraceContextSource) TraceContextFromText(
	traceContextID map[string]string,
	traceAttrs map[string]string,
) (TraceContext, error) {
	return defaultNoopTraceContext, nil
}
func (n noopTraceContextSource) NewRootTraceContext() TraceContext {
	return defaultNoopTraceContext
}
func (n noopTraceContextSource) NewChildTraceContext(parent TraceContext) (TraceContext, Tags) {
	return defaultNoopTraceContext, emptyTags
}

// noopTracer:
func (n noopTracer) StartTrace(operationName string) Span {
	return defaultNoopSpan
}

func (n noopTracer) StartSpanWithContext(operationName string, ctx TraceContext) Span {
	return defaultNoopSpan
}

func (n noopTracer) JoinTrace(operationName string, parent interface{}) Span {
	return defaultNoopSpan
}