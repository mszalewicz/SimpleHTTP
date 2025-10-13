package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":42069")
	defer listener.Close()

	if err != nil {
		log.Fatal("error: ", err)
	}

	for {
		conn, err := listener.Accept()
		defer conn.Close()

		if err != nil {
			log.Fatal("error: ", err)
		} else {
			fmt.Println("Connection has been accepted")
		}

		lines := getLinesChannel(conn)

		for line := range lines {
			fmt.Println(line)
		}

		fmt.Println("Connection has been closed")
	}
}

func getLinesChannel(f io.ReadCloser) <-chan string {
	out := make(chan string, 1)

	go func(f io.ReadCloser, out chan string) {
		defer f.Close()
		defer close(out)

		data := make([]byte, 8)
		line := ""
		for {
			n, err := f.Read(data)

			if err != nil {
				break
			}

			chunk := data[:n]

			if index := bytes.Index(chunk, []byte("\n")); index != -1 {
				line += string(chunk[:index])
				chunk = chunk[index+1:]
				out <- line
				line = ""
			}

			line += string(chunk)
		}

		if len(line) != 0 {
			fmt.Print("read: ", line, "\n")
		}
	}(f, out)

	return out
}
