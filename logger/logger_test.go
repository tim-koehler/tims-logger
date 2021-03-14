package logger

import (
	"strings"
	"testing"
)

func TestSetLogLevel(t *testing.T) {
	tests := []struct {
		name     string
		logLevel LogLevel
		level    int
	}{
		{DEBUG.String(), DEBUG, 0},
		{INFO.String(), INFO, 1},
		{WARNING.String(), WARNING, 2},
		{ERROR.String(), ERROR, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logLevel = 0
			if SetLogLevel(tt.logLevel.String()); int(logLevel) != tt.level {
				t.Error("SetLogLevel() failed to set correct log level")
			}
		})
	}
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
	type args struct {
		lvl LogLevel
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"DEBUG", args{lvl: DEBUG}, "\033[0;36m[DEBUG]   |\033[0m"},
		{"INFO", args{lvl: INFO}, "\033[1;34m[INFO]    |\033[0m"},
		{"WARNING", args{lvl: WARNING}, "\033[1;33m[WARNING] |\033[0m"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getColoredPrefix(tt.args.lvl); got != tt.want {
				t.Errorf("getColoredPrefix() = %v, want %v", got, tt.want)
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
		name string
		args args
		want string
	}{
		{
			"Debug",
			args{
				lvl:    DEBUG,
				format: "%v %s",
				_args:  []interface{}{200, "Test"},
			},
			"\033[0;36m[DEBUG]   |\033[0m 200 Test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateString(tt.args.lvl, tt.args.format, tt.args._args...); !strings.HasSuffix(got, tt.want) {
				t.Errorf("CreateString() = %v, want %v", got, tt.want)
			}
		})
	}
}
