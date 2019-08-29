package oss

import (
	"gsman/lib/log"
	"io/ioutil"
	"os"
)

func Cwd() (o string) {
	log.TraceStart()
	defer func() {log.TraceEnd(o)}()
	a, b := os.Getwd()
	if b != nil {
		panic(b)
	}
	return a
}

func DirExists(path string) (out bool) {
	log.TraceStart(path)
	defer func() {log.TraceEnd(out)}()
	a, b := os.Stat(path)

	if b != nil {
		if os.IsNotExist(b) {
			return false
		}

		panic(b)
	}

	return a.IsDir()
}

func FileExists(path string) (out bool) {
	log.TraceStart(path)
	defer func() {log.TraceEnd(out)}()

	stat, err := os.Stat(path)

	if err != nil {
		if os.IsNotExist(err) {
			return false
		}

		panic(err)
	}

	return !stat.IsDir()
}

func Dirs() (dirs []string) {
	log.TraceStart()
	defer func() {log.TraceEnd(dirs)}()
	ls, err := ioutil.ReadDir(Cwd())
	if err != nil {
		panic(err)
	}
	dirs = make([]string, 0)
	for _, i := range ls {
		if i.IsDir() {
			dirs = append(dirs, i.Name())
		}
	}
	return
}

func Cd(path string) {
	log.TraceStart()
	err := os.Chdir(path)
	if err != nil {
		panic(err)
	}
	log.TraceEnd()
}
