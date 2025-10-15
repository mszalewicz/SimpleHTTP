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

	if len(requestLineParts) != 3 {
		return nil, errors.New("Request Line does not contain all required parts.")
	}

	requestLine.Method = requestLineParts[0]
	requestLine.RequestTarget = requestLineParts[1]
	requestLine.HttpVersion = strings.Split(requestLineParts[2], "/")[1]

	return &requestLine, nil
}
