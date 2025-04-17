package logger

import (
	"log"
	"os"
	"sync"
	"time"
)

var (
	Log         *log.Logger
	logFile     *os.File
	currentDate string
	mutex       sync.Mutex
)

func init() {
	mutex.Lock()
	defer mutex.Unlock()

	currentDate = time.Now().Format("2006-01-02")
	openLogFile(currentDate)

	Log = log.New(logFile, "", 0)

	go autoRotate()
}

func openLogFile(date string) {
	logDir := "logs"
	os.MkdirAll(logDir, os.ModePerm)

	filePath := logDir + "/" + date + ".log"

	f, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("File to open log file: %v", err)
	}

	if logFile != nil {
		logFile.Close()
	}

	logFile = f
}

func autoRotate() {
	for {
		time.Sleep(1 * time.Minute)
		checkRotate()
	}
}

func checkRotate() {

}
