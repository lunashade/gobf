package bf

import (
	"fmt"
	"io"
)

const (
	INCR  = '+'
	DECR  = '-'
	NEXT  = '>'
	PREV  = '<'
	PUT   = '.'
	GET   = ','
	BEGIN = '['
	END   = ']'
)
type Program struct {
	in   io.Reader
	out  io.Writer
	data []byte
	at   int
}

func NewProgram(in io.Reader, out io.Writer) *Program {
	prog := new(Program)
	prog.in = in
	prog.out = out
	prog.data = make([]byte, 1)
	prog.at = 0
	return prog
}

func (prog *Program) Execute(src []rune) {
	for pc := 0; pc < len(src); pc++ {
		switch src[pc] {
		case INCR:
			prog.data[prog.at]++
		case DECR:
			prog.data[prog.at]--
		case NEXT:
			prog.at++
			if prog.at <= len(prog.data) {
				prog.data = append(prog.data, 0)
			}
		case PREV:
			if prog.at > 0 {
				prog.at--
			}
		case PUT:
			fmt.Fprintf(prog.out, "%c", prog.data[prog.at])
		case GET:
			b := make([]byte, 1)
			prog.in.Read(b)
			prog.data[prog.at] = b[0]
		case BEGIN:
			if prog.data[prog.at] != 0 {
				break
			}
			for depth := 1; depth > 0; {
				pc++
				switch src[pc] {
				case BEGIN:
					depth++
				case END:
					depth--
				}
			}
		case END:
			if prog.data[prog.at] == 0 {
				break
			}
			for depth := 1; depth > 0; {
				pc--
				switch src[pc] {
				case BEGIN:
					depth--
				case END:
					depth++
				}
			}
		}
	}
}
