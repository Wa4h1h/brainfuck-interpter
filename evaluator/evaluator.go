package evaluator

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
)

type Evaluator struct {
	memory []byte
	ptr    int
}

func New(memorySize int) *Evaluator {
	return &Evaluator{
		memory: make([]byte, memorySize),
		ptr:    0,
	}
}

func (e *Evaluator) Evaluate(in io.ByteReader) error {
	r := bufio.NewReader(os.Stdin)

	for {
		b, err := in.ReadByte()
		if err != nil {
			if errors.Is(err, io.EOF) {
				return nil
			}

			return fmt.Errorf("reading byte %s produced error: %w", string(b), err)
		}

		switch b {
		case '+':
			e.memory[e.ptr]++
		case '-':
			e.memory[e.ptr]--
		case '>':
			if e.ptr < len(e.memory)-1 {
				e.ptr++
			}
		case '<':
			if e.ptr > 0 {
				e.ptr--
			}
		case '.':
			fmt.Fprintln(os.Stdout, string(e.memory[e.ptr]))
		case ',':

			for {
				rb, err := r.ReadByte()
				if err != nil {
					return fmt.Errorf("reading byte %s produced error: %w", string(b), err)
				}

				if rb == '\n' {
					continue
				}

				e.memory[e.ptr] = rb
				break
			}

		case '[':
			fmt.Println("[")
		case ']':
			fmt.Println("]")
		default:
			// skip character
		}
	}
}
