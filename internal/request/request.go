package request

import (
	"errors"
	"io"
	"strings"
)

type Request struct {
	RequestLine RequestLine
	// Header      map[string]string
	// Body        []byte
}

type RequestLine struct {
	HttpVersion   string
	RequestTarget string
	Method        string
}

func RequestFromReader(reader io.Reader) (*Request, error) {
	content, err := io.ReadAll(reader)

	if err != nil {
		return nil, err
	}

	requestLine, err := parseRequestLine(string(content))

	if err != nil {
		return nil, err
	}

	request := Request{RequestLine: *requestLine}

	return &request, nil
}

func parseRequestLine(requestContent string) (*RequestLine, error) {
	requestLine := RequestLine{}

	requestParts := strings.Split(requestContent, "\r\n")
	requestLineParts := strings.Split(requestParts[0], " ")

	// Check if request line contains all 3 required parts
	if len(requestLineParts) != 3 {
		return nil, errors.New("Request Line does not contain all required parts.")
	}

	method := requestLineParts[0]

	if !isAllUpperLetters(requestLineParts[0]) {
		return nil, errors.New("Request method is not correct http method")
	}

	httpVersion := strings.Split(requestLineParts[2], "/")[1]

	if httpVersion != "1.1" {
		return nil, errors.New("Invalid http version. We only support 1.1")
	}

	requestLine.Method = method
	requestLine.RequestTarget = requestLineParts[1]
	requestLine.HttpVersion = httpVersion

	return &requestLine, nil
}

func isAllUpperLetters(word string) bool {
	if len(word) == 0 {
		return false
	}

	for i := 0; i < len(word); i++ {
		letter := word[i]
		if letter < 'A' || letter > 'Z' {
			return false
		}
	}

	return true
}
