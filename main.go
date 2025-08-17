package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Wa4h1h/brainfuck-interpter/repl"
)

func main() {
	var bfSrc string

	flag.StringVar(&bfSrc, "bf-src", "", "bf file path")

	flag.Parse()

	if len(bfSrc) != 0 {
		content, err := os.ReadFile(bfSrc)
		if err != nil {
			fmt.Fscanln(os.Stderr, err)

			return
		}

		fmt.Fprintln(os.Stdout, string(content))

		return
	}

	err := repl.Start(os.Stdin, os.Stdout)
	if err != nil {
		fmt.Fscanln(os.Stderr, err)
	}
}
