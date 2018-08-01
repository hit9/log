// Copyright 2018 Chao Wang <hit9@icloud.com>

package log

import (
	"testing"
)

func TestLog(t *testing.T) {
	// No assertions.
	SetLevel(DEBUG)
	Debug(nil)
	Info(nil)
	Warn(nil)
	Error(nil)
	Debugf("hello %s", "world")
	Infof("hello %s", "world")
	Warnf("hello %s", "world")
	Errorf("hello %s", "world")
}
