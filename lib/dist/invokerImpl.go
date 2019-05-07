package dist

import (
	"fmt"
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

	err = s.Start()
	if err != nil {
		return err
	}

	defer s.Stop()
	fmt.Println("Invoker.invoke - conexão aberta")

	for {

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
		msgReceived.Body.ReplyBody = 1 // todo implementar chamada do objeto remoto

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
	}

	return nil
}
