package log

//go:generate mockgen -source=interfaces.go -destination=mocks/log.go -package=mock_log

type SugaredLogger interface {
	Debug(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
	Fatal(args ...interface{})
}
