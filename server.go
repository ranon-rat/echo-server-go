
package main

import (
  	"log"
        "net"
	"fmt"   
	
)

type user struct{
  message []string
}
func (u *user) newMessage(text string){
  u.message=append(u.message,string(text))
}
func (message *user) messageRequest(conn net.Conn) {
	defer conn.Close()
	defer log.Println("Closed connection.")

	for {
		buf := make([]byte, 1024)
		size, err := conn.Read(buf)
		if err != nil {
			return
		}
		data := buf[:size]
		message.newMessage(string(data))
		fmt.Print(message.message)
		fmt.Println("asi es")
		conn.Write(data)
	}
}

func main() {
	message:=user{}

	l, err := net.Listen("tcp",":500")
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


