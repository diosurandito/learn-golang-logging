package golang_logging

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestField(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.WithField("username", "kjx").Info("Hello world")

	logger.WithField("username", "dio").
		WithField("name", "Dio Surandito").
		Info("Hello world")

}

func TestFields(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.WithFields(logrus.Fields{
		"username": "diosurandito",
		"name":     "Dio Surandito",
	}).Info("Hello world")

}
