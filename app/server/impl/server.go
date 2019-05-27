package impl

import (
	"jankenpo/shared"
	"middleware/lib"
	"middleware/lib/dist"
	"middleware/lib/services/common"
	"sync"
)

const NAME = "jankenpo/mid/server"

/*func waitForConection(inv rpc.RPC, idx int) {
	lib.PrintlnInfo(NAME, "Connection", strconv.Itoa(idx), "started")

	// fecha o socket
	defer rpc.CloseConnection()

	// aceita conexões na porta
	rpc.WaitForConnection(idx)

	lib.PrintlnInfo(NAME, "Servidor finalizado (MyMiddleware)")
	lib.PrintlnInfo(NAME, "Connection", strconv.Itoa(idx), "ended")
}*/

func StartJankenpoServer() {
	var wg sync.WaitGroup
	lib.PrintlnInfo(NAME, "Initializing server MyMiddleware")

	lp := dist.NewLookupProxy(shared.NAME_SERVER_IP, shared.NAME_SERVER_PORT)
	err := lp.Bind("jankenpo", common.ClientProxy{"127.0.0.1", shared.MID_PORT, 1500})
	if err != nil {
		lib.PrintlnError(NAME, "Error at lookup")
	}
	err = lp.Close()
	if err != nil {
		lib.PrintlnError(NAME, "Error at closing lookup")
	}

	// escuta na porta tcp configurada
	var inv dist.InvokerImpl
	//inv.StartServer("", strconv.Itoa(shared.RPC_PORT))
	//defer inv.StopServer()

	go inv.Invoke(shared.MID_PORT)
	wg.Add(1)
	/*for idx := 0; idx < shared.CONECTIONS; idx++ {
		wg.Add(1)
		go func(i int) {
			waitForConection(inv, i)

			wg.Done()
		}(idx)
	}*/
	wg.Wait()
	lib.PrintlnInfo(NAME, "Fim do Servidor MyMiddleware")
}
