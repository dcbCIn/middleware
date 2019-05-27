package common

import (
	"middleware/lib"
)

type ClientProxy struct {
	Ip   string
	Port int
	// Todo Quando houver mais de um protocolo adicionar protocolo ao AOR
	//protocol string // Quando o Client Request Handler usa mais de um protocolo, a AOR também inclui a identificação do protocolo.
	ObjectId int
}

type NamingRecord struct {
	serviceName string
	clientProxy ClientProxy
}

type ILookup interface {
	Bind(sn string, cp ClientProxy) (err error)
	Lookup(serviceName string) (cp ClientProxy, err error)
}

type Lookup struct {
	services []NamingRecord
}

func (l *Lookup) Bind(sn string, cp ClientProxy) (err error) {
	lib.PrintlnInfo("Lookup", "Service bind =", sn)
	l.services = append(l.services, NamingRecord{sn, cp})
	return nil
}

func (l Lookup) Lookup(serviceName string) (cp ClientProxy, err error) {
	lib.PrintlnInfo("Lookup", "Service lookup =", serviceName)
	for _, nr := range l.services {
		return nr.clientProxy, nil
	}
	return ClientProxy{}, nil
}
