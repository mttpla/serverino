//Copyright 2021 Matteo Paoli - mttpla@gmail.com

package util

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

var (
	//Trace deeply info
	Trace *log.Logger
	//Info standard logs
	Info *log.Logger
	//Warning sameting to take care of
	Warning *log.Logger
	//Error somethings go wrong
	Error *log.Logger
)

func LogInit() {
	logLevel := Configuration.LogLevel
	switch logLevel {
	case "TRACE":
		logSettings(os.Stdout, os.Stdout, os.Stdout, os.Stderr)
	case "INFO":
		logSettings(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)
	case "WARNING":
		logSettings(ioutil.Discard, ioutil.Discard, os.Stdout, os.Stderr)
	case "ERROR":
		logSettings(ioutil.Discard, ioutil.Discard, ioutil.Discard, os.Stderr)
	default:
		logSettings(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)
	}
	Trace.Println("LogLevel: ", logLevel)
}

func logSettings(
	traceHandle io.Writer,
	infoHandle io.Writer,
	warningHandle io.Writer,
	errorHandle io.Writer) {

	Trace = log.New(traceHandle,
		"TRACE: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Info = log.New(infoHandle,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Warning = log.New(warningHandle,
		"WARNING: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Error = log.New(errorHandle,
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}
