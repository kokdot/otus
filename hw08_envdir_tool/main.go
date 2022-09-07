package main

import (
	"log"
	"os"
)

func main() {
	cmd := os.Args[2:]
	env, err := ReadDir(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	_ = RunCmd(cmd, env)
}
