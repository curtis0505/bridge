package logger

import (
	"fmt"
	"io"
	"os"
	"runtime"
)

func TraceCallStack(skip int) string {
	pc := make([]uintptr, 15)
	n := runtime.Callers(skip, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	return fmt.Sprintf("%s:%d", frame.File, frame.Line)
}

func ErrorWithTrace(err error) error {
	return fmt.Errorf("%s:%v", TraceCallStack(3), err)
}

func Fatalf(format string, args ...interface{}) {
	w := io.MultiWriter(os.Stdout, os.Stderr)
	if runtime.GOOS == "windows" {
		// The SameFile check below doesn't work on Windows.
		// stdout is unlikely to get redirected though, so just print there.
		w = os.Stdout
	} else {
		outf, _ := os.Stdout.Stat()
		errf, _ := os.Stderr.Stat()
		if outf != nil && errf != nil && os.SameFile(outf, errf) {
			w = os.Stderr
		}
	}
	fmt.Fprintf(w, "Fatal: "+format+"\n", args...)
	os.Exit(1)
}

func createLogs(ctx ...interface{}) []interface{} {
	logs := make([]interface{}, 0)

	for _, log := range ctx {
		logs = append(logs, log)
	}

	return logs
}
