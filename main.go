package main

import (
	"gsman/lib/color"
	"gsman/lib/git"
	"gsman/lib/log"
	"gsman/lib/oss"
	"path"
)

const GIT_HEAD_PATH string = ".git/refs/heads"

func main() {

	var proc git.Processor

	for _, dir := range oss.Dirs() {
		log.WriteLn("Entering path", color.Cyan(dir))
		oss.Cd(dir)

		log.WriteLn(color.Gray("  Fetching changes from SVN"))
		branches := proc.Fetch(dir)

		if len(branches) == 0 {
			log.WriteLn(color.Gray("  No new changes"))
			oss.Cd("..")
			continue
		}

		log.WriteLn("  Processing branches:")

		for _, branch := range branches {

			local := toLocalBranchName(branch)

			log.WriteLn("    " + local)

			if !oss.FileExists(path.Join(GIT_HEAD_PATH, local)) {
				log.WriteLn(color.Blue("      Creating"))
				if !proc.CreateBranch(dir, local, branch) {
					log.WriteLn(color.Red("      FAILED TO CREATE BRANCH"))
					continue
				}
			}

			log.WriteLn("      Checking out")
			if !proc.Checkout(dir, local) {
				log.WriteLn(color.Red("      FAILED TO CHECKOUT BRANCH"))
				continue
			}

			log.WriteLn("      Rebasing onto SVN remote")
			if !proc.Rebase(dir, branch) {
				log.WriteLn(color.Red("      Rebasing Failed!"))
				continue
			}

			log.WriteLn("      Pushing up new changes")
			if !proc.Push(dir, branch) {
				log.WriteLn(color.Red("      Push failed!"))
			}
		}

		oss.Cd("..")
	}

	proc.WriteErrors()
}

func toLocalBranchName(branch string) (o string) {
	log.TraceStart(branch)
	defer func() {log.TraceEnd(o)}()
	_, o = path.Split(branch)
	return
}