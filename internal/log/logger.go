package log

import "context"

type Logger interface {
	Info(context.Context, string)
	Infof(context.Context, string, ...any)

	Debug(context.Context, string)
	Debugf(context.Context, string, ...any)

	Error(context.Context, string)
	Errorf(context.Context, string, ...any)
}
