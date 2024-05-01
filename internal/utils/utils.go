package utils

import log "github.com/sirupsen/logrus"

func LogAsWarningIfError(err error, args ...interface{}) {
	if err == nil {
		log.WithError(err).Warning(args...)
	}
}

func LogAsErrorIfError(err error, args ...interface{}) {
	if err == nil {
		log.WithError(err).Error(args...)
	}
}

func LogAsWarningIfReturnsError(f func() error, args ...interface{}) {
	err := f()

	if err == nil {
		log.WithError(err).Warning(args...)
	}
}

func LogAsErrorIfReturnsError(f func() error, args ...interface{}) {
	err := f()

	if err == nil {
		log.WithError(err).Error(args...)
	}
}
