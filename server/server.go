package server

import (
	"context"
	"log"
	"net"
	"sync"
	"time"

	"github.com/elahe-dastan/interview/config"
	"github.com/elahe-dastan/interview/protocol"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

type ChatServer struct {
	Seq    int32
	Mutex  sync.Mutex
	Queues map[int32]chan protocol.Data
}

func (s *ChatServer) Send(c context.Context, data *protocol.Data) (*empty.Empty, error) {
	for id, ch := range s.Queues {
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

func (s *ChatServer) Receive(id *protocol.ID, con protocol.Chat_ReceiveServer) error {
	messages := s.Queues[id.Id]
	openConnTime := 5 * time.Second
	ticker := time.NewTicker(openConnTime)
	defer ticker.Stop()

	for {
		select {
		case message := <-messages:
			if err := con.Send(&message); err != nil {
				log.Println(err)
			}
		case <-ticker.C:
			return nil
		}
	}
}

func (s *ChatServer) Who(context.Context, *empty.Empty) (*protocol.ID, error) {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()
	s.Seq++
	chanSize := 100
	s.Queues[s.Seq] = make(chan protocol.Data, chanSize)

	return &protocol.ID{
		Id: s.Seq,
	}, nil
}

func (s *ChatServer) Start(c config.Config) error {
	l, err := net.Listen("tcp", c.Address)

	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer()
	protocol.RegisterChatServer(grpcServer, s)

	if err := grpcServer.Serve(l); err != nil {
		return err
	}

	return nil
}
