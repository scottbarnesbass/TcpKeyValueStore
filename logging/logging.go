package logging

import (
	"errors"
	"fmt"
	"log"
	"os"
)

var (
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
	ErrUnsupportedLogDestinationChosen = errors.New("unsupported log destination chosen")
	ErrValidFilenameRequiredForLogToFile = errors.New("valid filename required for LogToFile")
)

const (
	maxLogStringLength = 50
	LogToFile LogDestination = "logToFile"
	LogToConsole LogDestination = "logToConsole"
)

type LogDestination string

func SetupLogging(fileName string, logDestination LogDestination) (*os.File, error) {

	switch logDestination{
	case LogToFile:
		if fileName == "" {
			return nil, ErrValidFilenameRequiredForLogToFile
		}

		file, err := os.OpenFile(fmt.Sprint(fileName, ".log"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

		if err == nil {
			log.SetOutput(file)
			InfoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile|log.Llongfile)
			WarningLogger = log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile|log.Llongfile)
			ErrorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile|log.Llongfile)
		}
		return file, err
	case LogToConsole:
		log.SetOutput(os.Stderr)
		InfoLogger = log.New(os.Stderr,"INFO: ", log.Ldate|log.Ltime|log.Lshortfile|log.Llongfile)
		WarningLogger = log.New(os.Stderr, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile|log.Llongfile)
		ErrorLogger = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile|log.Llongfile)
		return nil, nil
	}

	return nil, ErrUnsupportedLogDestinationChosen
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