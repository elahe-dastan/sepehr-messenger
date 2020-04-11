package server

import (
	"alibaba/client"
	"log"
	"net"
)

type Server struct {
	clients []client.Client
	connReaders chan client.Client
}

func New() Server {
	return Server{clients:make([]client.Client, 0),
		connReaders:make(chan client.Client, 100)}
}

func (s *Server) Start()  {
	go s.accept()

	for {
		r := <-s.connReaders
		go s.Handler(r)
	}
}

func (s *Server) accept()  {
	l, err := net.Listen("tcp",":1373")

	if err != nil {
		log.Println(err)
	}

	for {
		c, err := l.Accept()
		if err != nil {
			log.Println(err)
		}

		cli := client.New(c)
		s.clients = append(s.clients, cli)
		s.connReaders<- cli
	}
}

func (s *Server) Handler(r client.Client)  {
	netData, err := r.Reader.ReadString('\n')

	if err != nil {
		log.Println(err)
	}

	for _, cli := range s.clients {
		if r != cli {
			cli.Send(netData)
		}
	}
}