package log

import (
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

var (
	alreadyExistMsg               = "Already Exist"
	notFoundMessage               = "Not Found"
	incorrectNoOfArgsMsg          = "Incorrect No of Args. Expecting"
	logFormatMsg                  = "Set log format"
	jsonMarshalFailedMsg          = "Json Marshal Failed"
	jsonUnMarshalFailedMsg        = "Json UnMarshal Failed"
	successMsg                    = "Success"
	invalidFunctionNameMsg        = "Invalid Function Name"
	errorWritingOnLedgerMsg       = "Error Writing on ledger"
	errorFetchingDataMsg          = "Error Fetching data"
	failedToGetCurrentDateTimeMsg = "Failed to get current date and time"
	dataFetchSuccessMsg           = "Data fetch successful"
	dataEmptyMsg                  = "Empty Data"
	generateRandomNoFailedMsg     = "Failed to generate random number"
)

/* ==============================ERROR MESSAGES============================== */

// EmptyList Standard eror message
func (m *MainLogger) EmptyList(message string) {
	m.ExecLog(INFO, WrapMessage(dataEmptyMsg, message))
}

// NotFound Standard error message
func (m *MainLogger) NotFound(customMessage string) {
	m.ExecLog(ERROR, WrapMessage(notFoundMessage, customMessage))
}

// IncorrectNoOfArgs Standard error message
func (m *MainLogger) IncorrectNoOfArgs(expectedValue string) {
	m.ExecLog(ERROR, WrapMessage(incorrectNoOfArgsMsg, expectedValue))
}

// AlreadyExist Standard error message
func (m *MainLogger) AlreadyExist(customMessage string) {
	m.ExecLog(ERROR, WrapMessage(alreadyExistMsg, customMessage))
}

// JSONMarshalFailed error message
func (m *MainLogger) JSONMarshalFailed(err error, message string) {
	m.ExecLog(ERROR, errors.Wrap(err, WrapMessage(jsonMarshalFailedMsg, message)))
}

// JSONUnMarshalFailed error message
func (m *MainLogger) JSONUnMarshalFailed(err error, message string) {
	m.ExecLog(ERROR, errors.Wrap(err, WrapMessage(jsonUnMarshalFailedMsg, message)))
}

// InvalidFunctionName error message
func (m *MainLogger) InvalidFunctionName(functionName string) {
	m.ExecLog(ERROR, WrapMessage(invalidFunctionNameMsg, functionName))
}

// ErrorWritingOnLedger error message
func (m *MainLogger) ErrorWritingOnLedger(err error, message string) {
	m.ExecLog(ERROR, errors.Wrap(err, WrapMessage(errorWritingOnLedgerMsg, message)))
}

// ErrorFetchingData error message
func (m *MainLogger) ErrorFetchingData(err error, message string) {
	m.ExecLog(ERROR, errors.Wrap(err, WrapMessage(errorFetchingDataMsg, message)))
}

// GetCurrentDateTimeFailed error message
func (m *MainLogger) GetCurrentDateTimeFailed(err error) {
	m.ExecLog(ERROR, errors.Wrap(err, failedToGetCurrentDateTimeMsg))
}

// GenerateRandomNoFailed error message
func (m *MainLogger) GenerateRandomNoFailed(err error, message string) {
	m.ExecLog(ERROR, errors.Wrap(err, WrapMessage(generateRandomNoFailedMsg, message)))
}

/* ==============================ERROR MESSAGES============================== */

/* ==============================INFO MESSAGES============================== */

// LogFormat Info message
func (m *MainLogger) LogFormat(format *logrus.JSONFormatter) {
	m.ExecLog(INFO, logFormatMsg, format.FieldMap)
}

// Success Info Message
func (m *MainLogger) Success(customMessage string) {
	m.ExecLog(INFO, WrapMessage(successMsg, customMessage))
}

// DataFetchSuccess Info Message
func (m *MainLogger) DataFetchSuccess(customMessage string) {
	m.ExecLog(INFO, WrapMessage(dataFetchSuccessMsg, customMessage))
}

/* ==============================INFO MESSAGES============================== */

/* ==============================CUSTOM MESSAGE============================== */

// CustomMessage Standard message
func (m *MainLogger) CustomMessage(logLevel string, args ...interface{}) {
	m.ExecLog(logLevel, args...)
}

/* ==============================CUSTOM MESSAGE============================== */

// WrapMessage function
func WrapMessage(defaultMessage string, customMessage string) string {
	return defaultMessage + ":" + customMessage
}
