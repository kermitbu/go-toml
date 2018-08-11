package tomlx

import "fmt"

type InputStream []byte

var pos = 0
var line = 1
var col = 0

func (input InputStream) next() byte {
	var ch = input[pos]
	pos++
	if ch == '\n' {
		line++
		col = 0
	} else {

		col++
	}
	return ch
}

func (input InputStream) peek() byte {
	return input[pos]
}

func (input InputStream) eof() bool {
	return len(input) == pos
}
func (input InputStream) croak(msg string) error {
	fmt.Println("")
	//  new Error(msg + " (" + line + ":" + col + ")");
	return fmt.Errorf("%s (%d:%d)", msg, line, col)
}
