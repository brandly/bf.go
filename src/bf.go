package main

import (
    "fmt"
    "os"
    "io/ioutil"
)

func brainfuck(src, input string) {
    var tape = []int{0}
    var tapeIndex = 0
    fmt.Println(tape)

    for i := range src {
        var char = string(src[i])

        switch {
        case char == ">":
            tapeIndex += 1
            if len(tape) <= tapeIndex {
                fmt.Println("extending tape")
                tape = append(tape, 0)
            }

        case char == "<":
            if tapeIndex > 0 {
                tapeIndex -= 1
            }
            fmt.Println("<")

        case char == "+":
            tape[tapeIndex] += 1
            fmt.Println("+")

        case char == "-":
            tape[tapeIndex] -= 1
            fmt.Println("-")

        case char == ".":
            fmt.Println(".")

        case char == ",":
            fmt.Println(",")

        case char == "[":
            fmt.Println("[")

        case char == "]":
            fmt.Println("]")

        }
    }
}

func main() {
    content, err := ioutil.ReadFile(os.Args[1])
    if err != nil {
        fmt.Println("Error reading file")
        fmt.Println(err)
    } else {
        var src = string(content)
        var input = os.Args[2]

        brainfuck(src, input)
    }
}
