package utils

import (
	"encoding/json"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"io"
	"os"
)

func LogAsWarningIfError(err error, args ...interface{}) {
	if err != nil {
		log.WithError(err).Warning(args...)
	}
}

func LogAsErrorIfError(err error, args ...interface{}) {
	if err != nil {
		log.WithError(err).Error(args...)
	}
}

func LogAsWarningIfReturnsError(f func() error, args ...interface{}) {
	err := f()

	if err != nil {
		log.WithError(err).Warning(args...)
	}
}

func LogAsErrorIfReturnsError(f func() error, args ...interface{}) {
	err := f()

	if err != nil {
		log.WithError(err).Error(args...)
	}
}

func Must(err error) {
	if err != nil {
		panic(err)
	}
}

func ReadFile(path string) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer func() { LogAsWarningIfError(f.Close()) }()

	data, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func UnmarshalJSONFromFile(path string, v any) error {
	jsonByteValue, err := ReadFile(path)
	if err != nil {
		return errors.Wrapf(err, "couldn't read json file %s", path)
	}

	if err = json.Unmarshal(jsonByteValue, v); err != nil {
		return errors.Wrapf(err, "couldn't unmarshal json file %s", path)
	}

	return nil
}
