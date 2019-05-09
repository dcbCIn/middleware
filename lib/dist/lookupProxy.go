package dist

import "middleware/lib/services/common"

type LookupProxy struct {
	Host string
	Port int
}

func (lp LookupProxy) Bind(sn string, cp common.ClientProxy) (err error) {
	inv := *NewInvocation(0, lp.Host, lp.Port, "Bind", []interface{}{sn, cp})
	requestor := RequestorImpl{}
	_, err = requestor.Invoke(inv)
	if err != nil {
		return err
	}
	return nil
}

func (lp LookupProxy) Lookup(serviceName string) (cp common.ClientProxy, err error) {
	inv := *NewInvocation(0, lp.Host, lp.Port, "Lookup", []interface{}{serviceName})
	requestor := RequestorImpl{}
	termination, err := requestor.Invoke(inv)
	if err != nil {
		return cp, err
	}

	clientProxyMap := termination.Result.(map[string]interface{})
	cp = common.ClientProxy{clientProxyMap["Ip"].(string), int(clientProxyMap["Port"].(float64)), int(clientProxyMap["ObjectId"].(float64))}

	return cp, nil
}

func NewLookupProxy(host string, port int) *LookupProxy {
	return &LookupProxy{host, port}
}
