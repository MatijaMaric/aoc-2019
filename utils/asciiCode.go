package utils

import (
	"fmt"
	"strings"
)

type AsciiCode struct {
	code   []int
	memory []int
	input  chan int
	output chan int
}

func (ac *AsciiCode) Init(code []int) {
	ac.code = make([]int, len(code))
	copy(ac.code, code)
	ac.Reset()
}

func (ac *AsciiCode) Reset() {
	ac.memory = make([]int, len(ac.code))
	copy(ac.memory, ac.code)
	SafeClose(ac.input)
	SafeClose(ac.output)

	ac.input = make(chan int, 1024)
	ac.output = make(chan int)

	go IntCodeMachine(ac.code, ac.input, ac.output)
}

func (ac *AsciiCode) WriteInt(input int) {
	ac.input <- input
}

func (ac *AsciiCode) Write(str string) {
	for _, c := range str {
		ac.WriteInt(int(c))
	}
}

func (ac *AsciiCode) WriteLn(str string) {
	ac.Write(str)
	ac.WriteInt(int('\n'))
}

func (ac *AsciiCode) Flush() string {
	var builder strings.Builder
	for a := range ac.output {
		if a < 255 {
			builder.WriteRune(rune(a))
		} else {
			builder.WriteString(fmt.Sprintf("%d", a))
		}
	}
	ac.Reset()
	return builder.String()
}
