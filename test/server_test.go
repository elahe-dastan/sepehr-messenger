package test

import (
	"context"
	"sync"
	"testing"

	"github.com/elahe-dastan/interview/client"
	"github.com/elahe-dastan/interview/config"
	"github.com/elahe-dastan/interview/protocol"
	"github.com/elahe-dastan/interview/server"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/stretchr/testify/assert"
)

func TestServer(t *testing.T) {
	// starting a server
	c := config.Read()
	s := server.ChatServer {
		Seq:    0,
		Mutex:  sync.Mutex{},
		Queues: make(map[int32]chan protocol.Data),
	}

	go func() {
		err := s.Start(c)
		assert.NoError(t, err, "cannot start server")
	}()

	// New three clients
	fCli, err := client.New("127.0.0.1:1373")
	assert.NoError(t, err, "cannot initiate client")
	sCli, err := client.New("127.0.0.1:1373")
	assert.NoError(t, err, "cannot initiate client")
	tCli, err := client.New("127.0.0.1:1373")
	assert.NoError(t, err, "cannot initiate client")

	var wg sync.WaitGroup
	wg.Add(3)

	var fID *protocol.ID
	var sID *protocol.ID
	var tID *protocol.ID

	go func() {
		fID, _ = fCli.Who(context.Background(), &empty.Empty{})
	}()

	go func() {
		sID, _ = sCli.Who(context.Background(), &empty.Empty{})
	}()

	go func() {
		tID, _ = tCli.Who(context.Background(), &empty.Empty{})
	}()

	wg.Wait()

	assert.NotEqual(t, fID.Id, sID.Id)
	assert.NotEqual(t, fID.Id, tID.Id)
	assert.NotEqual(t, sID.Id, tID.Id)

	fCli.Send(context.Background(), &protocol.Data {
		Id:                   fID,
		Text:                 "Hello from client one",
	})

	sCli.Send(context.Background(), &protocol.Data {
		Id:                   sID,
		Text:                 "Hello from client two",
	})

	tCli.Send(context.Background(), &protocol.Data {
		Id:                   tID,
		Text:                 "Hello from client three",
	})

	fChannel := s.Queues[fID.Id]
	assert.Equal(t, "Hello from client two", <-fChannel)
}