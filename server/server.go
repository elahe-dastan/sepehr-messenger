package server

import (
	"alibaba/protocol"
	"context"
	"log"
	"net"
	"sync"
	"time"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

//type Server struct {
//	clients []client.Client
//	connReaders chan client.Client
//}

//func New() Server {
//	return Server{clients:make([]client.Client, 0),
//		connReaders:make(chan client.Client, 100)}
//}

type SimpleChat struct{
	seq    int32
	mutex  sync.Mutex
	queues map[int32]chan protocol.Data
}

func (s *SimpleChat) Send(c context.Context, data *protocol.Data) (*empty.Empty, error) {
	for id, ch := range s.queues {
		if id == data.Id.Id {
			continue
		}

		select {
		case ch <- *data:
		default:
			continue
		}
	}

	return &empty.Empty{}, nil
}

func (s *SimpleChat) Receive(id *protocol.ID, con protocol.SimpleChat_ReceiveServer) error {
	messages := s.queues[id.Id]

	for {
		tick := time.Tick(time.Minute * 5)

		select {
		case message := <-messages:
			if err := con.Send(&message); err != nil {
				log.Println(err)
			}
		case <-tick :
			break
		}
	}

	return nil
}

func (s *SimpleChat) Who(context.Context, *empty.Empty) (*protocol.ID, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.seq++
	s.queues[s.seq] = make(chan protocol.Data, 100)
	return &protocol.ID{
		Id: s.seq,
	}, nil
}

func Start()  {
	l, err := net.Listen("tcp",":1373")

	if err != nil {
		log.Println(err)
	}

	grpcServer := grpc.NewServer()
	protocol.RegisterSimpleChatServer(grpcServer, &SimpleChat{
		seq:    0,
		mutex:  sync.Mutex{},
		queues: make(map[int32]chan protocol.Data),
	})
	grpcServer.Serve(l)

	//go s.accept(l)
	//
	//for {
	//	r := <-s.connReaders
	//	go s.Handler(r)
	//}
}

//func (s *Server) accept(l net.Listener)  {
//	for {
//		c, err := l.Accept()
//		if err != nil {
//			log.Println(err)
//		}
//
//		cli := client.New(c)
//		s.clients = append(s.clients, cli)
//		s.connReaders<- cli
//	}
//}

//func (s *Server) Handler(r client.Client)  {
//	netData, err := r.Reader.ReadString('\n')
//
//	if err != nil {
//		log.Println(err)
//	}
//
//	for _, cli := range s.clients {
//		if r != cli {
//			cli.Send(netData)
//		}
//	}
//}