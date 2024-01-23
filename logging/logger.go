package logging

import (
	"github.com/TechSir3n/analytics-platform/assistance"
	"github.com/sirupsen/logrus"
	"os"
)

var Log *logrus.Logger

func init() {
	Log = logrus.New() 
	file, err := os.OpenFile(assistance.LogFile, os.O_APPEND | os.O_CREATE | os.O_RDWR, 0600)
	if err != nil {
		panic(err)
	}
	
	Log.SetOutput(file)
	Log.SetLevel(logrus.DebugLevel)
	Log.SetReportCaller(true)

	Log.SetFormatter(&logrus.JSONFormatter{})
}
