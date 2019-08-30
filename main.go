package main

import (
	"github.com/Foxcapades/eupath-svn2git/lib/git"
	"github.com/Foxcapades/eupath-svn2git/lib/log"
	"github.com/Foxcapades/eupath-svn2git/lib/oss"
	"github.com/Foxcapades/lib-go-ansi-esc/color"
	"github.com/Foxcapades/lib-go-xos"
	"path"
)

const GIT_HEAD_PATH string = ".git/refs/heads"

func main() {

	var proc git.Processor

	for _, dir := range oss.Dirs() {
		log.WriteLn("Entering path", color.FgLightCyanText(dir))
		xos.Chdir(dir)

		log.WriteLn(color.FgDarkGrayText("  Fetching changes from SVN"))
		branches := proc.Fetch(dir)

		if len(branches) == 0 {
			log.WriteLn(color.FgDarkGrayText("  No new changes"))
			xos.Chdir("..")
			continue
		}

		log.WriteLn(color.FgDarkGrayText("  Processing branches:"))

		for _, branch := range branches {

			local := toLocalBranchName(branch)

			log.WriteLn(color.FgDarkGreenText("    " + local))

			if !xos.FileExists(path.Join(GIT_HEAD_PATH, local)) {
				log.WriteLn(color.FgLightYellowText("      Creating"))
				if !proc.CreateBranch(dir, local, branch) {
					log.WriteLn(color.FgDarkRedText("      FAILED TO CREATE BRANCH"))
					continue
				}
			}

			log.WriteLn(color.FgDarkGrayText("      Checking out"))
			if !proc.Checkout(dir, local) {
				log.WriteLn(color.FgDarkRedText("      FAILED TO CHECKOUT BRANCH"))
				continue
			}

			log.WriteLn(color.FgDarkGrayText("      Rebasing onto SVN remote"))
			if !proc.Rebase(dir, branch) {
				log.WriteLn(color.FgDarkRedText("      Rebasing Failed!"))
				continue
			}

			log.WriteLn(color.FgDarkGrayText("      Pushing up new changes"))
			if !proc.Push(dir, branch) {
				log.WriteLn(color.FgDarkRedText("      Push failed!"))
			}
		}

		xos.Chdir("..")
	}

	proc.WriteErrors()
}

func toLocalBranchName(branch string) (o string) {
	log.TraceStart(branch)
	defer func() {log.TraceEnd(o)}()
	_, o = path.Split(branch)
	return
}