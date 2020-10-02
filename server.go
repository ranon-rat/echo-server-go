
package main

import (
        "bufio"
        "fmt"
        "net"
        "strings"

)

type user struct{
  message []string
}
func (u *user) newMessage(text string){
  u.message=append(u.message,string(text))
}
func main() {
        mensaje:=user{}

        l, err := net.Listen("tcp", ":500")
        if err != nil {
                fmt.Println(err)
                return
        }
        defer l.Close()

        c, err := l.Accept()
        if err != nil {
        	c.Close()
                fmt.Println(err)
                return
        }


        for {
                go readAndSend(mensaje,c)
        }
}

func readAndSend(mensaje user,c net.Conn){

  for {
          netData, err := bufio.NewReader(c).ReadString('\n')
          if err != nil {
                  fmt.Println(err)
                  c.Close()
                  return
                  mensaje.newMessage(string(netData))
          }


          if strings.TrimSpace(string(netData)) == "STOP" {
                  fmt.Println("Exiting TCP server!")
                  return
          }

          fmt.Print("-> ", string(netData))
          c.Write([]byte(netData))
          value:=strings.Join(mensaje.message,"")
          c.Write([]byte(value))
        }
  }
