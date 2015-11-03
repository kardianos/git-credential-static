package main

import (
	"fmt"
	"io"
	"strings"
)

type KV map[string]string

func read(r io.Reader) (KV, error) {
	lookup := make(KV, 3)
	for {
		var line string
		_, err := fmt.Fscanf(r, "%s\n", &line)
		if err == io.EOF {
			return lookup, nil
		}
		if err != nil {
			return lookup, err
		}
		kv := strings.Split(line, "=")
		if len(kv) != 2 {
			return lookup, fmt.Errorf("Not in format of key=value: %q", line)
		}
		lookup[kv[0]] = kv[1]
	}
	return lookup, nil
}
