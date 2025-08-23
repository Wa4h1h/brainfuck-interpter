package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"

	"github.com/Wa4h1h/brainfuck-interpter/evaluator"
	"github.com/Wa4h1h/brainfuck-interpter/repl"
)

const memorySize = 30000

func main() {
	var (
		useRepl bool
		exec    string
		memSize int
	)

	flag.BoolVar(&useRepl, "repl", false, "start repl")
	flag.StringVar(&exec, "exec", "", "bf file to execute")
	flag.IntVar(&memSize, "mem-size", memorySize, "memory size")

	flag.Parse()

	if useRepl {
		err := repl.Start(os.Stdin, os.Stdout)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		return
	}

	e := evaluator.New(memSize)

	if len(exec) != 0 {
		content, err := os.ReadFile(exec)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)

			return
		}

		err = e.Evaluate(bytes.NewReader(content))
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}
