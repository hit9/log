Log
===

Package log implements leveled logging.

https://godoc.org/github.com/hit9/log

Example
-------

```go
package main

import (
	"github.com/hit9/log"
)

func main() {
	log.SetLevel(log.INFO)
	log.Debug("This is a debug message")
	log.Info("This is a info message")
	log.Warn("This is a warning message")
	log.Error("This is an error message")
	log.Warnf("This is a number %v", 1)
}
```

License
-------

BSD.
