package log

import (
	"github.com/sirupsen/logrus"
)

var jsonFormat *logrus.JSONFormatter

// SetFormat to set desired key name for the given key
func SetFormat(timeStamp string, logLevel string, message string) {
	jsonFormat = &logrus.JSONFormatter{
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  timeStamp,
			logrus.FieldKeyLevel: logLevel,
			logrus.FieldKeyMsg:   message,
		},
	}
}

// GetFormat function to get the format of json
func GetFormat() *logrus.JSONFormatter {
	return jsonFormat
}
