package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	// listen
	downStream, err := net.Listen("tcp", ":9090")

	handleError(err)
	defer downStream.Close()

	// keep accepting incoming connections
	for {
		conn, err := downStream.Accept()
		fmt.Println("New client: ", conn.RemoteAddr().String())
		handleError(err)

		// concurrently handle new connections
		go handleConnection(conn)
	}
}

func handleError(e error) {
	if e != nil {
		log.Fatal(e.Error())
	}
}

func handleConnection(conn net.Conn) {
	remoteAddress := conn.RemoteAddr().String()
	for {
		message, err := bufio.NewReader(conn).ReadString('\n')
		handleError(err)

		end := "q/" // client types 'q/' to exit
		if strings.TrimSpace(string(message)) == end {
			break
		}

		fmt.Print((remoteAddress + ": "), string(message))
	}
	fmt.Println("Closing client: ", remoteAddress)
	conn.Close()
}
