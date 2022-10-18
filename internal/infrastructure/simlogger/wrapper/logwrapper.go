package wrapper

import "go.uber.org/zap"

type Logger interface {
	Info(msg string, fields ...zap.Field)
	Error(msg string, fields ...zap.Field)
	Warn(msg string, fields ...zap.Field)
	Fatal(msg string, fields ...zap.Field)
	Debug(msg string, fields ...zap.Field)
}

type logWrapper struct {
	logger  Logger
	traceID string
	version string
}

func (l *logWrapper) Info(msg string, fields ...zap.Field) {
	f := l.mergeField(fields...)
	l.logger.Info(msg, f...)
}

func (l *logWrapper) Warn(msg string, fields ...zap.Field) {
	f := l.mergeField(fields...)
	l.logger.Warn(msg, f...)
}

func (l *logWrapper) Error(msg string, fields ...zap.Field) {
	f := l.mergeField(fields...)
	l.logger.Error(msg, f...)
}

func (l *logWrapper) Fatal(msg string, fields ...zap.Field) {
	f := l.mergeField(fields...)
	l.logger.Fatal(msg, f...)
}

func (l *logWrapper) Debug(msg string, fields ...zap.Field) {
	f := l.mergeField(fields...)
	l.logger.Debug(msg, f...)
}

func (l *logWrapper) mergeField(fields ...zap.Field) []zap.Field {

	parentID := ""
	s := []zap.Field{
		zap.String("version", l.version),
		zap.String("trace_id", l.traceID),
		zap.String("span_parent_id", parentID),
	}

	s = append(s, fields...)

	return s
}
