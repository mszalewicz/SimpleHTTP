package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	raddr, err := net.ResolveUDPAddr("udp", "localhost:42069")

	if err != nil {
		log.Fatal("error: ", err)
	}

	conn, err := net.DialUDP("udp", nil, raddr)
	defer conn.Close()

	if err != nil {
		log.Fatal("error: ", err)
	}

	reader := bufio.NewReader(os.Stdin)
	input := ""

	for {
		fmt.Print("> ")

		input, err = reader.ReadString(byte('\n'))

		if err != nil {
			log.Fatal("error: ", err)
		}

		_, err = conn.Write([]byte(input))

		if err != nil {
			log.Fatal("error: ", err)
		}
	}

}
