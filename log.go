package log

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

// File write definition
var (
	filePermission os.FileMode = 0666
	fileFlag                   = os.O_APPEND | os.O_CREATE | os.O_RDWR
	filePath                   = "../logFile"
)

// MainLogger struct
type MainLogger struct {
	*logrus.Logger
	ModuleName   string
	ModuleDesc   string
	FunctionName string
}

// Logging level definition
var (
	INFO  = "INFO"
	ERROR = "ERROR"
	WARN  = "WARN"
	FATAL = "FATAL"
)

// ExecLog for executing log as desired
func (m *MainLogger) ExecLog(level string, args ...interface{}) {
	switch strings.ToUpper(level) {
	case INFO:
		m.WithFields(logrus.Fields{
			"moduleName":   m.ModuleName,
			"moduleDesc":   m.ModuleDesc,
			"functionName": m.FunctionName,
		}).Info(args...)
	case ERROR:
		m.WithFields(logrus.Fields{
			"moduleName":   m.ModuleName,
			"moduleDesc":   m.ModuleDesc,
			"functionName": m.FunctionName,
		}).Error(args...)
	case WARN:
		m.WithFields(logrus.Fields{
			"moduleName":   m.ModuleName,
			"moduleDesc":   m.ModuleDesc,
			"functionName": m.FunctionName,
		}).Warn(args...)
	case FATAL:
		m.WithFields(logrus.Fields{
			"moduleName":   m.ModuleName,
			"moduleDesc":   m.ModuleDesc,
			"functionName": m.FunctionName,
		}).Fatal(args...)
	default:
		m.WithFields(logrus.Fields{
			"moduleName":   m.ModuleName,
			"moduleDesc":   m.ModuleDesc,
			"functionName": m.FunctionName,
		}).Print(args...)
	}
}

// NewLogger Function to create new logger
func NewLogger(moduleName string, moduleDesc string, functionName string) *MainLogger {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		os.Mkdir(filePath, filePermission)
	}
	f, err := os.OpenFile(filePath+"/"+moduleDesc+".log", fileFlag, filePermission)
	if err != nil {
		fmt.Printf("Error Opening File: %v", err)
	}
	var baseLogger = logrus.New()
	var standardLogger = &MainLogger{baseLogger, moduleName, moduleDesc, functionName}
	if GetFormat() == nil {
		SetFormat("time", "logLevel", "msg")
	}
	standardLogger.Formatter = GetFormat()
	standardLogger.ModuleName = moduleName
	standardLogger.ModuleDesc = moduleDesc
	standardLogger.FunctionName = functionName
	mw := io.MultiWriter(os.Stdout, f)
	standardLogger.SetOutput(mw)
	return standardLogger
}
