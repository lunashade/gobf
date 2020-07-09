package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
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

func main() {
	var contents bytes.Buffer
	for _, f := range os.Args[1:] {
		b, err := ioutil.ReadFile(f)
		if err != nil {
			fmt.Fprintln(os.Stderr, "read file: ", err)
			panic(err)
		}
		contents.Write(b)
	}
	if contents.Len() == 0 {
		log.Fatal("no input")
	}
	execute([]rune(contents.String()))
}

type Program struct {
	data []byte
	at   int
}

func execute(src []rune) {
	Len := len(src)
	prog := new(Program)
	prog.data = make([]byte, 1)
	prog.at = 0

	for pc := 0; pc < Len; pc++ {
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
			fmt.Printf("%c", prog.data[prog.at])
		case GET:
			b := make([]byte, 1)
			os.Stdin.Read(b)
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
