package impl

import (
	"jankenpo/shared"
	"middleware/lib"
	"middleware/lib/dist"
	"time"
)

const NAME = "jankenpo/mid/client"

func PlayJanKenPo(auto bool) (elapsed time.Duration) {
	lp := dist.NewLookupProxy(shared.NAME_SERVER_IP, shared.NAME_SERVER_PORT)
	cp, err := lp.Lookup("jankenpo")
	if err != nil {
		lib.PrintlnError(NAME, "Error at lookup")
	}
	err = lp.Close()
	if err != nil {
		lib.PrintlnError(NAME, "Error at closing lookup")
	}

	var jp dist.JankenpoProxy
	// connect to server
	//rpc.ConnectToServer("localhost", strconv.Itoa(shared.RPC_PORT))
	jp = *dist.NewJankenpoProxy(cp.Ip, cp.Port, cp.ObjectId)

	lib.PrintlnInfo(NAME, "Connected successfully")
	lib.PrintlnInfo(NAME)

	var player1Move, player2Move string
	// loop
	start := time.Now()
	for i := 0; i < shared.SAMPLE_SIZE; i++ {
		lib.PrintlnMessage(NAME, "Game", i)

		player1Move, player2Move = shared.GetMoves(auto)

		// send request to server and receive reply at the same time
		result, err := jp.Play(player1Move, player2Move)
		if err != nil {
			lib.FailOnError(NAME, err, "Erro ao obter resultado do jogo no servidor. Erro:")
		}

		lib.PrintlnMessage(NAME)
		switch result {
		case 1, 2:
			lib.PrintlnMessage(NAME, "The winner is Player", result)
		case 0:
			lib.PrintlnMessage(NAME, "Draw")
		default:
			lib.PrintlnMessage(NAME, "Invalid move")
		}
		lib.PrintlnMessage(NAME, "------------------------------------------------------------------")
		lib.PrintlnMessage(NAME)
		time.Sleep(shared.WAIT * time.Millisecond)
	}
	elapsed = time.Since(start)
	return elapsed
}
