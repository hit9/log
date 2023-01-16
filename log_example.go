// Copyright 2018 Chao Wang <hit9@icloud.com>

// +build ignore

package main

import (
	"github.com/hit9/log"
)

var logger = log.Get("ExampleName")

func main() {
	logger.SetLevel(log.INFO)
	logger.Debug("This is a debug message")
	logger.Info("This is a info message")
	logger.Warn("This is a warning message")
	logger.Error("This is an error message")
	logger.Warn("This is a number %v", 1)
}
