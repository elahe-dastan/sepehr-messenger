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

type ChatServer struct{
	seq    int32
	mutex  sync.Mutex
	queues map[int32]chan protocol.Data
}

func (s *ChatServer) Send(c context.Context, data *protocol.Data) (*empty.Empty, error) {
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

func (s *ChatServer) Receive(id *protocol.ID, con protocol.SimpleChat_ReceiveServer) error {
	messages := s.queues[id.Id]

	for {
		tick := time.Tick(time.Second * 10)


		select {
		case message := <-messages:
			if err := con.Send(&message); err != nil {
				log.Println(err)
			}
		case <-tick :
			return nil
		}
	}
}

func (s *ChatServer) Who(context.Context, *empty.Empty) (*protocol.ID, error) {
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
	protocol.RegisterSimpleChatServer(grpcServer, &ChatServer{
		seq:    0,
		mutex:  sync.Mutex{},
		queues: make(map[int32]chan protocol.Data),
	})
	grpcServer.Serve(l)
}