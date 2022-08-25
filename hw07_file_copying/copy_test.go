package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	srcDir    = "./testdata/"
	dstDir    = "./testdata/testdir/"
	fileInput = "input.txt"
	fileTest1 = "out_offset0_limit0.txt"
	fileTest2 = "out_offset0_limit10.txt"
	fileTest3 = "out_offset0_limit1000.txt"
	fileTest4 = "out_offset0_limit10000.txt"
	fileTest5 = "out_offset100_limit1000.txt"
	fileTest6 = "out_offset6000_limit1000.txt"
)

func TestCopy(t *testing.T) {
	_ = os.Mkdir(dstDir, 0o755)
	defer os.RemoveAll(dstDir)
	fromPath := srcDir + fileInput

	t.Run("case offset0_limit0", func(t *testing.T) {
		t.Skip()
		toPath := dstDir + fileTest1
		toTest := srcDir + fileTest1
		_ = Copy(fromPath, toPath, 0, 0)
		result, _ := os.ReadFile(toPath)
		test, _ := os.ReadFile(toTest)
		require.Equal(t, result, test)
	})

	t.Run("case out_offset0_limit10", func(t *testing.T) {
		toPath := dstDir + fileTest2
		toTest := srcDir + fileTest2
		_ = Copy(fromPath, toPath, 0, 10)
		result, _ := os.ReadFile(toPath)
		test, _ := os.ReadFile(toTest)
		require.Equal(t, result, test)
	})

	t.Run("case out_offset0_limit1000", func(t *testing.T) {
		toPath := dstDir + fileTest3
		toTest := srcDir + fileTest3
		_ = Copy(fromPath, toPath, 0, 1000)
		result, _ := os.ReadFile(toPath)
		test, _ := os.ReadFile(toTest)
		require.Equal(t, result, test)
	})

	t.Run("case out_offset0_limit100000", func(t *testing.T) {
		toPath := dstDir + fileTest4
		toTest := srcDir + fileTest4
		_ = Copy(fromPath, toPath, 0, 10000)
		result, _ := os.ReadFile(toPath)
		test, _ := os.ReadFile(toTest)
		require.Equal(t, result, test)
	})

	t.Run("case out_offset100_limit1000", func(t *testing.T) {
		toPath := dstDir + fileTest5
		toTest := srcDir + fileTest5
		_ = Copy(fromPath, toPath, 100, 1000)
		result, _ := os.ReadFile(toPath)
		test, _ := os.ReadFile(toTest)
		require.Equal(t, result, test)
	})

	t.Run("case out_offset6000_limit1000", func(t *testing.T) {
		toPath := dstDir + fileTest6
		toTest := srcDir + fileTest6
		_ = Copy(fromPath, toPath, 6000, 1000)
		result, _ := os.ReadFile(toPath)
		test, _ := os.ReadFile(toTest)
		require.Equal(t, result, test)
	})
}
