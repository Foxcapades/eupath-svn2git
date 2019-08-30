package main

import (
	"github.com/Foxcapades/eupath-svn2git/lib/git"
	"github.com/Foxcapades/eupath-svn2git/lib/log"
	"github.com/Foxcapades/eupath-svn2git/lib/oss"
	"github.com/Foxcapades/lib-go-ansi-esc/color"
	"github.com/Foxcapades/lib-go-xos"
	"path"
	"strings"
)

// User facing messages
const (
	msgFetch     = "Fetching changes from SVN"
	msgNoChanges = "No new changes"
	msgBranches  = "Processing branches:"
	msgCreating  = "Creating"
	msgCd        = "Entering path"
	msgCo        = "Checking out"
	msgRebase    = "Rebasing onto SVN remote branch"
	msgPush      = "Pushing changes to github"

	errBranch = "Failed to create branch"
	errCo     = "Failed to checkout branch"
	errRebase = "Failed to rebase branch"
	errPush   = "Git push failed"
)

// Paths
const (
	dirUp   = ".."
	refHead = ".git/refs/heads")

func main() {

	var proc git.Processor

	for _, dir := range oss.Dirs() {
		log.WriteLn(msgCd, color.FgLightCyanText(dir))
		xos.Chdir(dir)

		pInfo(indent(msgFetch, 1))
		branches := proc.Fetch(dir)

		if len(branches) == 0 {
			pInfo(indent(msgNoChanges, 1))
			xos.Chdir(dirUp)
			continue
		}

		pInfo(indent(msgBranches, 1))

		for _, branch := range branches {

			local := toLocalBranchName(branch)

			log.WriteLn(color.FgDarkGreenText(indent(local, 2)))

			if !xos.FileExists(path.Join(refHead, local)) {
				log.WriteLn(color.FgLightYellowText(indent(msgCreating, 2)))
				if !proc.CreateBranch(dir, local, branch) {
					pErr(indent(errBranch, 3))
					continue
				}
			}

			pInfo(indent(msgCo, 3))
			if !proc.Checkout(dir, local) {
				pErr(indent(errCo, 3))
				continue
			}

			pInfo(indent(msgRebase, 3))
			if !proc.Rebase(dir, branch) {
				pErr(indent(errRebase, 3))
				continue
			}

			pInfo(indent(msgPush, 3))
			if !proc.Push(dir, branch) {
				pErr(indent(errPush, 3))
			}
		}

		xos.Chdir(dirUp)
	}

	proc.WriteErrors()
}

func toLocalBranchName(branch string) (o string) {
	log.TraceStart(branch)
	defer func() { log.TraceEnd(o) }()
	_, o = path.Split(branch)
	return
}

func indent(txt string, n int) string {
	return strings.Repeat("  ", n) + txt
}

func pErr(txt string) {
	log.WriteLn(color.FgDarkRedText(txt))
}
func pInfo(txt string) {
	log.WriteLn(color.FgDarkGrayText(txt))
}
