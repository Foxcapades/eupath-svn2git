package main

import (
	"github.com/Foxcapades/eupath-svn2git/lib/git"
	"github.com/Foxcapades/eupath-svn2git/lib/log"
	"github.com/Foxcapades/eupath-svn2git/lib/oss"
	"github.com/Foxcapades/lib-go-ansi-esc/color"
	tl "github.com/Foxcapades/lib-go-tree-log"
	"github.com/Foxcapades/lib-go-xos"
	"path"
)

// User facing messages
const (
	msgFetch     = "Fetching changes from SVN"
	msgNoChanges = "No new changes"
	msgBranches  = "Processing branches:"
	msgCreating  = "Creating"
	msgCd        = "Project: "
	msgCo        = "Checking out"
	msgRebase    = "Rebasing onto SVN remote branch"
	msgPush      = "Pushing changes to github"
	msgNoGit     = "Not a git repo"

	errBranch = "Failed to create branch"
	errCo     = "Failed to checkout branch"
	errRebase = "Failed to rebase branch"
	errPush   = "Git push failed"
)

// Paths
const (
	dirUp   = ".."
	dirGit  = ".git"
	refHead = ".git/refs/heads"
)

var tree = tl.DefaultLogger()

func main() {

	proc := git.Processor{}

	for _, dir := range oss.Dirs() {
		tree.WriteLn(msgCd, color.FgLightCyanText(dir))
		tree.Indent()
		handleProject(&proc, dir)
		tree.UnIndent()
	}

	proc.WriteErrors()
}

func handleProject(git *git.Processor, project string) {
	xos.Chdir(project)
	defer xos.Chdir(dirUp)

	if !xos.DirExists(dirGit) {
		pInfo(msgNoGit)
		return
	}

	pInfo(msgFetch)
	branches := git.Fetch(project)

	if len(branches) == 0 {
		pInfo(msgNoChanges)
		return
	}

	pInfo(msgBranches)

	for _, branch := range branches {
		tree.Indent()
		handleBranch(git, project, branch)
		tree.UnIndent()
	}
}

func handleBranch(git *git.Processor, project, branch string) {
	local := toLocalBranchName(branch)

	tree.WriteLn(color.FgDarkGreenText(local)).Indent()
	defer tree.UnIndent()

	if !xos.FileExists(path.Join(refHead, local)) {
		tree.WriteLn(color.FgLightYellowText(msgCreating))
		if !git.CreateBranch(project, local, branch) {
			pErr(errBranch)
			return
		}
	}

	pInfo(msgCo)
	if !git.Checkout(project, local) {
		pErr(errCo)
		return
	}

	pInfo(msgRebase)
	if !git.Rebase(project, branch) {
		pErr(errRebase)
		return
	}

	pInfo(msgPush)
	if !git.Push(project, branch) {
		pErr(errPush)
	}
}

func toLocalBranchName(branch string) (o string) {
	log.TraceStart(branch)
	defer func() { log.TraceEnd(o) }()
	_, o = path.Split(branch)
	return
}

func pErr(txt string) {
	tree.WriteLn(color.FgDarkRedText(txt))
}

func pInfo(txt string) {
	tree.WriteLn(color.FgDarkGrayText(txt))
}
