package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	client    []Clients
	curClient int
	curConn   net.Conn
	reader    *bufio.Reader
	timeout   = 500 * time.Millisecond
)

//Clients structure to Handle Multiple Users
type Clients struct {
	Client    net.Conn //Client Specific Connection
	IPAddress string
}

func handleConnection() {
	for {
		if curClient != 0 {
			clientIndex := curClient - 1
			curConn = client[clientIndex].Client
			fmt.Print("Client: " + strconv.Itoa(clientIndex) + ">> ")
		} else {
			fmt.Print(">> ")
		}
		reader = bufio.NewReader(os.Stdin)
		textInput, _ := reader.ReadString('\n')
		textInput = strings.Replace(textInput, "\r", "", -1)
		textInput = strings.Replace(textInput, "\n", "", -1)

		if textInput == "clients" && curClient == 0 {
			if len(client) == 0 {
				fmt.Println("No Users Are Connected")
			} else {
				for i := 1; i < len(client)+1; i++ {
					clientIndex := strconv.Itoa(i)
					clientIP := client[i-1].IPAddress
					fmt.Println(clientIndex + " | " + clientIP)
				}
			}
		} else if strings.Split(textInput, " ")[0] == "use" && curClient == 0 {
			textSplit := strings.Split(textInput, " ")
			if len(textSplit) == 2 {
				interactWith, err := strconv.Atoi(textSplit[1])
				if err != nil {
					fmt.Println(err)
				}
				curClient = interactWith
			}
		} else if textInput == "quit" && curClient != 0 {
			curClient = 0
		} else {
			if curClient != 0 {
				client[curClient-1].Client.Write([]byte(textInput))
			}
		}
	}
}

func checkMessage() {
	for {
		for i, v := range client {
			v.Client.SetReadDeadline(time.Now().Add(timeout))
			buf := make([]byte, 1024)
			size, _ := v.Client.Read(buf)
			data := string(buf[:size])
			if size != 0 {
				fmt.Println("Received From Client " + strconv.Itoa(i) + ": " + data)
			}
		}
	}
}

func main() {

	l, err := net.Listen("tcp", ":500") //Listen for connections
	if err != nil {
		log.Panicln(err)
	}
	defer l.Close()

	go handleConnection() //run function in a goroutine
	go checkMessage()     //run function in a goroutine

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Panicln(err)
		}
		IPAddress := conn.RemoteAddr().String()
		structData := Clients{Client: conn, IPAddress: IPAddress}
		client = append(client, structData)
	}
}
