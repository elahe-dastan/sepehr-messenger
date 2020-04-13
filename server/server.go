package server

import (
	"context"
	"log"
	"net"
	"sync"
	"time"

	"github.com/elahe-dastan/gossip/config"
	"github.com/elahe-dastan/gossip/protocol"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

const size = 100 // This is the buffer size for each client's channel

type ChatServer struct {
	Seq    int32
	Mutex  sync.Mutex
	Queues map[int32]chan protocol.Data
}

func NewChatServer() *ChatServer {
	return &ChatServer{
		Seq:    0,
		Mutex:  sync.Mutex{},
		Queues: make(map[int32]chan protocol.Data),
	}
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
	const openConnTime = 5 * time.Second

	messages := s.Queues[id.Id]

	for {
		ticker := time.NewTicker(openConnTime)
		select {
		case message := <-messages:
			if err := con.Send(&message); err != nil {
				log.Println(err)
			}
		case <-ticker.C:
			ticker.Stop()
			return nil
		}
	}
}

func (s *ChatServer) Who(context.Context, *empty.Empty) (*protocol.ID, error) {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()
	s.Seq++
	s.Queues[s.Seq] = make(chan protocol.Data, size)

	return &protocol.ID{
		Id: s.Seq,
	}, nil
}

func (s *ChatServer) Start(c config.Config) error {
	l, err := net.Listen("tcp", c.Address)
	if err != nil {
		return err
	}

	server := grpc.NewServer()
	protocol.RegisterChatServer(server, s)

	return server.Serve(l)
}
