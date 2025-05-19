package util

import (
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

func HomeDir() string {
	if home := os.Getenv("HOME"); home != "" {
		return home
	}
	if usr, err := user.Current(); err == nil {
		return usr.HomeDir
	}
	return ""
}

func WorkingDir() string {
	workdir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	} else {
		if strings.Contains(workdir, "/var/folders/") == true {
			workdir = "./"
		}
	}
	return workdir
}
