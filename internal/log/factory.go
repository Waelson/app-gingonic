package log

import "errors"

type TypeLogger int64

const (
	TypeLogrus TypeLogger = 0
)

var (
	InvalidLoggerError = errors.New("invalid logger")
)

func NewLogger(name string) Logger {
	log, _ := New(TypeLogrus, name)
	return log
}

func New(typeLogger TypeLogger, name string) (Logger, error) {
	if typeLogger == TypeLogrus {
		return newLogrus(name), nil
	} else {
		return nil, InvalidLoggerError
	}
}
