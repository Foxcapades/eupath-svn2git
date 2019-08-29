package git

import (
	"fmt"
	"gsman/lib/log"
	"strings"
)

func isRev(in string) (o bool) {
	log.TraceStart(in)
	defer func() {log.TraceEnd(o)}()
	return in[0] == 'r' && in[1] >= '0' && in[1] <= '9'
}

func lineToBranch(in string) (o string) {
	log.TraceStart(in)
	defer func() {log.TraceEnd(o)}()
	return in[strings.IndexByte(in, '(')+1:len(in)-1]
}

func printErrs(errs []recErr, head string) {
	fmt.Println(head)
	for _, i := range errs {
		fmt.Println(i.String())
		fmt.Println(strings.Repeat("-", 60))
		fmt.Println()
	}
}
