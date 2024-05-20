package utils

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/jszwec/csvutil"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"io"
	"math/rand"
	"os"
	"runtime"
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

func must(err error) {
	if err != nil {
		_, file, no, _ := runtime.Caller(1)
		log.WithError(err).WithField("caller", fmt.Sprintf("%s:%d", file, no)).Fatal("must-run error")
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

func UnmarshalCSVFromFile[T any](path string, divider rune, rowType T) ([]T, error) {
	var decodedContent []T

	file, err := os.Open(path)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to open file %s", path)
	}

	defer func() { LogAsWarningIfError(file.Close()) }()

	csvReader := csv.NewReader(file)
	csvReader.Comma = divider

	decoder, err := csvutil.NewDecoder(csvReader)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to parse csv file %s", path)
	}

	for {
		var decodedContentItem T
		if err = decoder.Decode(&decodedContentItem); err == io.EOF {
			break
		} else if err != nil {
			return nil, errors.Wrapf(err, "unable to parse csv file row %s", path)
		}

		decodedContent = append(decodedContent, decodedContentItem)
	}

	return decodedContent, nil
}

func MustUnmarshalCSVFromFile[T any](path string, divider rune, rowType T) []T {
	content, err := UnmarshalCSVFromFile(path, divider, rowType)
	must(err)
	return content
}

func GetRandomElement[T any](elements []T) T {
	if len(elements) == 0 {
		var emptyValue T

		return emptyValue
	}

	return elements[rand.Intn(len(elements))]
}
