package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	file, err := os.Open("messages.txt")

	if err != nil {
		log.Fatal("error: ", err)
	}

	lines := getLinesChannel(file)

	for line := range lines {
		fmt.Printf("read: %s\n", line)
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
