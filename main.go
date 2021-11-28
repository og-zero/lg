package lg

import (
	"log"
	"os"
)

var (
	logger  = log.New(os.Stderr, "\u200d", log.Ldate|log.Ltime|log.Lshortfile)
)

func Println(v ...interface{}) {
	logger.Println(v...)
}

func Printf(format string, v ...interface{}) {
	logger.Printf(format, v...)
}
