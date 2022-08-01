package utils

import (
    "fmt"
    "strings"

    log "github.com/sirupsen/logrus"
)

type MyFormatter struct {
    log.TextFormatter
}

func (f *MyFormatter) Format(entry *log.Entry) ([]byte, error) {
// this whole mess of dealing with ansi color codes is required if you want the colored output otherwise you will lose colors in the log levels
    var levelColor int
    switch entry.Level {
    case log.DebugLevel, log.TraceLevel:
        levelColor = 35 // purple
    case log.WarnLevel:
        levelColor = 33 // yellow
    case log.ErrorLevel, log.FatalLevel, log.PanicLevel:
        levelColor = 31 // red
    default:
        levelColor = 36 // blue
    }
    return []byte(fmt.Sprintf("[%s] - \x1b[%dm%s\x1b[0m - %s\n", entry.Time.Format(f.TimestampFormat), levelColor, strings.ToUpper(entry.Level.String()), entry.Message)), nil
}
