package main

import (
	"strings"
	"testing"
)

var input = `protocol=https
host=files.sembit.com
username=danielt
`

func TestRead(t *testing.T) {
	kv, err := read(strings.NewReader(input))
	if err != nil {
		t.Fatal(err)
	}
	_ = kv
}
