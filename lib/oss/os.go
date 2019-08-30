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

