package main

import (
	"log"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadDir(t *testing.T) {
	envExpected := Environment{
		"BAR":   EnvValue{Value: "bar", NeedRemove: false},
		"FOO":   EnvValue{Value: "   foo", NeedRemove: false},
		"UNSET": EnvValue{Value: "", NeedRemove: true},
		"EMPTY": EnvValue{Value: "", NeedRemove: false},
		"HELLO": EnvValue{Value: "\"hello\"", NeedRemove: false},
	}

	envActual, err := ReadDir("testdata/env/")
	if err != nil {
		log.Fatal(err)
	}
	require.Equal(t, envExpected, envActual)
}
