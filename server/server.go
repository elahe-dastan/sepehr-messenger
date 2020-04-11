package server

import (
	"alibaba/client"
	"bufio"
	"fmt"
	"log"
	"net"
)

type Server struct {
	clients []client.Client
	connReaders []*bufio.Reader
}

func New() Server {
	return Server{clients:make([]client.Client, 0)}
}

func (s Server) Start()  {
	l, err := net.Listen("tcp",":1373")

	if err != nil {
		log.Println(err)
	}

	for i := 0 ; i < 2; i++ {
		c, erro := l.Accept()
		if erro != nil {
			log.Println(err)
		}

		s.clients = append(s.clients, client.New(c))
		s.connReaders = append(s.connReaders, bufio.NewReader(c))
	}

	for {
		fmt.Println("inside")
		netData, err := s.connReaders[0].ReadString('\n')

		if err != nil {
			log.Println(err)
		}

		fmt.Println(netData)
	}
}