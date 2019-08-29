package git

import (
	"gsman/lib/log"
	"strings"
)

const (
	HEAD_PROJECT = "    Project: "
	HEAD_BRANCH  = "\n      Branch: "
	HEAD_ERROR   = "\n      Error: "
	HEAD_STDERR  = "\n      Stderr:\n"
	TAIL         = "\n"
)

const (
	LEN_MOD_PROJECT = len(HEAD_PROJECT)
	LEN_MOD_BRANCH  = len(HEAD_BRANCH)
	LEN_MOD_ERROR   = len(HEAD_ERROR)
	LEN_MOD_STDERR  = len(HEAD_STDERR)
	LEN_MOD_TAIL    = len(TAIL)
)

type recErr struct {
	project string
	branch  string
	error   string
	err     error
}

func newRecErr(project, branch, error string, err error) (o recErr) {
	log.TraceStart(project, branch, error, err)
	defer func() {log.TraceEnd(o)}()
	return recErr{
		project: project,
		branch:  branch,
		error:   error,
		err:     err,
	}
}

func (r recErr) len() (o int) {
	log.TraceStart()
	defer func() {log.TraceEnd(o)}()
	return len(r.project) + LEN_MOD_PROJECT +
		len(r.branch) + LEN_MOD_BRANCH +
		len(r.err.Error()) + LEN_MOD_ERROR +
		len(r.error) + LEN_MOD_STDERR +
		LEN_MOD_TAIL
}

func (r recErr) String() string {
	var out strings.Builder
	out.Grow(r.len())
	out.WriteString(HEAD_PROJECT)
	out.WriteString(r.project)
	out.WriteString(HEAD_BRANCH)
	out.WriteString(r.branch)
	out.WriteString(HEAD_ERROR)
	out.WriteString(r.err.Error())
	out.WriteString(HEAD_STDERR)
	out.WriteString(r.error)
	out.WriteString(TAIL)
	return out.String()
}
