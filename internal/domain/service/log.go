package service

type Logger interface {
	Debug(msg interface{}, keyvals ...interface{})
	Debugf(format string, args ...interface{})
	Info(msg interface{}, keyvals ...interface{})
	Infof(format string, args ...interface{})
	Warn(msg interface{}, keyvals ...interface{})
	Warnf(format string, args ...interface{})
	Error(msg interface{}, keyvals ...interface{})
	Errorf(format string, args ...interface{})
	Fatal(msg interface{}, keyvals ...interface{})
	Fatalf(format string, args ...interface{})
}
