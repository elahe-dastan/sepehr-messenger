package client

import (
	"bufio"
	"net"
)

type Client struct {
	con net.Conn
	Writer *bufio.Writer
	Reader *bufio.Reader
}

func New(c net.Conn) Client {
	return Client{
		con:    c,
		Writer: bufio.NewWriter(c),
		Reader: bufio.NewReader(c),
	}
}