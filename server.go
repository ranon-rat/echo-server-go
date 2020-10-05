package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

type user struct {
	message []string
}

func (u *user) newMessage(text string) {
	u.message = append(u.message, string(text))
}
func (u *user) messageRequest(conn net.Conn) {

	defer conn.Close()
	defer log.Println("Closed connection.")

	for {
		io.WriteString(conn, fmt.Sprintln(u.message[0:]))
		buf := make([]byte, 1024)
		size, err := conn.Read(buf)
		if err != nil {
			return
		}
		data := buf[:size]
		u.newMessage(string(data))
		fmt.Print(u.message)

	}
}

func main() {
	message := user{}

	l, err := net.Listen("tcp", ":500")
	if err != nil {
		log.Panicln(err)
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Panicln(err)
		}

		fmt.Println(message.message)

		go message.messageRequest(conn)

	}
}
