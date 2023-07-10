package logger

import (
	"log"
	"os"
)

var (
	InfoLog  *log.Logger = log.New(os.Stdout, "Info: ", log.Ltime)
	ErrorLog *log.Logger = log.New(os.Stderr, "Error: ", log.Ltime|log.Lshortfile)
)
