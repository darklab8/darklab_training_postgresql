package main

import (
	"flag"
	"fmt"
)

func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}

func main() {
	forkPtr := flag.Bool("fork", false, "a bool")
	flag.Parse()
	fmt.Println(
		"fork:", *forkPtr,
		"passed:", isFlagPassed("fork"),
	)
}
