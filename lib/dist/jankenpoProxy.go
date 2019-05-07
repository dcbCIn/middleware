package dist

type JankenpoProxy struct {
	Host     string
	Port     int
	objectId int
}

func NewJankenpoProxy(host string, port int, objectId int) *JankenpoProxy {
	jp := JankenpoProxy{host, port, objectId}
	return &jp
}

func (jp JankenpoProxy) Play(player1Move, player2Move string) (float64, error) {
	inv := *NewInvocation(jp.objectId, jp.Host, jp.Port, "jankenpo.play", []interface{}{player1Move, player2Move})
	requestor := RequestorImpl{}
	termination, err := requestor.Invoke(inv)
	if err != nil {
		return 0, err
	}
	winner := termination.Result.(float64)
	return winner, nil
}
