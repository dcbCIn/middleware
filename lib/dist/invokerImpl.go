package dist

import (
	"fmt"
	"middleware/app/remoteObjects"
	"middleware/lib/infra/server"
)

type InvokerImpl struct {
	stop bool
}

func (inv InvokerImpl) Stop() {
	inv.stop = true
}

func (inv InvokerImpl) Invoke() (err error) {
	s, err := server.NewServerRequestHandlerImpl("1234")

	if err != nil {
		return err
	}

	defer s.StopServer()
	fmt.Println("Invoker.invoke - conexão aberta")

	var jankenpo = remoteObjects.Jankenpo{}

	for {
		err = s.Start()
		if err != nil {
			return err
		}

		fmt.Println("Invoker.invoke - Aguardando mensagem")

		msgToBeUnmarshalled, err := s.Receive()
		if err != nil {
			return err
		}
		fmt.Println("Invoker.invoke - Mensagem recebida")

		msgReceived, err := Unmarshall(msgToBeUnmarshalled)

		if err != nil {
			return err
		}

		fmt.Println("Invoker.invoke - Mensagem unmarshalled")

		msgReceived.Body.ReplyHeader = ReplyHeader{"", msgReceived.Body.RequestHeader.RequestId, 1}
		player1Move := msgReceived.Body.RequestBody.Parameters[0].(string)
		player2Move := msgReceived.Body.RequestBody.Parameters[1].(string)
		msgReceived.Body.ReplyBody = jankenpo.Process(player1Move, player2Move)

		var bytes []byte
		bytes, err = Marshall(msgReceived)
		if err != nil {
			return err
		}

		fmt.Println("Invoker.invoke - Retorno marshalled")

		err = s.Send(bytes)
		if err != nil {
			return err
		}

		fmt.Println("Invoker.invoke - Mensagem enviada")
		err = s.CloseConnection()
		if err != nil {
			return err
		}
	}

	return nil
}
