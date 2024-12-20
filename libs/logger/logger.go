package logger

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"time"
)

func init() {
	lg, err := NewLogger()
	if err != nil {
		panic(err)
	}
	DefaultLogger = lg
	//log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.SetFormatter(&log.TextFormatter{
		ForceColors:   true, // Force colors
		FullTimestamp: true, // Show full timestamp
	})
	log.SetReportCaller(true)
	log.SetOutput(lg)
}

var DefaultLogger *LoggerInst

type LoggerInst struct {
	logDir    string
	printFunc func(string)
}

func NewLogger() (*LoggerInst, error) {
	logDir := "./log"
	err := os.MkdirAll(logDir, os.ModePerm)
	if err != nil {
		return nil, err
	}

	return &LoggerInst{logDir: logDir}, nil
}

func (l *LoggerInst) Write(p []byte) (n int, err error) {
	if l.printFunc != nil {
		l.printFunc(string(p))
	}
	err = l.writeToFile(p)
	if err != nil {
		return 0, err
	}
	return len(p), nil
}

func (l *LoggerInst) writeToFile(p []byte) error {
	t := time.Now()
	logFile := fmt.Sprintf("%s/%s.log", l.logDir, t.Format("2006-01-02"))

	file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = file.Write(p)
	if err != nil {
		return err
	}

	return nil
}

func (l *LoggerInst) SetTermPrinter(printFunc func(string)) {
	l.printFunc = printFunc
}
