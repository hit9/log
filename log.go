// Package log implements a minimal leveled logging library.
//
// Example
//
//	var logger = log.Get("ExampleName")
//
//	func main() {
//		logger.SetLevel(log.INFO)
//		logger.Debug("This is a debug message")
//		logger.Info("This is a info message")
//		logger.Warn("This is a warning message")
//		logger.Error("This is an error message")
//		logger.Warnf("This is a number %v", 1)
//	}
package log

import (
	"fmt"
	"io"
	"os"
	"path"
	"runtime"
	"sync"
	"time"

	"github.com/willf/pad"
)

// Level
const (
	DEBUG int = iota
	INFO
	WARN
	ERROR
)

// default
const (
	DefaultCallerDepth = 2
)

// Global registry
var m = sync.Mutex{}
var registry = make(map[string]*Logger, 0)

// Level name
var levelNames = [4]string{"DEBUG", "INFO", "WARN", "ERROR"}

// Logger abstraction.
type Logger struct {
	name                      string
	nameLength                int
	level                     int
	w                         io.Writer
	colored                   bool
	callerDepth               int
	enableCallerSourceLogging bool
}

// get log registry
func GetRegistry() map[string]*Logger {
	return registry
}

// New creates a new Logger.
func Get(name string) *Logger {
	m.Lock()
	defer m.Unlock()
	l, ok := registry[name]
	if ok {
		return l
	}
	l = &Logger{
		name:                      name,
		nameLength:                0,
		level:                     INFO,
		w:                         os.Stdout,
		colored:                   true,
		callerDepth:               DefaultCallerDepth,
		enableCallerSourceLogging: true,
	}
	registry[name] = l
	return l
}

// colors to ansi code map
var colors = map[string]int{
	"black":   0,
	"red":     1,
	"green":   2,
	"yellow":  3,
	"blue":    4,
	"magenta": 5,
	"cyan":    6,
	"white":   7,
}

// levelColors
var levelColors = map[int]string{
	DEBUG: "blue",
	INFO:  "green",
	WARN:  "yellow",
	ERROR: "red",
}

// SetColored sets the color enability.
func (l *Logger) SetColored(b bool) {
	l.colored = b
}

// SetLevel sets the logging level.
func (l *Logger) SetLevel(level int) {
	l.level = level % len(levelNames)
}

// GetLevel gets the logging level.
func (l *Logger) GetLevel() int {
	return l.level
}

// SetWriter sets the writer.
func (l *Logger) SetWriter(w io.Writer) {
	l.w = w
}

// SetNameLength sets the name length.
func (l *Logger) SetNameLength(n int) { l.nameLength = n }

// SetCallerDepth sets the caller depth for this logger.
func (l *Logger) SetCallerDepth(callerDepth int) {
	l.callerDepth = callerDepth
}

// DisableCallerSourceLogging disables the logging for caller source.
// This sets to true by default.
func (l *Logger) DisableCallerSourceLogging() {
	l.enableCallerSourceLogging = false
}

// Debug formats and logs message with level DEBUG.
func (l *Logger) Debug(format string, a ...interface{}) error {
	return l.log(DEBUG, fmt.Sprintf(format, a...))
}

// Info formats and logs message with level INFO.
func (l *Logger) Info(format string, a ...interface{}) error {
	return l.log(INFO, fmt.Sprintf(format, a...))
}

// Warn formats and logs message with level WARN.
func (l *Logger) Warn(format string, a ...interface{}) error {
	return l.log(WARN, fmt.Sprintf(format, a...))
}

// Error formats and logs message with level ERROR.
func (l *Logger) Error(format string, a ...interface{}) error {
	return l.log(ERROR, fmt.Sprintf(format, a...))
}

// Fatal formats and logs message with level FATAL.
func (l *Logger) Fatal(format string, a ...interface{}) {
	l.log(ERROR, fmt.Sprintf(format, a...))
	os.Exit(1)
}

// Smart formats and logs message in different levels according to given error.
func (l *Logger) Smart(e error, format string, a ...interface{}) {
	if e == nil {
		l.Info(format, a...)
	} else {
		l.Error(format, a...)
	}
}

func (l *Logger) Log(level int, format string, a ...interface{}) error {
	return l.log(level, fmt.Sprintf(format, a...))
}

// Colored returns text in color.
func Colored(color string, text string) string {
	return fmt.Sprintf("\033[3%dm%s\033[0m", colors[color], text)
}

// log dose logging.
func (l *Logger) log(level int, msg string) error {
	if level >= l.level {
		// Caller pkg.
		_, fileName, line, _ := runtime.Caller(l.callerDepth)
		pkgName := path.Base(path.Dir(fileName))
		filepath := path.Join(pkgName, path.Base(fileName))
		// Datetime and pid.
		now := time.Now().String()[:19]
		// Message
		levelName := levelNames[level]
		levelName = pad.Right(levelName, 5, " ")
		// Name
		name := l.name
		if l.nameLength > 0 {
			name = pad.Right(l.name, l.nameLength, " ")
		}
		// Whether to log the caller source.
		var headerString string
		if !l.enableCallerSourceLogging {
			headerString = fmt.Sprintf("[ %s ] %s %s", name, levelName, now)
		} else {
			headerString = fmt.Sprintf("[ %s ] %s %s %s:%d", name, levelName, now, filepath, line)
		}

		header := headerString
		if l.colored {
			header = Colored(levelColors[level], headerString)
		}

		_, err := fmt.Fprintf(l.w, "%s %s\n", header, msg)
		return err
	}
	return nil
}
