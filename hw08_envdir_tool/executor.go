package main

import (
	"fmt"
	"os"
	"os/exec"
)

// RunCmd runs a command + arguments (cmd) with environment variables from env.
func RunCmd(cmd []string, env Environment) (returnCode int) {
	// fmt.Println("env  --  :", env)
	// fmt.Println("cmd  --  :", cmd)
	cmdProg := exec.Command(cmd[0])
	cmdProg.Args = cmd
	// fmt.Println("cmdProg.Args  --  :", cmdProg.Args)
	for key, envElem := range env {
		if envElem.NeedRemove {
			os.Unsetenv(key)
		}
		// fmt.Println("key  --  :", key, "; envElem.Value  --  :", envElem.Value)
		// fmt.Printf("\n%v=%v\n", key, envElem.Value)
		cmdProg.Env = append(cmdProg.Environ(), fmt.Sprintf("%v=%v", key, envElem.Value))
	}
	// fmt.Println(cmdProg.Environ())
	cmdProg.Stdout = os.Stdout
	cmdProg.Stderr = os.Stderr
	cmdProg.Run()
	return 1
}
