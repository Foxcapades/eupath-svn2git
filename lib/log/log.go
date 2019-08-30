package log

import (
	tl "github.com/Foxcapades/lib-go-tree-log"
	"runtime"
	"strings"
)

const (
	showTrace = false
)

var tree = tl.DefaultLogger()

func TraceStart(params ...interface{}) {
	//noinspection GoBoolExpressions
	if showTrace {
		tree.Indent()
		tree.WriteLn(traceName(), "<-", params)
	}
}

func TraceEnd(out ...interface{}) {
	//noinspection GoBoolExpressions
	if showTrace {
		tree.WriteLn(traceName(), "->", out)
		tree.UnIndent()
	}
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
