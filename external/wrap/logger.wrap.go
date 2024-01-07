package wrap

import (
	"go.uber.org/zap"
)

type ExternalLogger = *zap.SugaredLogger

type InternalLogger struct {
	logger ExternalLogger
}

func NewInternalLogger(logger ExternalLogger) *InternalLogger {
	return &InternalLogger{
		logger,
	}
}

func (il *InternalLogger) Info(msg string) {
	il.logger.Info(msg)
}

func (il *InternalLogger) Error(err error) error {
	if err != nil {
		il.logger.Error(err.Error())
	}
	return err
}

func (il *InternalLogger) Warning(msg string) {
	il.logger.Warn(msg)
}
