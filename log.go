// Copyright 2018 Chao Wang <hit9@icloud.com>

// Package log implements leveled logging.
package log

import (
	"fmt"
	"io"
	"os"
	"path"
	"runtime"
	"sync"
	"time"
)

// Level
const (
	DEBUG int = iota
	INFO
	WARN
	ERROR
)

// Global registry
var m = sync.Mutex{}
var registry = make(map[string]*Logger, 0)

// Level name
var levelNames = [4]string{"DEBUG", "INFO", "WARN", "ERROR"}

// Logger abstraction.
type Logger struct {
	name    string
	level   int
	w       io.Writer
	colored bool
	enabled bool
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
		name:    name,
		level:   INFO,
		w:       os.Stdout,
		colored: true,
		enabled: true,
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

// SetWriter sets the writer.
func (l *Logger) SetWriter(w io.Writer) {
	l.w = w
}

// Disable the logging.
func (l *Logger) Disable() {
	l.enabled = false
}

// Enable the logging.
func (l *Logger) Enable() {
	l.enabled = true
}

// Debug logs message with level DEBUG.
func (l *Logger) Debug(a ...interface{}) error {
	return l.log(DEBUG, fmt.Sprint(a...))
}

// Info logs message with level INFO.
func (l *Logger) Info(a ...interface{}) error {
	return l.log(INFO, fmt.Sprint(a...))
}

// Warn logs message with level WARN.
func (l *Logger) Warn(a ...interface{}) error {
	return l.log(WARN, fmt.Sprint(a...))
}

// Error logs message with level ERROR.
func (l *Logger) Error(a ...interface{}) error {
	return l.log(ERROR, fmt.Sprint(a...))
}

// Fatal and logs message with level FATAL.
func (l *Logger) Fatal(a ...interface{}) {
	l.log(ERROR, fmt.Sprint(a...))
	os.Exit(1)
}

// Debugf formats and logs message with level DEBUG.
func (l *Logger) Debugf(format string, a ...interface{}) error {
	return l.log(DEBUG, fmt.Sprintf(format, a...))
}

// Infof formats and logs message with level INFO.
func (l *Logger) Infof(format string, a ...interface{}) error {
	return l.log(INFO, fmt.Sprintf(format, a...))
}

// Warnf formats and logs message with level WARN.
func (l *Logger) Warnf(format string, a ...interface{}) error {
	return l.log(WARN, fmt.Sprintf(format, a...))
}

// Errorf formats and logs message with level ERROR.
func (l *Logger) Errorf(format string, a ...interface{}) error {
	return l.log(ERROR, fmt.Sprintf(format, a...))
}

// Fatalf formats and logs message with level FATAL.
func (l *Logger) Fatalf(format string, a ...interface{}) {
	l.log(ERROR, fmt.Sprintf(format, a...))
	os.Exit(1)
}

// Colored returns text in color.
func Colored(color string, text string) string {
	return fmt.Sprintf("\033[3%dm%s\033[0m", colors[color], text)
}

// log dose logging.
func (l *Logger) log(level int, msg string) error {
	if l.enabled && level >= l.level {
		// Caller pkg.
		_, fileName, line, _ := runtime.Caller(2)
		pkgName := path.Base(path.Dir(fileName))
		filepath := path.Join(pkgName, path.Base(fileName))
		// Datetime and pid.
		now := time.Now().String()[:19]
		// Message
		levelName := levelNames[level]
		header := Colored(levelColors[level], fmt.Sprintf("[%s] %s %s %s:%d", l.name, levelName, now, filepath, line))
		_, err := fmt.Fprintf(l.w, "%s %s\n", header, msg)
		return err
	}
	return nil
}
