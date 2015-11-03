package main

import "os"

func main() {
	if len(os.Args) < 2 {
		printUsageExit("not enough args")
	}
	var err error
	verb := os.Args[1]
	switch verb {
	default:
		printUsageExit("Unknown verb: " + verb)
	case "get":
		err = get(os.Stdin, os.Stdout)
	case "store":
		err = store(os.Stdin)
	case "erase":
		// noop
	}
	if err != nil {
		printUsageExit(verb + ": " + err.Error())
	}
}

func printUsageExit(msg string) {
	if len(msg) > 0 {
		os.Stderr.Write([]byte(msg))
		os.Stderr.Write([]byte("\n"))
	}
	os.Stderr.Write([]byte(usage))
	os.Exit(1)
}

var usage = `<exec> {get | store}
`
