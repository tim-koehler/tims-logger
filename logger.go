package logger

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

const (
	JSON LogType = "JSON"
	TEXT LogType = "TEXT"
)

const (
	DEBUG   LogLevel = 0
	INFO    LogLevel = 1
	WARNING LogLevel = 2
	ERROR   LogLevel = 3
)

type LogType string
type LogLevel int
type Custom map[string]interface{}

var logLevel LogLevel = DEBUG
var logType LogType = TEXT
var colors bool = true
var dateFormat = "2006-01-02T15:04:05.515-07:00"

func SetLogLevel(strLogFormat string) {
	switch strings.ToUpper(strLogFormat) {
	case DEBUG.String():
		logLevel = DEBUG
	case INFO.String():
		logLevel = INFO
	case WARNING.String():
		logLevel = WARNING
	case ERROR.String():
		logLevel = ERROR
	default:
		logLevel = DEBUG
	}
}

func GetLogLevel() LogLevel {
	return logLevel
}

func SetLogType(strLogType string) {
	switch strings.ToUpper(strLogType) {
	case TEXT.String():
		logType = TEXT
	case JSON.String():
		logType = JSON
	default:
		logType = TEXT
	}
}

func GetLogType() LogType {
	return logType
}

func SetDateFormat(newDateFormat string) {
	dateFormat = newDateFormat
}

func GetDateFormat() string {
	return dateFormat
}

func SetColoredLogs(enabled bool) {
	colors = enabled
}

func Debugln(v ...interface{}) {
	if !checkLogLevel(DEBUG) {
		return
	}
	if logType == TEXT {
		coloredPrefix := getColoredPrefix(DEBUG)
		fmt.Println(getTime(), coloredPrefix, removeBrackets(v))
	} else {
		fmt.Println(buildJsonLog(DEBUG, v))
	}
}

func Debugf(formatString string, v ...interface{}) {
	if !checkLogLevel(DEBUG) {
		return
	}
	if logType == TEXT {
		str := fmt.Sprintf(formatString, v...)
		coloredPrefix := getColoredPrefix(DEBUG)
		fmt.Println(getTime(), coloredPrefix, str)
	} else {
		str := fmt.Sprintf(formatString, v...)
		fmt.Println(buildJsonLog(DEBUG, []interface{}{str}))
	}
}

func Infoln(v ...interface{}) {
	if !checkLogLevel(INFO) {
		return
	}
	if logType == TEXT {
		coloredPrefix := getColoredPrefix(INFO)
		fmt.Println(getTime(), coloredPrefix, removeBrackets(v))
	} else {
		fmt.Println(buildJsonLog(INFO, v))
	}
}

func Infof(formatString string, v ...interface{}) {
	if !checkLogLevel(INFO) {
		return
	}
	if logType == TEXT {
		str := fmt.Sprintf(formatString, v...)
		coloredPrefix := getColoredPrefix(INFO)
		fmt.Println(getTime(), coloredPrefix, str)
	} else {
		str := fmt.Sprintf(formatString, v...)
		fmt.Println(buildJsonLog(INFO, []interface{}{str}))
	}
}

func Warningln(v ...interface{}) {
	if !checkLogLevel(WARNING) {
		return
	}
	if logType == TEXT {
		coloredPrefix := getColoredPrefix(WARNING)
		fmt.Println(getTime(), coloredPrefix, removeBrackets(v))
	} else {
		fmt.Println(buildJsonLog(WARNING, v))
	}
}

func Warningf(formatString string, v ...interface{}) {
	if !checkLogLevel(WARNING) {
		return
	}
	if logType == TEXT {
		str := fmt.Sprintf(formatString, v...)
		coloredPrefix := getColoredPrefix(WARNING)
		fmt.Println(getTime(), coloredPrefix, str)
	} else {
		str := fmt.Sprintf(formatString, v...)
		fmt.Println(buildJsonLog(WARNING, []interface{}{str}))
	}
}

func Errorln(v ...interface{}) {
	if !checkLogLevel(ERROR) {
		return
	}
	if logType == TEXT {
		coloredPrefix := getColoredPrefix(ERROR)
		fmt.Println(getTime(), coloredPrefix, removeBrackets(v))
	} else {
		fmt.Println(buildJsonLog(ERROR, v))
	}
}

func Errorf(formatString string, v ...interface{}) {
	if !checkLogLevel(ERROR) {
		return
	}
	if logType == TEXT {
		str := fmt.Sprintf(formatString, v...)
		coloredPrefix := getColoredPrefix(ERROR)
		fmt.Println(getTime(), coloredPrefix, str)
	} else {
		str := fmt.Sprintf(formatString, v...)
		fmt.Println(buildJsonLog(ERROR, []interface{}{str}))
	}
}

func CreateString(level LogLevel, formatString string, v ...interface{}) string {
	result := ""
	if logType == TEXT {
		str := fmt.Sprintf(formatString, v...)
		coloredPrefix := getColoredPrefix(level)
		result = fmt.Sprintf("%s %s %s", getTime(), coloredPrefix, str)
	} else {
		str := fmt.Sprintf(formatString, v...)
		result = fmt.Sprintln(buildJsonLog(level, []interface{}{str}))
	}
	return result
}

func checkLogLevel(t LogLevel) bool {
	return t >= logLevel
}

func getTime() string {
	return fmt.Sprintf("%30s", time.Now().Format(dateFormat))
}

func buildJsonLog(lvl LogLevel, v []interface{}) string {
	var data []byte
	for _, val := range v {
		if customMap, ok := val.(Custom); ok {
			combinedMap := map[string]string{}
			for key, val := range customMap {
				combinedMap[key] = fmt.Sprintf("%v", val)
			}

			combinedMap["@timestamp"] = getTime()
			combinedMap["level"] = lvl.String()

			data, _ := json.Marshal(combinedMap)
			return string(data)
		}
	}
	data, _ = json.Marshal(
		map[string]string{
			"@timestamp": getTime(),
			"level":      lvl.String(),
			"message":    removeBrackets(v),
		})
	return string(data)
}

func removeBrackets(v []interface{}) string {
	str := fmt.Sprintf("%v", v)
	str = strings.TrimPrefix(str, "[")
	return strings.TrimSuffix(str, "]")
}

func getColoredPrefix(t LogLevel) string {
	if !colors {
		return fmt.Sprintf("%-10s|", fmt.Sprintf("[%s]", t.String()))
	}

	switch t {
	case INFO:
		return fmt.Sprintf("\033[1;34m%-10s|\033[0m", fmt.Sprintf("[%s]", t.String()))
	case WARNING:
		return fmt.Sprintf("\033[1;33m%-10s|\033[0m", fmt.Sprintf("[%s]", t.String()))
	case ERROR:
		return fmt.Sprintf("\033[1;31m%-10s|\033[0m", fmt.Sprintf("[%s]", t.String()))
	default: // DEBUG
		return fmt.Sprintf("\033[0;36m%-10s|\033[0m", fmt.Sprintf("[%s]", t.String()))
	}
}

func (t LogLevel) String() string {
	mapping := map[LogLevel]string{
		DEBUG:   "DEBUG",
		INFO:    "INFO",
		WARNING: "WARNING",
		ERROR:   "ERROR",
	}
	return mapping[t]
}

func (t LogType) String() string {
	mapping := map[LogType]string{
		TEXT: "TEXT",
		JSON: "JSON",
	}
	return mapping[t]
}
