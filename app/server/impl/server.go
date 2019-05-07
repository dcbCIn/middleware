package impl

import (
	"jankenpo/shared"
	"middleware/lib/dist"
	"middleware/lib/services/common"
	"sync"
)

const NAME = "jankenpo/mid/server"

/*func waitForConection(inv rpc.RPC, idx int) {
	shared.PrintlnInfo(NAME, "Connection", strconv.Itoa(idx), "started")

	// fecha o socket
	defer rpc.CloseConnection()

	// aceita conexões na porta
	rpc.WaitForConnection(idx)

	shared.PrintlnInfo(NAME, "Servidor finalizado (MyMiddleware)")
	shared.PrintlnInfo(NAME, "Connection", strconv.Itoa(idx), "ended")
}*/

func StartJankenpoServer() {
	var wg sync.WaitGroup
	shared.PrintlnInfo(NAME, "Initializing server MyMiddleware")

	lp := dist.LookupProxy{shared.NAME_SERVER_IP, shared.NAME_SERVER_PORT}
	err := lp.Bind("jankenpo", common.ClientProxy{"127.0.0.1", shared.MID_PORT, 1500})
	if err != nil {
		shared.PrintlnError(NAME, "Error at lookup")
	}

	// escuta na porta tcp configurada
	var inv dist.InvokerImpl
	//inv.StartServer("", strconv.Itoa(shared.RPC_PORT))
	//defer inv.StopServer()

	go inv.Invoke(shared.MID_PORT, false)
	wg.Add(1)
	/*for idx := 0; idx < shared.CONECTIONS; idx++ {
		wg.Add(1)
		go func(i int) {
			waitForConection(inv, i)

			wg.Done()
		}(idx)
	}*/
	wg.Wait()
	shared.PrintlnInfo(NAME, "Fim do Servidor MyMiddleware")
}
