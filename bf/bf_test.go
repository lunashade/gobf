package bf

import (
	"bytes"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	in := new(bytes.Buffer)
	out := new(bytes.Buffer)
	prog := NewProgram(in, out)
	prog.Execute([]rune("+++++++++[>++++++++>+++++++++++>+++>+<<<<-]>.>++.+++++++..+++.  >+++++.<<+++++++++++++++.>.+++.------.--------.>+.>+. "))
	if out.String() != "Hello World!\n" {
		t.Fatal("want:\nHello World\ngot:\n", out.String())
	}

}
