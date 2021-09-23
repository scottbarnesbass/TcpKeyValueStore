package responseFormatting

import (
	"TcpKeyValueStore/globals"
	"errors"
	"testing"
)

func TestFormatResponse(t *testing.T) {
	t.Run("successful get response should be formatted correctly", func(t *testing.T){
		expectedResult := "val19something"
		result := FormatResponse(globals.RequestMethodGet, "something", nil)
		if string(result) != expectedResult {
			t.Error("incorrect result returned, expected", expectedResult, "but got", string(result))
		}
	})
	t.Run("failed get response should return correct failure response", func(t *testing.T) {
		coordinatorResults := []string{"nil", "err"}
		for _, coordinatorResult := range coordinatorResults {
			formattedResult :=  FormatResponse(globals.RequestMethodGet, coordinatorResult, errors.New("some error"))
			if string(formattedResult) != coordinatorResult {
				t.Error("incorrect result returned, expected", coordinatorResult, "but got", string(formattedResult))
			}
		}
	})
}

func BenchmarkFormatResponse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FormatResponse(globals.RequestMethodGet, "something", nil)
	}
}


