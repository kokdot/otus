package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

// RunCmd runs a command + arguments (cmd) with environment variables from env.
func RunCmd(cmd []string, env Environment) (returnCode int) {
	cmdProg := exec.Command(cmd[1])
	cmdProg.Args = cmd[1:]
	for key, envElem := range env {
		if envElem.NeedRemove {
			os.Unsetenv(key)
		}
		cmdProg.Env = append(cmdProg.Env, fmt.Sprintf("%v=%v", key, envElem.Value))
	}
	cmdProg.Env = append(os.Environ(), cmdProg.Env...)
	cmdProg.Stdout = os.Stdout
	cmdProg.Stderr = os.Stderr
	if err := cmdProg.Run(); err != nil {
		log.Fatal(err)
	}
	returnCode = 1
	return returnCode
}
