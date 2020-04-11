package client

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

type Client struct {
	Con      net.Conn
	Writer   *bufio.Writer
	Reader   *bufio.Reader
	Incoming chan string
}

func New(c net.Conn) Client {
	return Client{
		Con:      c,
		Writer:   bufio.NewWriter(c),
		Reader:   bufio.NewReader(c),
		Incoming: make(chan string, 100),
	}
}

func (cli *Client) Send(text string)  {
	cli.Writer.WriteString(text)
	cli.Writer.Flush()
}

func (cli *Client) Show()  {
	for {
		text, err := cli.Reader.ReadString('\n')

		if err != nil {
			log.Println(err)
		}

		cli.Incoming<- text

		fmt.Println(text)
	}
}

// API GRPC
// 2 days