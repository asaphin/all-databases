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
