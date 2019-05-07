package dist

type Invoker interface {
	Invoker() (err error)
	StopServer()
}
