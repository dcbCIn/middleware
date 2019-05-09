package dist

import (
	"jankenpo/shared"
	"middleware/lib/infra/client"
)

/*func NewInvocationImpl(objectId int, ipAddress string, portNumber int, operationName string, parameters []interface{}) *InvocationImpl {
	return &InvocationImpl{objectId: objectId, ipAddress: ipAddress, portNumber: portNumber, operationName: operationName, parameters: parameters}
}*/

// Implements requestor
type RequestorImpl struct{}

func (RequestorImpl) Invoke(inv Invocation) (t Termination, err error) {

	crh := client.NewClientRequestHandlerImpl(inv.IpAddress, inv.PortNumber)

	requestHeader := RequestHeader{inv.IpAddress, inv.ObjectId, true, inv.ObjectId, inv.OperationName}
	requestBody := RequestBody{inv.Parameters}

	msg := Message{
		Header{"GIOP", 1, true, 0, 0},
		Body{requestHeader, requestBody, ReplyHeader{}, nil}}

	var bytes []byte
	bytes, err = Marshall(msg)
	if err != nil {
		return Termination{}, err
	}

	err = crh.Send(bytes)
	if err != nil {
		return Termination{}, err
	}

	var msgReturned Message
	msgToBeUnmarshalled, err := crh.Receive()
	if err != nil {
		return Termination{}, err
	}
	msgReturned, err = Unmarshall(msgToBeUnmarshalled)
	if err != nil {
		return Termination{}, err
	}

	// Todo check if replyStatus of the message is valid

	shared.PrintlnInfo("RequestorImpl", "RequestorImpl.Invoke - Reply recebido e unmarshalled")
	t = Termination{msgReturned.Body.ReplyBody}

	return t, err
}
