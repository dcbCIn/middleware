package server

import (
	"fmt"
	"net"
)

type ServerRequestHandlerImpl struct {
	Port       string
	listener   net.Listener
	connection net.Conn
}

func NewServerRequestHandlerImpl(port string) (srh *ServerRequestHandlerImpl, err error) {
	srh = &ServerRequestHandlerImpl{Port: port}

	srh.listener, err = net.Listen("tcp", ":"+srh.Port)

	if err != nil {
		return nil, err
	}

	return srh, nil
}

func (s *ServerRequestHandlerImpl) Start() (err error) {
	fmt.Println("ServerRequestHandler.Start - Aceitando conexões...")

	s.connection, err = s.listener.Accept()

	if err != nil {
		fmt.Println("ServerRequestHandler.Start - Erro ao abrir conexão")
		return err
	}

	fmt.Println("ServerRequestHandler.Start - Conexão aceita...")
	return nil
}

func (s *ServerRequestHandlerImpl) Stop() (err error) {
	fmt.Println("ServerRequestHandler.Stop - Closing connection")
	//err = s.connection.Close()
	if err != nil {
		return err
	}
	fmt.Println("ServerRequestHandler.Stop - Connection closed")
	err = s.listener.Close()
	if err != nil {
		return err
	}
	fmt.Println("ServerRequestHandler.Stop - Listener closed")
	return nil
}

func (s *ServerRequestHandlerImpl) Receive() (msg []byte, err error) {
	msg = make([]byte, 10240)
	n, err := s.connection.Read(msg)
	if err != nil {
		return nil, err
	}

	return msg[:n], nil
}

func (s *ServerRequestHandlerImpl) Send(msg []byte) (err error) {
	_, err = s.connection.Write(msg)
	if err != nil {
		return err
	}
	return nil
}
