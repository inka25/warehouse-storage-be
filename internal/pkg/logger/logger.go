package logger

import (
	"fmt"
	"time"
)

type LogWriter struct {
	AppName string
	Loc     *time.Location
	Env     string
}

func (l *LogWriter) Write(bytes []byte) (int, error) {
	return fmt.Printf("%s [%s-%s] %s", time.Now().In(l.Loc).Format("2006-01-02 15:04:05"), l.AppName, l.Env, string(bytes))
}
