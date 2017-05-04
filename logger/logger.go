package logger

import (
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/pkg/errors"
)

var logOut = logrus.New()
var logErr = logrus.New()

//InitLogger ...
func InitLogger() {
	logOut.Formatter = &logrus.JSONFormatter{}
	logOut.Level = logrus.DebugLevel
	logOut.Out = os.Stdout

	logErr.Formatter = &logrus.JSONFormatter{}
	logErr.Level = logrus.ErrorLevel
	logErr.Out = os.Stderr
}

//Debug ...
func Debug(value ...interface{}) {
	logOut.Debug(value)
}

//Info ...
func Info(value ...interface{}) {
	logOut.Info(value)
}

//Error ...
func Error(err error, msg string) {
	logErr.Error(errors.Wrap(err, msg).Error())
}

//Fatal ...
func Fatal(err error, msg string) {
	logErr.Fatal(errors.Wrap(err, msg).Error())
}
