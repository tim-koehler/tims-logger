# TimsLogger [![Tests](https://github.com/tim-koehler/tims-logger/actions/workflows/go.yml/badge.svg)](https://github.com/tim-koehler/tims-logger/actions/workflows/go.yml) [![Coverage Status](https://coveralls.io/repos/github/tim-koehler/tims-logger/badge.svg)](https://coveralls.io/github/tim-koehler/tims-logger) [![Go Report Card](https://goreportcard.com/badge/github.com/tim-koehler/tims-logger)](https://goreportcard.com/report/github.com/tim-koehler/tims-logger)

A logger build for my personal needs. Maybe you will also enjoy it :)

## Install

```bash
go get -u github.com/tim-koehler/tims-logger
```

## Example

```go
package main

import logger "github.com/tim-koehler/tims-logger"

func main() {
	logger.SetColoredLogs(true)
	logger.SetLogLevel(logger.DEBUG.String())

	// TEXT Logging
	logger.SetLogType(logger.TEXT.String()) // logger.SetLogType("TEXT")
	logger.Debugln("Debug log")
	logger.Infoln("Info log")
	logger.Warningln("Warning log")
	logger.Errorln("Error log")

	// JSON Logging
	logger.SetLogType(logger.JSON.String()) // logger.SetLogType("JSON")
	logger.Debugln("Some error")
	logger.Infoln(logger.Custom{
		"foo": "bar",
		"baz": 123,
	})
}
```

![image](https://i.imgur.com/norZIPN.png)
