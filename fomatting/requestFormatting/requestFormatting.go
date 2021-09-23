package requestFormatting

import (
	"TcpKeyValueStore/globals"
	"errors"
	"io"
	"strconv"
)

var (
	ErrStreamShorterThanRequested = errors.New("stream is shorter than the length requested")
	ErrArgumentLengthCannotBeLessThan1 = errors.New("argumentLength cannot be less then 1")
	ErrInvalidRequestMethod = errors.New("invalid request method")
)

func GetRequestMethod(r io.Reader) (string, error) {
	return streamReader(r, 3)
}

func GetRequestArguments(r io.Reader, requiredArgumentsCount int) ([]string, error) {
	var arguments []string
	for i := 0; i < requiredArgumentsCount; i++ {
		argumentSizeLengthString, err := streamReader(r, 1)
		if err != nil {
			return nil, err
		}
		argumentSizeLength, err := strconv.Atoi(argumentSizeLengthString)
		if err != nil {
			return nil, err
		}
		argumentLengthString, err := streamReader(r, argumentSizeLength)
		if err != nil {
			return nil, err
		}
		argumentLength, err := strconv.Atoi(argumentLengthString)
		if err != nil {
			return nil, err
		} else if argumentLength < 1 {
			return nil, ErrArgumentLengthCannotBeLessThan1
		}
		argument, err := streamReader(r, argumentLength)
		if err != nil {
			return nil, err
		}
		arguments = append(arguments, argument)
	}
	return arguments, nil
}

func GetRequiredArgumentsCount(method string) (int, error) {
	switch method {
	case globals.RequestMethodPut:
		return 2, nil
	case globals.RequestMethodGet, globals.RequestMethodDelete:
		return 1, nil
	}
	return 0, ErrInvalidRequestMethod
}

func streamReader(r io.Reader, length int) (string, error) {
	buffer := make([]byte, length)
	singleCharBuffer := make([]byte, 1)
	var count int
	for {
		if n, err := r.Read(singleCharBuffer); err != nil {
			return "", err
		} else if n == 1 {
			if count > length {
				return "", ErrStreamShorterThanRequested
			}
			buffer[count] = singleCharBuffer[0]
			count++
		}
		if count == length {
			return string(buffer), nil
		}
	}
}