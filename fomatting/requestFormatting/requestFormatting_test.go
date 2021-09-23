package requestFormatting

import (
	"TcpKeyValueStore/globals"
	"bytes"
	"testing"
)

func TestGetRequestMethod(t *testing.T) {
	t.Run("stream shorter than 3 characters should return error", func(t *testing.T) {
		input := []byte("pu")
		reader := bytes.NewReader(input)
		if _, err := GetRequestMethod(reader); err == nil {
			t.Error("should have returned an error")
		} else if err.Error() != "EOF" {
			t.Error("should have returned an EOF error")
		}
	})
	t.Run("stream 3 characters long should return method string", func(t *testing.T) {
		input := []byte("put")
		reader := bytes.NewReader(input)
		if result, err := GetRequestMethod(reader); err != nil {
			t.Error("should not have returned an error")
		} else if result != "put" {
			t.Error("incorrect string returned")
		}
	})
	t.Run("stream longer than 3 characters long should return first three as a string", func(t *testing.T) {
		input := []byte("put223")
		reader := bytes.NewReader(input)
		if result, err := GetRequestMethod(reader); err != nil {
			t.Error("should not have returned an error")
		} else if result != "put" {
			t.Error("incorrect string returned")
		}
	})
	t.Run("stream longer than 3 character should leave remaining chars on stream", func(t *testing.T) {
		input := []byte("put223")
		reader := bytes.NewReader(input)
		if _, err := GetRequestMethod(reader); err != nil {
			t.Error("should not have returned an error")
		}
		if remainingLength, err := reader.Read(input);
		err != nil && remainingLength != 3 {
			t.Error("input stream should contain 3 characters")
		}
	})
}

func TestGetRequestArguments(t *testing.T) {
	t.Run("valid input should return 1 expected argument", func(t *testing.T) {
		input := []byte("15stuff")
		reader := bytes.NewReader(input)
		if args, err := GetRequestArguments(reader, 1); err != nil {
			t.Error("should not have returned an error")
		} else if len(args) != 1 {
			t.Error("1 argument should have been returned")
		} else if args[0] != "stuff" {
			t.Error("incorrect argument string returned")
		}
	})
	t.Run("valid input should return 2 expected argument", func(t *testing.T) {
		input := []byte("15stuff210onetwothre")
		reader := bytes.NewReader(input)
		if args, err := GetRequestArguments(reader, 2); err != nil {
			t.Error("should not have returned an error")
		} else if len(args) != 2 {
			t.Error("2 argument should have been returned")
		} else if args[0] != "stuff" {
			t.Error("incorrect argument string returned")
		} else if args[1] != "onetwothre" {
			t.Error("incorrect argument string returned")
		}
	})
	t.Run("invalid input should return error", func(t *testing.T) {
		input := []byte("15stuf210onetwothre")
		reader := bytes.NewReader(input)
		if _, err := GetRequestArguments(reader, 2); err != ErrArgumentLengthCannotBeLessThan1 {
			t.Error("ErrStreamShorterThanRequested should have been returned")
		}
	})
	t.Run("invalid input stream should return an error", func(t *testing.T) {
		invalidInputs := []string{"1stuff", "12a", "23ty2"}
		for _, input := range invalidInputs {
			inputStream := []byte(input)
			reader := bytes.NewReader(inputStream)
			if _, err := GetRequestArguments(reader, 2); err == nil {
				t.Error("error should have been returned")
			}
		}
	})
}

func TestGetRequiredArgumentsCount(t *testing.T) {
	t.Run("valid RequestMethod should return correct argument count", func (t *testing.T) {
		methods := map[string]int{ globals.RequestMethodGet: 1, globals.RequestMethodDelete: 1, globals.RequestMethodPut: 2 }
		for method, expectedCount := range methods {
			if result, err := GetRequiredArgumentsCount(method); err != nil {
				t.Error("should not have returned an error")
			}  else if result != expectedCount {
				t.Error("incorrect result returned, expected", expectedCount, "but got", result)
			}
		}
	})
	t.Run("invalid request mtehod should return error", func(t *testing.T) {
		if _, err := GetRequiredArgumentsCount("clap"); err == nil {
			t.Error("should have returned an error")
		}
	})
}

func BenchmarkGetRequestMethod(b *testing.B) {
	input := []byte("put")
	reader := bytes.NewReader(input)
	for i := 0; i < b.N; i++ {
		GetRequestMethod(reader)
	}
}

func BenchmarkGetRequestArguments(b *testing.B) {
	input := []byte("15stuff")
	reader := bytes.NewReader(input)
	for i := 0; i < b.N; i++ {
		GetRequestArguments(reader, 1)
	}
}

func BenchmarkGetRequiredArgumentsCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetRequiredArgumentsCount(globals.RequestMethodGet)
	}
}
