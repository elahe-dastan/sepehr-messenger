package client

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

type Client struct {
	con net.Conn
	Writer *bufio.Writer
	Reader *bufio.Reader
	Incoming chan string
}

func New(c net.Conn) Client {
	return Client{
		con:    c,
		Writer: bufio.NewWriter(c),
		Reader: bufio.NewReader(c),
		Incoming:make(chan string, 100),
	}
}

//func (cli Client) Send()  {
//	for {
//		fmt.Println("send")
//		consoleReader := bufio.NewReader(os.Stdin)
//		fmt.Print(" >> ")
//		text, er := consoleReader.ReadString('\n')
//		if er != nil {
//			log.Println(er)
//		}
//
//		fmt.Println(text)
//		cli.Writer.WriteString(text)
//		cli.Writer.Flush()
//	}
//}

func (cli Client) Show()  {
	for {
		fmt.Println("inshow")
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