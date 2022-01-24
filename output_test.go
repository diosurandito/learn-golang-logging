package golang_logging

import (
	"os"
	"testing"

	"github.com/sirupsen/logrus"
)

func Test(t *testing.T) {
	logger := logrus.New()

	file, _ := os.OpenFile("application.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	logger.SetOutput(file)

	logger.Info("Hello Logging(Info)")
	logger.Warn("Hello Logging(Warn)")
	logger.Error("Hello Logging(Error)")
}
