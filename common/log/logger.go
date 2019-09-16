package log

// Level loglevel
type Level int

const (
	// Debug debug
	Debug Level = iota
	// Trace trace
	Trace
	// Info info
	Info
	// Warn warn
	Warn
	// Error error
	Error
)

func (rm Level) String() string {
	switch rm {
	case Debug:
		return "Debug"
	case Trace:
		return "Trace"
	case Info:
		return "Info"
	case Warn:
		return "Warn"
	case Error:
		return "Error"
	default:
		return "Unknown"
	}
}

// LoggerImpl logwrite impl
type LoggerImpl interface {
	LogWrite(lv Level, msg string)
	LogWriteWithObj(lv Level, obj interface{})
	LogWriteWithMsgAndObj(lv Level, msg string, obj interface{})
}
