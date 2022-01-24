package golang_logging

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestSingleton(t *testing.T) {
	//Tidak Disarankan menggunakan singleton seperti di bawah
	logrus.Info("Hello Info")
	logrus.Warn("Hello Warn")

	logrus.SetFormatter(&logrus.JSONFormatter{})

	logrus.Info("Hello Info JSON")
	logrus.Warn("Hello Warn JSON")
}
