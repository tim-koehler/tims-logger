# TimsLogger

A logger build for my personal needs. Maybe you will also enjoy it :)

## Example

```go
package main

import "github.com/tim-koehler/TimsLogger/logger"

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
