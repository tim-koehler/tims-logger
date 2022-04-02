package logger

import (
	"io"
	"os"
	"strings"
	"testing"
)

func TestSetLogLevel(t *testing.T) {
	tests := []struct {
		name         string
		strLogFormat string
		shouldLevel  int
	}{
		{DEBUG.String(), "DEBUG", 0},
		{INFO.String(), "info", 1},
		{WARNING.String(), "wArnInG", 2},
		{ERROR.String(), "errOR", 3},
		{"Wrong Level String", "foobar", 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logLevel = 0
			if SetLogLevel(tt.strLogFormat); int(logLevel) != tt.shouldLevel {
				t.Error("should:", tt.shouldLevel, "is:", logLevel)
			}
		})
	}
}

func TestGetLogLevel(t *testing.T) {
	SetLogLevel(WARNING.String())
	if GetLogLevel() != WARNING {
		t.Error("Does not match")
	}
	SetLogLevel(DEBUG.String())
	if GetLogLevel() != DEBUG {
		t.Error("Does not match")
	}
}

func TestSetLogType(t *testing.T) {
	tests := []struct {
		name       string
		strLogType string
		shouldType string
	}{
		{TEXT.String(), "text", "TEXT"},
		{JSON.String(), "JSON", "JSON"},
		{"Wrong Type String", "foobar", "TEXT"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if SetLogType(tt.strLogType); logType.String() != tt.shouldType {
				t.Error("should:", tt.shouldType, "is:", logType)
			}
		})
	}
}

func TestGetLogType(t *testing.T) {
	SetLogType("jsOn")
	if GetLogType() != JSON {
		t.Error("Does not match")
	}
	SetLogType("TExt")
	if GetLogType() != TEXT {
		t.Error("Does not match")
	}
}

func TestSetColoredLogs(t *testing.T) {
	colors = false
	if SetColoredLogs(true); colors != true {
		t.Error("should be true")
	}
	colors = true
	if SetColoredLogs(false); colors != false {
		t.Error("should be false")
	}
}

func TestAllPrints(t *testing.T) {
	tests := []struct {
		name                   string
		printFunction          func()
		shouldOutputTextSuffix string
		shouldOutputJsonSuffix string
	}{
		{"Debugln", func() { Debugln("Foo", "Bar") }, "[DEBUG]   | Foo Bar\n", `"level":"DEBUG","message":"Foo Bar"}` + "\n"},
		{"Debugln", func() { Debugln("Foo and Bar") }, "[DEBUG]   | Foo and Bar\n", `"level":"DEBUG","message":"Foo and Bar"}` + "\n"},
		{"Debugf", func() { Debugf("%s %s", "Foo", "Bar") }, "[DEBUG]   | Foo Bar\n", `"level":"DEBUG","message":"Foo Bar"}` + "\n"},
		{"Debugf", func() { Debugf("%s and %s", "Foo", "Bar") }, "[DEBUG]   | Foo and Bar\n", `"level":"DEBUG","message":"Foo and Bar"}` + "\n"},

		{"Infoln", func() { Infoln("Foo", "Bar") }, "[INFO]    | Foo Bar\n", `"level":"INFO","message":"Foo Bar"}` + "\n"},
		{"Infoln", func() { Infoln("Foo and Bar") }, "[INFO]    | Foo and Bar\n", `"level":"INFO","message":"Foo and Bar"}` + "\n"},
		{"Infof", func() { Infof("%s %s", "Foo", "Bar") }, "[INFO]    | Foo Bar\n", `"level":"INFO","message":"Foo Bar"}` + "\n"},
		{"Infof", func() { Infof("%s and %s", "Foo", "Bar") }, "[INFO]    | Foo and Bar\n", `"level":"INFO","message":"Foo and Bar"}` + "\n"},

		{"Warningln", func() { Warningln("Foo", "Bar", "BAZ") }, "[WARNING] | Foo Bar BAZ\n", `"level":"WARNING","message":"Foo Bar BAZ"}` + "\n"},
		{"Warningln", func() { Warningln("Foo and Bar + BAZ") }, "[WARNING] | Foo and Bar + BAZ\n", `"level":"WARNING","message":"Foo and Bar + BAZ"}` + "\n"},
		{"Warningf", func() { Warningf("%s%s %s", "Foo", "Bar", "BAZ") }, "[WARNING] | FooBar BAZ\n", `"level":"WARNING","message":"FooBar BAZ"}` + "\n"},
		{"Warningf", func() { Warningf("%s and %s & %s", "Foo", "Bar", "BAZ") }, "[WARNING] | Foo and Bar & BAZ\n", `"level":"WARNING","message":"Foo and Bar \u0026 BAZ"}` + "\n"},

		{"Errorln", func() { Errorln("Foo", "Bar", "BAZ") }, "[ERROR]   | Foo Bar BAZ\n", `"level":"ERROR","message":"Foo Bar BAZ"}` + "\n"},
		{"Errorln", func() { Errorln("Foo and Bar + BAZ") }, "[ERROR]   | Foo and Bar + BAZ\n", `"level":"ERROR","message":"Foo and Bar + BAZ"}` + "\n"},
		{"Errorf", func() { Errorf("%s%s %s", "Foo", "Bar", "BAZ") }, "[ERROR]   | FooBar BAZ\n", `"level":"ERROR","message":"FooBar BAZ"}` + "\n"},
		{"Errorf", func() { Errorf("%s and %s & %s", "Foo", "Bar", "BAZ") }, "[ERROR]   | Foo and Bar & BAZ\n", `"level":"ERROR","message":"Foo and Bar \u0026 BAZ"}` + "\n"},
	}

	rescueStdout := os.Stdout
	SetColoredLogs(false)

	for _, tt := range tests {
		SetLogType(TEXT.String())
		t.Run("[TEXT]"+" "+tt.name, func(t *testing.T) {
			r, w, err := os.Pipe()
			if err != nil {
				t.Error(err.Error())
			}

			os.Stdout = w
			tt.printFunction()
			w.Close()

			isOutput, err := io.ReadAll(r)
			if err != nil {
				t.Error(err.Error())
			}

			if !strings.HasSuffix(string(isOutput), tt.shouldOutputTextSuffix) {
				os.Stdout = rescueStdout
				t.Errorf("Expected '%s', got '%s'", tt.shouldOutputTextSuffix, string(isOutput))
			}
		})
		SetLogType(JSON.String())
		t.Run("[JSON]"+" "+tt.name, func(t *testing.T) {
			r, w, err := os.Pipe()
			if err != nil {
				t.Error(err.Error())
			}

			os.Stdout = w
			tt.printFunction()
			w.Close()

			isOutput, err := io.ReadAll(r)
			if err != nil {
				t.Error(err.Error())
			}

			if !strings.HasSuffix(string(isOutput), tt.shouldOutputJsonSuffix) {
				os.Stdout = rescueStdout
				t.Errorf("Expected '%s', got '%s'", tt.shouldOutputJsonSuffix, string(isOutput))
			}
		})
	}
	os.Stdout = rescueStdout
	SetLogType(TEXT.String())
}

func TestCheckLogLevel(t *testing.T) {
	tests := []struct {
		name     string
		logLevel LogLevel
		level    int
		want     bool
	}{
		{DEBUG.String(), DEBUG, 0, true},
		{INFO.String(), INFO, 3, false},
		{WARNING.String(), WARNING, 3, false},
		{ERROR.String(), INFO, 1, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logLevel = LogLevel(tt.level)
			if checkLogLevel(tt.logLevel) != tt.want {
				t.Error("CheckLogLevel() returned wrong value")
			}
		})
	}
}

func TestGetColoredPrefix(t *testing.T) {
	SetColoredLogs(true)

	tests := []struct {
		name     string
		logLevel LogLevel
		want     string
	}{
		{"DEBUG", DEBUG, "\033[0;36m[DEBUG]   |\033[0m"},
		{"INFO", INFO, "\033[1;34m[INFO]    |\033[0m"},
		{"WARNING", WARNING, "\033[1;33m[WARNING] |\033[0m"},
		{"ERROR", ERROR, "\033[1;31m[ERROR]   |\033[0m"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getColoredPrefix(tt.logLevel); got != tt.want {
				t.Errorf("getColoredPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBuildJsonLog(t *testing.T) {
	tests := []struct {
		name       string
		logLevel   LogLevel
		params     Custom
		wantSuffix string
	}{
		{"DEBUG", DEBUG, Custom{"foo": "bar", "baz": 123}, `"baz":"123","foo":"bar","level":"DEBUG"}`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildJsonLog(tt.logLevel, []interface{}{tt.params}); !strings.HasSuffix(got, tt.wantSuffix) {
				t.Errorf("getColoredPrefix() = %v, want %v", got, tt.wantSuffix)
			}
		})
	}
}

func TestRemoveBrackets(t *testing.T) {
	type args struct {
		v []interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"single", args{v: []interface{}{"foo"}}, "foo"},
		{"multiple", args{v: []interface{}{"foo", "bar"}}, "foo bar"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeBrackets(tt.args.v); got != tt.want {
				t.Errorf("removeBrackets() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateString(t *testing.T) {
	type args struct {
		lvl    LogLevel
		format string
		_args  []interface{}
	}
	tests := []struct {
		name     string
		args     args
		wantText string
		wantJson string
	}{
		{"Debug", args{lvl: DEBUG, format: "%v %s", _args: []interface{}{200, "Test"}}, "\033[0;36m[DEBUG]   |\033[0m 200 Test", `level":"DEBUG","message":"200 Test"}` + "\n"},
	}
	for _, tt := range tests {
		SetLogType(TEXT.String())
		t.Run("[TEXT]"+tt.name, func(t *testing.T) {
			if got := CreateString(tt.args.lvl, tt.args.format, tt.args._args...); !strings.HasSuffix(got, tt.wantText) {
				t.Errorf("CreateString() = %v, want %v", got, tt.wantText)
			}
		})
		SetLogType(JSON.String())
		t.Run("[JSON]"+tt.name, func(t *testing.T) {
			if got := CreateString(tt.args.lvl, tt.args.format, tt.args._args...); !strings.HasSuffix(got, tt.wantJson) {
				t.Errorf("CreateString() = %v, want %v", got, tt.wantJson)
			}
		})
	}
}
