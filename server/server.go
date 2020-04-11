package server

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

type Server struct {
	conn []net.Conn
	connReaders chan net.Conn
}

func New() Server {
	return Server{conn:make([]net.Conn, 0),
		connReaders:make(chan net.Conn, 100)}
}

func (s Server) Start()  {
	go s.accept()

	for {
		fmt.Println("inside")
		r := <-s.connReaders
		netData, err := bufio.NewReader(r).ReadString('\n')

		if err != nil {
			log.Println(err)
		}

		for _, c := range s.conn {
			if r != c {
				w := bufio.NewWriter(c)
				w.WriteString(netData)
				w.Flush()
			}
		}
		fmt.Println(netData)
	}
}

func (s *Server) accept()  {
	l, err := net.Listen("tcp",":1373")

	if err != nil {
		log.Println(err)
	}

	for {
		c, erro := l.Accept()
		if erro != nil {
			log.Println(err)
		}

		s.conn = append(s.conn, c)
		s.connReaders <- c
	}
}