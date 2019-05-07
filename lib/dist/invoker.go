package dist

type Invoker interface {
	Invoke(port int, nameServer bool) (err error)
	//StopServer()
}
