package main

import (
	"net"
)

func main() {

	ln, _ := net.Listen("tcp", ":8081")
    for {    
      conn, _ := ln.Accept()
      go process(conn)
    }

}

func process(conn net.Conn) {
	defer conn.Close()

	for {
		msg, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Println("Message received:", string(msg))
		conn.Write([]byte(msg + "\n"))

	}
}
