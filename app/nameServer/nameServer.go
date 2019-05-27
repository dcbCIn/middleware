package main

import (
	"jankenpo/shared"
	"middleware/lib"
	"middleware/lib/dist"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	lib.PrintlnInfo("nameServer", "Initializing server MyMiddleware(NameServer)")

	// escuta na porta tcp configurada
	var inv dist.InvokerImpl
	go inv.Invoke(shared.NAME_SERVER_PORT)
	wg.Add(1)

	/*for idx := 0; idx < shared.CONECTIONS; idx++ {
		wg.Add(1)
		go func(i int) {
			waitForConection(inv, i)

			wg.Done()
		}(idx)
	}*/
	wg.Wait()
	lib.PrintlnInfo("nameServer", "Fim do Servidor MyMiddleware(NameServer)")
}
