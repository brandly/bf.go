package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func brainfuck(src string) []uint8 {
	tape := []uint8{0}
	tapeIndex := 0

	srcLength := len(src)

	for srcIndex := 0; srcIndex < srcLength; srcIndex++ {
		char := string(src[srcIndex])

		switch {
		case char == ">":
			tapeIndex += 1
			if len(tape) <= tapeIndex {
				tape = append(tape, 0)
			}

		case char == "<":
			if tapeIndex > 0 {
				tapeIndex -= 1
			}

		case char == "+":
			tape[tapeIndex] += 1

		case char == "-":
			tape[tapeIndex] -= 1

		case char == ".":
			print(string(uint8(tape[tapeIndex])))

		case char == ",":
			b := make([]byte, 1)
			os.Stdin.Read(b)
			tape[tapeIndex] = b[0]

		case char == "[":
			if tape[tapeIndex] == 0 {
				depth := 1
				for depth > 0 {
					srcIndex++
					srcCharacter := string(src[srcIndex])
					if srcCharacter == "[" {
						depth++
					} else if srcCharacter == "]" {
						depth--
					}
				}
			}

		case char == "]":
			depth := 1
			for depth > 0 {
				srcIndex--
				srcCharacter := string(src[srcIndex])
				if srcCharacter == "[" {
					depth--
				} else if srcCharacter == "]" {
					depth++
				}
			}
			srcIndex--
		}
	}

	return tape
}

func main() {
	content, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Error reading file")
		fmt.Println(err)
	} else {
		var src = string(content)

		brainfuck(src)
	}
}
