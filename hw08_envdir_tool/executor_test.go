package main

import (
	"log"
	"testing"
)

func TestRunCmd(t *testing.T) {
	env, err := ReadDir("testdata/env/")
	if err != nil {
		log.Fatal(err)
	}
	// _ = RunCmd([]string{"ls"}, env)
	_ = RunCmd([]string{"/mnt/c/Users/user/Documents/go/otus/otus/hw08_envdir_tool//testdata/env",
		"/bin/bash", "/mnt/c/Users/user/Documents/go/otus/otus/hw08_envdir_tool//testdata/echo.sh", "arg1=1", "arg2=2"}, env)
	// _ = RunCmd([]string{"./testdata/echo.sh", "arg1=1", "arg2=2"}, env)
	// return result
}
