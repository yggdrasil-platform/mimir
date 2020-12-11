package logger

import (
	"log"
	"os"
)

var (
	Info  = log.New(os.Stdout, "\033[1;33m[INFO]\033[0m ", log.Ldate|log.Ltime)
	Error = log.New(os.Stderr, "\033[0;31m[ERROR]\033[0m ", log.Ldate|log.Ltime|log.Lshortfile)
)
