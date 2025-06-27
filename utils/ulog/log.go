package ulog

import (
	"bytes"

	"github.com/sirupsen/logrus"
)

type _logFormat struct {
}

func (l *_logFormat) Format(entry *logrus.Entry) ([]byte, error) {
	b := entry.Buffer
	if b == nil {
		b = &bytes.Buffer{}
	}

	b.WriteRune('[')
	b.WriteString(entry.Time.Format("2006-01-02 15:04:05"))
	b.WriteRune(']')

	b.WriteRune(' ')

	b.WriteRune('[')
	b.WriteString(entry.Level.String())
	b.WriteRune(']')

	b.WriteRune(' ')
	b.WriteString(entry.Message)

	b.WriteRune('\n')

	return b.Bytes(), nil
}

func New() *logrus.Logger {
	logger := logrus.New()
	logger.SetLevel(logrus.DebugLevel)
	logger.SetOutput(&NullWriter{})

	logger.SetFormatter(&_logFormat{})
	// logger.SetFormatter(&logrus.TextFormatter{
	// 	DisableColors:   true,
	// 	TimestampFormat: "2006-01-02 15:04:05",
	// })

	return logger
}
