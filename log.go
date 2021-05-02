package ui

import "fmt"

func Log(a ...interface{}) {
	logView, err := gui.View("log")
	if err == nil {
		fmt.Fprint(logView, a...)
	}
}
func Logf(format string, a ...interface{}) {
	logView, err := gui.View("log")
	if err == nil {
		fmt.Fprintf(logView, format, a...)
	}
}
func Logfln(format string, a ...interface{}) {
	logView, err := gui.View("log")
	if err == nil {
		fmt.Fprintf(logView, format, a...)
		fmt.Fprintln(logView)
	}
}
func Logln(a ...interface{}) {
	logView, err := gui.View("log")
	if err == nil {
		fmt.Fprintln(logView, a...)
	}
}
