package logger

import (
	"github.com/geekshacking/geekhub-backend/config"
	"go.uber.org/zap"
)

type Logger struct {
	Zap *zap.SugaredLogger
}

func NewLogger(config config.Config) (Logger, error) {
	var l *zap.Logger
	var err error

	if config.IsProduction() {
		l, err = zap.NewProduction()
	} else {
		l, err = zap.NewDevelopment()
	}

	if err != nil {
		return Logger{}, err
	}

	defer func(logger *zap.Logger) {
		_ = logger.Sync()
	}(l)

	return Logger{
		Zap: l.Sugar(),
	}, nil
}
