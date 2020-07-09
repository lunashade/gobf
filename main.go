package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"gobf/bf"
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
	prog := bf.NewProgram(os.Stdin, os.Stdout)
	prog.Execute([]rune(contents.String()))
}
