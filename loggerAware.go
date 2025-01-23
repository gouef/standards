package standards

type LoggerAware interface {
	SetLogger(logger Logger)
}
