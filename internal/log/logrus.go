package log

import (
	"context"
	"github.com/sirupsen/logrus"
	"os"
)

type instance struct {
	Logger
	Name   string
	Logrus *logrus.Logger
}

func (a *instance) Info(ctx context.Context, value string) {
	a.Logrus.WithField("name", a.Name).Info(value)
}

func (a *instance) Infof(ctx context.Context, value string, params ...any) {
	a.Logrus.WithField("name", a.Name).Infof(value, params)
}

func (a *instance) Debug(ctx context.Context, value string) {
	a.Logrus.WithField("name", a.Name).Debug(value)
}

func (a *instance) Debugf(ctx context.Context, value string, params ...any) {
	a.Logrus.WithField("name", a.Name).Debugf(value, params)
}

func (a *instance) Error(ctx context.Context, value string) {
	a.Logrus.WithField("name", a.Name).Error(value)
}

func (a *instance) Errorf(ctx context.Context, value string, params ...any) {
	a.Logrus.WithField("name", a.Name).Errorf(value, params)
}

func newLogrus(name string) Logger {
	logLogrus := logrus.New()
	logLogrus.Out = os.Stdout

	logLogrus.SetFormatter(&logrus.TextFormatter{FullTimestamp: true, DisableColors: false})

	return &instance{
		Name:   name,
		Logrus: logLogrus,
	}
}
