// Copyright 2018 Chao Wang <hit9@icloud.com>

package log

import (
	"testing"
)

func TestLog(t *testing.T) {
	logger := Get("ExampleName")
	// No assertions.
	logger.SetLevel(DEBUG)
	logger.Debug(nil)
	logger.Info(nil)
	logger.Warn(nil)
	logger.Error(nil)
	logger.Debugf("hello %s", "world")
	logger.Infof("hello %s", "world")
	logger.Warnf("hello %s", "world")
	logger.Errorf("hello %s", "world")
}
