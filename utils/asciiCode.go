package utils

import (
	"fmt"
	"strings"
)

type AsciiCode struct {
	intcode IntCodeVM
}

func (ac *AsciiCode) Wrap(vm IntCodeVM) {
	ac.intcode = vm
}

func (ac *AsciiCode) InitFromFile(path string) {
	ac.intcode.InitFromFile(path)
}

func (ac *AsciiCode) Init(code []int) {
	ac.intcode.Init(code)
}

func (ac *AsciiCode) Reset() {
	ac.intcode.Reset()
}

func (ac *AsciiCode) WriteInt(input int) {
	ac.intcode.Write(input)
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
	for _, a := range ac.intcode.Flush() {
		if a < 255 {
			builder.WriteRune(rune(a))
		} else {
			builder.WriteString(fmt.Sprintf("%d", a))
		}
	}
	return builder.String()
}
