package responseFormatting

import (
	"TcpKeyValueStore/globals"
	"fmt"
	"strconv"
)

func FormatResponse(method string, coordinatorResult string, err error) []byte {
	if method == globals.RequestMethodGet && err == nil {
		responseString := getResultToResponseString(coordinatorResult)
		return []byte(responseString)
	}
	return []byte(coordinatorResult)
}

func getResultToResponseString(value string) string {
	valueSizeString := getStringLengthAsString(value)
	valueSizeLengthString := getStringLengthAsString(valueSizeString)
	return fmt.Sprint("val", valueSizeLengthString, valueSizeString, value)
}

func getStringLengthAsString(value string) string {
	valueLength := len(value)
	valueLengthString := strconv.Itoa(valueLength)
	return valueLengthString
}