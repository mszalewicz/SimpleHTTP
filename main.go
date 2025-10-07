package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("messages.txt")

	if err != nil {
		log.Fatal("error: ", err)
	}

	data := make([]byte, 8)
	line := ""
	for {
		n, err := file.Read(data)

		if err != nil {
			break
		}

		chunk := data[:n]

		if index := bytes.Index(chunk, []byte("\n")); index != -1 {
			line += string(chunk[:index])
			chunk = chunk[index+1:]
			fmt.Printf("read: %s\n", line)
			line = ""
		}

		line += string(chunk)
	}

	if len(line) != 0 {
		fmt.Print("read: ", line, "\n")
	}
}
