package utils

import "github.com/sirupsen/logrus"

func PanicOnError(err error, msg string) {
	if err != nil {
		logrus.Panicf("%s: %s", msg, err)
	}
}
