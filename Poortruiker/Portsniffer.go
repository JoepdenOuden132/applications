package main

import (
	"fmt"
	"log"
	"net"
	"strconv"
)

func main() {
	//Hier begint de code met aangeven welke port ik wil scannen
	ports := []int{80, 443}
	//Dit stukje is een for loop. In deze loop worden de port meerdere keren gescant.
	// range geeft aan dat alle intergers in ports moeten worden gebruikt
	for _, port := range ports {
		go listenPort(port)
	}

	// dit zorgt ervoor dat de bovenstaande for loop geactiveerd blijft
	select {}
}

// functie 1
func listenPort(port int) {

	// de bovenste lijn probeerd te luisteren op de ports, als dit mislukt word er een error gegeven
	listener, err := net.Listen("tcp", ":"+strconv.Itoa(port))
	if err != nil {
		log.Fatalf("Failed to listen on port %d: %v", port, err)
	}
	// Hij wacht op reactie en daarna sluit hij af
	defer listener.Close()

	fmt.Println("Port sniffer listening on port %d\n", port)

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

	remoteAddr := conn.RemoteAddr().String()
	localAddr := conn.LocalAddr().String()
	fmt.Printf("Connection accepted from %s to %s\n", remoteAddr, localAddr)
}
