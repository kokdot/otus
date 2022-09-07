package main

import (
	"log"
	"os"
	"testing"

	// "fmt".
	"github.com/stretchr/testify/require"
)

const (
	srcDir    = "./testdata/"
	fileInput = "input.txt"
	fileTest1 = "out_offset0_limit0.txt"
	fileTest2 = "out_offset0_limit10.txt"
	fileTest3 = "out_offset0_limit1000.txt"
	fileTest4 = "out_offset0_limit10000.txt"
	fileTest5 = "out_offset100_limit1000.txt"
	fileTest6 = "out_offset6000_limit1000.txt"
)

func TestCopy(t *testing.T) {
	dstDir, _ := os.MkdirTemp(
		"/mnt/c/Users/user/Documents/go/otus/otus/hw07_file_copying/testdata/",
		"testdir",
	)
	defer os.RemoveAll(dstDir)
	fromPath := srcDir + fileInput

	t.Run("case offset0_limit0", func(t *testing.T) {
		// t.Skip()
		toPath, err := os.CreateTemp(dstDir, "out")
		if err != nil {
			log.Fatal(err)
		}
		toTest := srcDir + fileTest1
		_ = Copy(fromPath, toPath.Name(), 0, 0)
		result, _ := os.ReadFile(toPath.Name())
		test, _ := os.ReadFile(toTest)
		require.Equal(t, test, result)
	})

	t.Run("case offset0_limit10", func(t *testing.T) {
		// t.Skip()
		toPath, err := os.CreateTemp(dstDir, "out")
		if err != nil {
			log.Fatal(err)
		}
		toTest := srcDir + fileTest2
		_ = Copy(fromPath, toPath.Name(), 0, 10)
		result, _ := os.ReadFile(toPath.Name())
		test, _ := os.ReadFile(toTest)
		require.Equal(t, test, result)
	})

	t.Run("case offset0_limit1000", func(t *testing.T) {
		// t.Skip()
		toPath, err := os.CreateTemp(dstDir, "out")
		if err != nil {
			log.Fatal(err)
		}
		toTest := srcDir + fileTest3
		_ = Copy(fromPath, toPath.Name(), 0, 1000)
		result, _ := os.ReadFile(toPath.Name())
		test, _ := os.ReadFile(toTest)
		require.Equal(t, test, result)
	})

	t.Run("case offset0_limit10000", func(t *testing.T) {
		// t.Skip()
		toPath, err := os.CreateTemp(dstDir, "out")
		if err != nil {
			log.Fatal(err)
		}
		toTest := srcDir + fileTest4
		_ = Copy(fromPath, toPath.Name(), 0, 10000)
		result, _ := os.ReadFile(toPath.Name())
		test, _ := os.ReadFile(toTest)
		require.Equal(t, test, result)
	})

	t.Run("case offset100_limit1000", func(t *testing.T) {
		// t.Skip()
		toPath, err := os.CreateTemp(dstDir, "out")
		if err != nil {
			log.Fatal(err)
		}
		toTest := srcDir + fileTest5
		_ = Copy(fromPath, toPath.Name(), 100, 1000)
		result, _ := os.ReadFile(toPath.Name())
		test, _ := os.ReadFile(toTest)
		require.Equal(t, test, result)
	})

	t.Run("case offset6000_limit1000", func(t *testing.T) {
		// t.Skip()
		toPath, err := os.CreateTemp(dstDir, "out")
		if err != nil {
			log.Fatal(err)
		}
		toTest := srcDir + fileTest6
		_ = Copy(fromPath, toPath.Name(), 6000, 1000)
		result, _ := os.ReadFile(toPath.Name())
		test, _ := os.ReadFile(toTest)
		require.Equal(t, test, result)
	})
}
