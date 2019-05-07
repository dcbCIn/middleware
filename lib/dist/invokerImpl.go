package dist

import (
	"jankenpo/shared"
	"middleware/app/server/remoteObjects"
	"middleware/lib/infra/server"
	"middleware/lib/services/common"
)

type InvokerImpl struct {
}

func (inv InvokerImpl) Invoke(port int, nameServer bool) (err error) {
	s, err := server.NewServerRequestHandlerImpl(port)
	if err != nil {
		return err
	}
	defer s.StopServer()
	shared.PrintlnInfo("InvokerImpl", "Invoker.invoke - conexão aberta")

	var lookup = common.Lookup{}
	var jankenpo = remoteObjects.Jankenpo{}

	for {
		err = s.Start()
		if err != nil {
			return err
		}

		shared.PrintlnInfo("InvokerImpl", "Invoker.invoke - Aguardando mensagem")

		msgToBeUnmarshalled, err := s.Receive()
		if err != nil {
			return err
		}
		shared.PrintlnInfo("InvokerImpl", "Invoker.invoke - Mensagem recebida")

		msgReceived, err := Unmarshall(msgToBeUnmarshalled)

		if err != nil {
			return err
		}

		shared.PrintlnInfo("InvokerImpl", "Invoker.invoke - Mensagem unmarshalled")

		switch msgReceived.Body.RequestHeader.Operation {
		case "jankenpo.Play":
			player1Move := msgReceived.Body.RequestBody.Parameters[0].(string)
			player2Move := msgReceived.Body.RequestBody.Parameters[1].(string)
			msgReceived.Body.ReplyHeader = ReplyHeader{"", msgReceived.Body.RequestHeader.RequestId, 1}
			msgReceived.Body.ReplyBody, _ = jankenpo.Play(player1Move, player2Move)
		case "lookup.Bind":
			serviceName := msgReceived.Body.RequestBody.Parameters[0].(string)
			clientProxyMap := msgReceived.Body.RequestBody.Parameters[1].(map[string]interface{})
			clientProxy := common.ClientProxy{clientProxyMap["Ip"].(string), int(clientProxyMap["Port"].(float64)), int(clientProxyMap["ObjectId"].(float64))}
			msgReceived.Body.ReplyHeader = ReplyHeader{"", msgReceived.Body.RequestHeader.RequestId, 1}
			msgReceived.Body.ReplyBody = lookup.Bind(serviceName, clientProxy)
		case "lookup.Lookup":
			serviceName := msgReceived.Body.RequestBody.Parameters[0].(string)
			msgReceived.Body.ReplyHeader = ReplyHeader{"", msgReceived.Body.RequestHeader.RequestId, 1}
			msgReceived.Body.ReplyBody, _ = lookup.Lookup(serviceName)
		}

		var bytes []byte
		bytes, err = Marshall(msgReceived)
		if err != nil {
			return err
		}

		shared.PrintlnInfo("InvokerImpl", "Invoker.invoke - Retorno marshalled")

		err = s.Send(bytes)
		if err != nil {
			return err
		}

		shared.PrintlnInfo("InvokerImpl", "Invoker.invoke - Mensagem enviada")
		err = s.CloseConnection()
		if err != nil {
			return err
		}
	}

	return nil
}
