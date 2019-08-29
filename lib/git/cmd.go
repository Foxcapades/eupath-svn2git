package git

import (
	"bufio"
	"fmt"
	"gsman/lib/color"
	"gsman/lib/log"
	"io"
	"os/exec"
	"strings"
)

const (
	comGit        = "git"
	argSvn        = "svn"
	argFetch      = "fetch"
	argCo         = "checkout"
	argRb         = "rebase"
	argBranch     = "branch"
	argPush       = "push"
	argOrigin     = "origin"
	flagForce     = "-f"
	flagAuthors   = "--authors-file=../authors.txt"
)

const CHECKOUT_SAME_BRANCH string = "Already on '%s'"

type Processor struct {
	rebaseErrs []recErr
	pushErrs   []recErr
	fetchErrs  []recErr
	checkErrs  []recErr
	branchErrs []recErr
}


func (p *Processor) Checkout(project, branch string) (out bool) {
	log.TraceStart(project, branch)
	defer func() {log.TraceEnd(out)}()

	var build strings.Builder
	cmd := exec.Command(comGit, argCo, branch)
	cmd.Stderr = &build
	err := cmd.Run()

	if err != nil {
		errStr := build.String()
		if errStr != fmt.Sprintf(CHECKOUT_SAME_BRANCH, branch) {
			p.coError(project, branch, errStr, err)
			return false
		}
	}

	return true
}


func (p *Processor) CreateBranch(project, branch, refs string) (out bool) {
	log.TraceStart(project, branch, refs, out)
	defer func() {log.TraceEnd(out)}()

	cmd := exec.Command(comGit, argBranch, branch, refs)
	var stderr strings.Builder

	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		p.branchError(project, refs, stderr.String(), err)
		return false
	}

	return true
}


func (p *Processor) Rebase(project, ref string) (out bool) {
	log.TraceStart(project, ref)
	defer func() {log.TraceEnd(out)}()

	var build strings.Builder

	cmd := exec.Command(comGit, argRb, ref)
	cmd.Stderr = &build

	if err := cmd.Run(); err != nil {
		p.rebaseError(project, ref, build.String(), err)
		return false
	}

	return true
}

func (p *Processor) Fetch(project string) (out []string) {
	log.TraceStart(project)
	defer func() {log.TraceEnd(out)}()

	var err error
	var build strings.Builder

	defer func() {if err != nil {p.fetchError(project, build.String(), err)}}()

	var pipe io.ReadCloser

	cmd := exec.Command(comGit, argSvn, argFetch, flagAuthors)
	cmd.Stderr = &build
	pipe, err = cmd.StdoutPipe()
	if err != nil {
		return
	}

	err = cmd.Start()
	if err != nil {
		return
	}

	branches := make(map[string]int)

	scanner := bufio.NewScanner(pipe)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()

		if !isRev(line) {
			continue
		}

		branch := lineToBranch(line)

		if val, ok := branches[branch]; ok {
			branches[branch] = val + 1
		} else {
			branches[branch] = 1
		}
	}
	err = cmd.Wait()
	if err != nil {
		return
	}

	out = make([]string, len(branches))

	i := 0
	for k := range branches {
		out[i] = k
		i++
	}

	return
}


func (p *Processor) Push(proj, bran string) (out bool) {
	log.TraceStart(proj, bran)
	defer func() {log.TraceEnd(out)}()

	var stderr strings.Builder

	cmd := exec.Command(comGit, argPush, argOrigin, bran, flagForce)
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		p.pushError(proj, bran, stderr.String(), err)
		return false
	}

	return true
}


func (p Processor) WriteErrors() {
	log.TraceStart()
	defer log.TraceEnd()

	head := false

	errs := []struct{ name string; errs []recErr }{
		{"SVN Fetch Errors", p.fetchErrs},
		{"Git Branch Errors", p.branchErrs},
		{"Git Checkout Errors", p.checkErrs},
		{"Git Rebase Errors", p.rebaseErrs},
		{"Git Push Errors", p.pushErrs},
	}

	for i := range errs {
		if len(errs[i].errs) == 0 {
			continue
		}

		if !head {
			p.printHead()
			head = true
		}

		printErrs(errs[i].errs, "  " + errs[i].name + ":")
	}
}

func (p Processor) printHead() {
	log.TraceStart()
	defer log.TraceEnd()
	fmt.Println(color.DarkRed("WARNING"), "The following errors occurred during execution")
}

func (p *Processor) fetchError(pro, log string, err error) {
	p.fetchErrs = append(p.fetchErrs, newRecErr(pro, "N/A", log, err))
}

func (p *Processor) rebaseError(pro, brn, log string, err error) {
	p.rebaseErrs = append(p.rebaseErrs, newRecErr(pro, brn, log, err))
}

func (p *Processor) coError(pro, brn, log string, err error) {
	p.checkErrs = append(p.checkErrs, newRecErr(pro, brn, log, err))
}

func (p *Processor) psError(pro, brn, log string, err error) {
	p.pushErrs = append(p.pushErrs, newRecErr(pro, brn, log, err))
}

func (p *Processor) branchError(pro, brn, log string, err error) {
	p.branchErrs = append(p.branchErrs, newRecErr(pro, brn, log, err))
}

func (p *Processor) pushError(pro, brn, log string, err error) {
	p.pushErrs = append(p.pushErrs, newRecErr(pro, brn, log, err))
}