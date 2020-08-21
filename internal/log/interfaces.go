package log

//go:generate mockgen -source=interfaces.go -destination=mocks/log.go -package=mock_log

type SugaredLogger interface {
	Debugw(string, ...interface{})
	Infow(string, ...interface{})
	Warnw(string, ...interface{})
	Errorw(string, ...interface{})
	Fatalw(string, ...interface{})
}
