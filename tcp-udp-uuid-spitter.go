package main

import (
	"fmt"
	"net"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
)

func handleTCPConnection(conn net.Conn) {
	defer conn.Close()
	uuid := uuid.New().String()
	fmt.Fprintf(conn, "%s\n", uuid)
}

func handleUDPConnection(conn *net.UDPConn) {
	buffer := make([]byte, 1024)
	n, _, err := conn.ReadFromUDP(buffer)
	if err != nil {
		fmt.Println(err)
		return
	}
	data := string(buffer[:n])
	data = strings.TrimSpace(data)
	uuid := uuid.New().String()
	conn.WriteTo([]byte(uuid+"\n"), nil)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		tcpAddr, err := net.ResolveTCPAddr("tcp", ":8080")
		if err != nil {
			fmt.Println(err)
			return
		}
		listener, err := net.ListenTCP("tcp", tcpAddr)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer listener.Close()
		for {
			conn, err := listener.Accept()
			if err != nil {
				fmt.Println(err)
				break
			}
			go handleTCPConnection(conn)
		}
	}()

	go func() {
		defer wg.Done()
		udpAddr, err := net.ResolveUDPAddr("udp", ":8080")
		if err != nil {
			fmt.Println(err)
			return
		}
		conn, err := net.ListenUDP("udp", udpAddr)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer conn.Close()
		for {
			handleUDPConnection(conn)
			time.Sleep(time.Second)
		}
	}()

	wg.Wait()
}
