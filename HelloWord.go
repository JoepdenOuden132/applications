package main

import (
	"fmt"
	"log"
	"net"
	"strconv"
)

func main() {
	ports := []int{80, 443}

	for _, port := range ports {
		go listenPort(port)
	}

	select {}
}

func listenPort(port int) {
	listener, err := net.Listen("tpc", ":"+strconv.Itoa(port))
	if err != nil {
		log.Fatalf("Failed to listen on port %d: %v", port, err)
	}

	defer listener.Close()

	fmt.Printf("Port sniffer listening on port %d\n", port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalf("Failed to accept connection on port %d: %v", port, err)
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

}
