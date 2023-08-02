package tracer

type Options struct {
	Service string
	Env     string
}

type Tracer interface {
	Start(opts Options)
	Stop()
}
