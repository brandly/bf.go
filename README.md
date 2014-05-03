# bf.go

idk how to write golang or an interpreter. basically a port of @evanhahn's [brainfuck-interpreter](https://github.com/EvanHahn/brainfuck-interpreter).

```shell
$ go run bf.go test/hello.bf
Hello World!
```

reads input from stdin

```shell
$ go run bf.go test/cat.bf < test/hello.bf
++++++++[>++++[>++>+++>+++>+<<<<-]>+>+>->>+[<]<-]>>.>---.+++++++..+++.>>.<-.<.+++.------.--------.>>+.>++.
```
