package main

import (
        "bufio"
        "fmt"
        "net"
        "os"
        "strings"
)

func main() {


        conn, err := net.Dial("tcp", ":500")
    
        if err != nil {
                fmt.Println(err)
                return
        }

        for {
                reader := bufio.NewReader(os.Stdin)
                fmt.Print(">> ")
                text, _ := reader.ReadString('\n')
                fmt.Fprintf(conn, text+"\n")
                conn.Write([]byte(text))
                message, _ := bufio.NewReader(conn).ReadString('\n')
                fmt.Print("->: " + message)
                if strings.TrimSpace(string(text)) == "STOP" {
                        fmt.Println("TCP client exiting...")
                        return
                }
        }
}
