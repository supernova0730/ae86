package logger

import (
	"github.com/sirupsen/logrus"
	"os"
	"sync"
)

var (
	Log  *logrus.Logger
	once sync.Once
)

func Init() {
	once.Do(func() {
		Log = logrus.New()
		Log.SetOutput(os.Stdout)
		Log.SetLevel(logrus.DebugLevel)
		Log.SetFormatter(&logrus.TextFormatter{
			FullTimestamp:          true,
			TimestampFormat:        "2006-01-02 15:04:05.000",
			DisableLevelTruncation: true,
			PadLevelText:           true,
			QuoteEmptyFields:       true,
		})
	})
}
