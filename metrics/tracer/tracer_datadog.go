package tracer

import dd_tracer "gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"

type ddTracer struct{}

func NewDDTracer(opts Options) Tracer {
	return nil
}

func (ddT *ddTracer) Start(opts Options) {
	dd_tracer.Start(
		dd_tracer.WithService(opts.Service),
		dd_tracer.WithEnv(opts.Env),
	)
}

func (ddt *ddTracer) Stop() {
	dd_tracer.Stop()
}
