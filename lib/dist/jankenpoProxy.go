package dist

import "middleware/lib"

type JankenpoProxy struct {
	Host      string
	Port      int
	ObjectId  int
	requestor Requestor
}

func NewJankenpoProxy(host string, port int, objectId int) *JankenpoProxy {
	return &JankenpoProxy{host, port, objectId, NewRequestorImpl(host, port)}
}

func (jp JankenpoProxy) Play(player1Move, player2Move string) (float64, error) {
	inv := *NewInvocation(jp.ObjectId, jp.Host, jp.Port, lib.FunctionName(), []interface{}{player1Move, player2Move})
	termination, err := jp.requestor.Invoke(inv)
	if err != nil {
		return -1, err
	}
	winner := termination.Result.(float64)
	return winner, nil
}
