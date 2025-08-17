package repl

import (
	"bufio"
	"fmt"
	"io"
)

const prompt = "bfi>>"

func Start(in io.Reader, out io.Writer) error {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprintf(out, "%s ", prompt)

		if !scanner.Scan() {

			if scanner.Err() != nil {
				return fmt.Errorf("reading input failed: %w", scanner.Err())
			}

			return nil
		}

		line := scanner.Text()

		if line == "exit" {
			return nil
		}

		fmt.Fprintln(out, line)
	}
}
