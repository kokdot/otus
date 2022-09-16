package main

import (
	"bytes"
	"log"
	"os"
	"regexp"
	"strings"
)

type Environment map[string]EnvValue

// EnvValue helps to distinguish between empty files and files with the first empty line.
type EnvValue struct {
	Value      string
	NeedRemove bool
}

// ReadDir reads a specified directory and returns map of env variables.
// Variables represented as files where filename is name of variable, file first line is a value.
func ReadDir(dir string) (Environment, error) {
	re := regexp.MustCompile(`=`)
	reFistString := regexp.MustCompile(`.*\n`)
	reNull := regexp.MustCompile(`\x00`)
	mapFileContent := make(map[string][]byte)
	env := make(Environment)
	fileList, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	for _, file := range fileList {
		name := file.Name()
		if re.Match([]byte(name)) {
			continue
		}
		fileContent, err := os.ReadFile(dir + name)
		if err != nil {
			log.Fatal(err)
		}
		mapFileContent[name] = fileContent
	}
	for nameFile, content := range mapFileContent {
		if len(content) == 0 {
			env[nameFile] = EnvValue{Value: "", NeedRemove: true}
			continue
		}
		if !reFistString.Match(content) {
			if reNull.Match(content) {
				content = bytes.ReplaceAll(content, []byte("\x00"), []byte("\n"))
			}
			contentStr := string(content)
			contentStr = strings.TrimRight(contentStr, " \t")
			env[nameFile] = EnvValue{Value: contentStr, NeedRemove: false}
			continue
		} else {
			content = reFistString.Find(content)
			if len(string(content)) == 1 {
				env[nameFile] = EnvValue{Value: "", NeedRemove: false}
				continue
			}
			content = content[:len(content)-1]
			if reNull.Match(content) {
				content = bytes.ReplaceAll(content, []byte("\x00"), []byte("\n"))
			}
			contentStr := string(content)
			contentStr = strings.TrimRight(contentStr, " \t")
			env[nameFile] = EnvValue{Value: contentStr, NeedRemove: false}
			continue
		}
	}
	return env, nil
}
