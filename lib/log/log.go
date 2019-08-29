package log

import (
	"fmt"
	"runtime"
	"strings"
)

const (
	showTrace = false
)

var stack uint8 = 0

func TraceStart(params ...interface{}) {
	stack++
	//noinspection GoBoolExpressions
	if showTrace {
		WriteLn(traceName(), "<-", params)
	}
}

func TraceEnd(out ...interface{}) {
	//noinspection GoBoolExpressions
	if showTrace {
		WriteLn(traceName(), "->", out)
	}
	stack--
}

func WriteLn(p ...interface{}) {
	fmt.Print(indent())
	fmt.Println(p...)
}

var indents = []string {
	"",
	"  ",
	"    ",
	"      ",
	"        ",
}
func indent() string {
	ln := len(indents)
	st := int(stack)
	if ln > st {
		return indents[st]
	}

	n := st - ln
	for i := 1; i <= n; i++ {
		indents = append(indents, strings.Repeat("  ", i + ln))
	}

	return indents[st]
}

func traceName() (name string) {
	pc := make([]uintptr, 2)
	runtime.Callers(3, pc)
	name = runtime.FuncForPC(pc[0]).Name()
	if strings.HasSuffix(name[:len(name)-1], ".func") {
		name = runtime.FuncForPC(pc[1]).Name()
	}
	return
}
