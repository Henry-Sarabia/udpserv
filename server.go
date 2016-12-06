package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func init() {
	service := ":1200"
	udpAddr, err := net.ResolveUDPAddr("up4", service)
	if err != nil {
		fmt.Println("Server: Address Resolution error: ", err)
		os.Exit(1)
	}

	conn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		fmt.Println("Server: Listening error: ", err)
		os.Exit(1)
	}

	for {
		handleClient(conn)
	}
}

func handleClient(conn *net.UDPConn) {
	buf := make([]byte, 512, 512) //Turn into slice later?

	_, addr, err := conn.ReadFromUDP(buf[0:])
	if err != nil {
		return
	}

	daytime := time.Now().String()

	conn.WriteToUDP([]byte(daytime), addr)
}
