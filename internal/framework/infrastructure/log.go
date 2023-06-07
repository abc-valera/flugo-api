package infrastructure

import (
	"os"

	"github.com/abc-valera/flugo-api/internal/domain/service"
	"github.com/charmbracelet/log"
)

type charmLogger struct {
	logger *log.Logger
}

func newLogger() service.Logger {
	return &charmLogger{
		logger: log.NewWithOptions(os.Stdout, log.Options{
			Level:           log.DebugLevel,
			ReportTimestamp: true,
			// Formatter:       log.JSONFormatter,
		}),
	}
}

func (s *charmLogger) Debug(msg interface{}, keyvals ...interface{}) {
	s.logger.Debug(msg, keyvals...)
}

func (s *charmLogger) Debugf(format string, args ...interface{}) {
	s.logger.Debugf(format, args...)
}

func (s *charmLogger) Info(msg interface{}, keyvals ...interface{}) {
	s.logger.Info(msg, keyvals...)
}

func (s *charmLogger) Infof(format string, args ...interface{}) {
	s.logger.Infof(format, args...)
}

func (s *charmLogger) Warn(msg interface{}, keyvals ...interface{}) {
	s.logger.Warn(msg, keyvals...)
}

func (s *charmLogger) Warnf(format string, args ...interface{}) {
	s.logger.Warnf(format, args...)
}

func (s *charmLogger) Error(msg interface{}, keyvals ...interface{}) {
	s.logger.Error(msg, keyvals...)
}

func (s *charmLogger) Errorf(format string, args ...interface{}) {
	s.logger.Errorf(format, args...)
}

func (s *charmLogger) Fatal(msg interface{}, keyvals ...interface{}) {
	s.logger.Fatal(msg, keyvals...)
}

func (s *charmLogger) Fatalf(format string, args ...interface{}) {
	s.logger.Fatalf(format, args...)
}
