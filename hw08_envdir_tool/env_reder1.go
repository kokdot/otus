// package main

// import (
// 	// "bytes"
// 	// "fmt"
// 	"log"
// 	"os"
// 	"regexp"
// 	"strings"
// )
// type Environment map[string]EnvValue

// // EnvValue helps to distinguish between empty files and files with the first empty line.
// type EnvValue struct {
// 	Value      string
// 	NeedRemove bool
// }

// // ReadDir reads a specified directory and returns map of env variables.
// // Variables represented as files where filename is name of variable, file first line is a value.
// func ReadDir(dir string) (Environment, error) {
// 	re := regexp.MustCompile(`=`)
// 	reNull := regexp.MustCompile(`\x00`)
// 	reFistString := regexp.MustCompile(`.*\n`)
// 	mapFileContent := make(map[string][]byte)
// 	env := make(Environment)
// 	fileList, err := os.ReadDir(dir)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	for _, file := range fileList {
// 		name := file.Name()
// 		if re.Match([]byte(name)) {
// 			continue
// 		}
// 		fileContent, err := os.ReadFile(dir + name)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		mapFileContent[name] = fileContent
// 	}
// 	for nameFile, content := range mapFileContent {
// 		if len(content) == 0 {
// 			env[nameFile] = EnvValue{Value: "", NeedRemove: true}
// 			continue
// 		}
// 		// if reNull.Match(content) {
// 		// 	content = bytes.Replace(content, []byte("\x00"), []byte("\n"), -1)
// 		// }
// 		// fmt.Println("content --  :", content)
// 		contentStr := string(content)
// 		// fmt.Println("contentStr --  :", contentStr)
// 		if !reFistString.MatchString(contentStr) {
// 			if reNull.MatchString(contentStr) {
// 				contentStr = strings.Replace(contentStr, "\x00", "\n", -1)
// 			}
// 			env[nameFile] = EnvValue{Value: contentStr, NeedRemove: false}
// 			continue
// 		}
// 		contentStr = reFistString.FindString(contentStr)
// 		if len(contentStr) == 1{
// 			env[nameFile] = EnvValue{Value: "", NeedRemove: false}
// 			continue
// 		}
// 		contentStr = strings.TrimSuffix(contentStr, "\n")
// 		contentStr = strings.TrimRight(string(contentStr), " \t")
// 		if reNull.MatchString(contentStr) {
// 			contentStr = strings.Replace(contentStr, "\x00", "\n", -1)
// 		}
// 		env[nameFile] = EnvValue{Value: contentStr, NeedRemove: false}

// 	}

// 	return env, nil
// }
