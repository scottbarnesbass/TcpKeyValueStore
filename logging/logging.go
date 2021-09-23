package logging

import (
	"fmt"
	"log"
	"os"
)

var (
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
)

const (
	maxLogStringLength = 50
)

// TODO add prop to switch to switch to standard output, not file
func SetupLogging(fileName string) (*os.File, error) {
	file, err := os.OpenFile(fmt.Sprint(fileName,".log"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

	if err == nil {
		log.SetOutput(file)
		InfoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile|log.Llongfile)
		WarningLogger = log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile|log.Llongfile)
		ErrorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile|log.Llongfile)
	}

	return file, err
}

func PrepLogArgs(args ...interface{}) []string {
	var resultStrings []string
	for _, v := range args {
		argString := fmt.Sprint(v)
		if len(argString) > maxLogStringLength {
			argString = argString[:maxLogStringLength]
		}
		resultStrings = append(resultStrings, argString)
	}
	return resultStrings
}