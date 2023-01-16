// Copyright 2018 Chao Wang <hit9@icloud.com>

package log

import (
	"testing"
)

func TestLog(t *testing.T) {
	logger := Get("ExampleName")
	// No assertions.
	logger.SetLevel(DEBUG)
	logger.Debug("hello %s", "world")
	logger.Info("hello %s", "world")
	logger.Warn("hello %s", "world")
	logger.Error("hello %s", "world")
}
