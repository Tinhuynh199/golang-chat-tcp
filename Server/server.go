package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

var (
	conns []net.Conn
	connCh = make(chan net.Conn)
	closeCh = make(chan net.Conn)
	msgCh = make(chan string)
)
func main() {
	server, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatal(err)
	}

	go func ()  {
		for {
			conn, err := server.Accept()
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("New client has connected")
			conns = append(conns, conn)
			connCh <- conn
		}
	} ()

	for {
		select {
		case conn := <- connCh:
			go onMessage(conn)
		case msg := <- msgCh:
			fmt.Print(msg)
		case conn := <- closeCh:
			fmt.Println("Client exit")
			removeConn(conn)
		}
	}
}

func removeConn(conn net.Conn) {
	var index int
	for index = range conns {
		if conns[index] == conn {
			break
		}
	}
	conns = append(conns[:index], conns[index + 1:]...)
	// conns = append(conns[index:], conns[:index + 1]...)
}

func publicMsg(conn net.Conn, msg string) {
	for i:= range conns {
		if conns[i] != conn {
			conns[i].Write([]byte(msg))
		}
	}
}

func onMessage(conn net.Conn) {
	for {
		reader := bufio.NewReader(conn)
		msg, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		msgCh <- msg
		publicMsg(conn, msg)
	}
	closeCh <- conn
}