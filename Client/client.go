package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func onMessage(conn net.Conn) {
	for {
		newReader := bufio.NewReader(conn)
		msg, err := newReader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print(msg)
	}
}
func main() {
	conn, err := net.Dial("tcp", "localhost:3000")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("Input your name: ")
	nameReader := bufio.NewReader(os.Stdin)
	nameInput, _ := nameReader.ReadString('\n')
	nameInput = nameInput[:len(nameInput) - 2]
	fmt.Println("*********Message*********")

	go onMessage(conn)

	for {
		msgReader := bufio.NewReader(os.Stdin)
		msg, err := msgReader.ReadString('\n')
		if err != nil {
			break
		}
		msg = fmt.Sprintf("%s: %s\n", nameInput, msg[:len(msg) - 2])
		conn.Write([]byte(msg))
	}
	conn.Close()
} 