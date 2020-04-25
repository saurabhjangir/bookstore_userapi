package log

import (
	"flag"
	"fmt"
	"github.com/google/logger"
	"os"
)


const logPath = "/tmp/bookstore.log"

var (
	Log *logger.Logger
	verbose = flag.Bool("verbose", false, "print info level logs to stdout")
)

func init(){
	lf, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)
	if err != nil {
		logger.Fatalf("Failed to open log file: %v", err)
	}
	Log = logger.Init("LoggerExample", *verbose, true, lf)
	if Log == nil {
		panic("Failed to initialize log service")
	}
	fmt.Println("logger service initialized successfully")
}
