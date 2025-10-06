package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("messages.txt")

	if err != nil {
		log.Fatal("error", "error", err)
	}

	data := make([]byte, 8)
	for {
		n, err := file.Read(data)

		if err != nil {
			break
		}

		fmt.Printf("read: %s\n", string(data[:n]))
	}
}
