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
	u.message = append(u.message, string(text))//add a new value to a string
}
func (u *user) messageRequest(conn net.Conn) {

	defer conn.Close()//close the conection
	defer log.Println("Closed connection.")

	for {
		io.WriteString(conn, fmt.Sprintln(u.message[0:]))//send the string
		buf := make([]byte, 1024)//read the conection
		size, err := conn.Read(buf)
		if err != nil {
			return//if err return
		}
		data := buf[:size]//convert
		u.newMessage(string(data))//add to  the string
		fmt.Print(u.message)

	}
}

func main() {
	message := user{}

	l, err := net.Listen("tcp", ":500")//listen 
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
