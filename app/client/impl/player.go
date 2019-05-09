package impl

import (
	"jankenpo/shared"
	"middleware/lib/dist"
	"time"
)

const NAME = "jankenpo/mid/client"

func PlayJanKenPo(auto bool) (elapsed time.Duration) {
	lp := dist.LookupProxy{shared.NAME_SERVER_IP, shared.NAME_SERVER_PORT}
	cp, err := lp.Lookup("jankenpo")
	if err != nil {
		shared.PrintlnError(NAME, "Error at lookup")
	}

	var jp dist.JankenpoProxy
	// connect to server
	//rpc.ConnectToServer("localhost", strconv.Itoa(shared.RPC_PORT))
	jp = *dist.NewJankenpoProxy(cp.Ip, cp.Port, cp.ObjectId)

	shared.PrintlnInfo(NAME, "Connected successfully")
	shared.PrintlnInfo(NAME)

	var player1Move, player2Move string
	// loop
	start := time.Now()
	for i := 0; i < shared.SAMPLE_SIZE; i++ {
		shared.PrintlnMessage(NAME, "Game", i)

		player1Move, player2Move = shared.GetMoves(auto)

		// send request to server and receive reply at the same time
		result, err := jp.Play(player1Move, player2Move)
		if err != nil {
			shared.PrintlnError(NAME, "Erro ao obter resultado do jogo no servidor. Erro:", err)
		}

		shared.PrintlnMessage(NAME)
		switch result {
		case 1, 2:
			shared.PrintlnMessage(NAME, "The winner is Player", result)
		case 0:
			shared.PrintlnMessage(NAME, "Draw")
		default:
			shared.PrintlnMessage(NAME, "Invalid move")
		}
		shared.PrintlnMessage(NAME, "------------------------------------------------------------------")
		shared.PrintlnMessage(NAME)
		time.Sleep(shared.WAIT * time.Millisecond)
	}
	elapsed = time.Since(start)
	return elapsed
}