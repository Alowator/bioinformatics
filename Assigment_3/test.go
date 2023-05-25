package main

import (
    sp "github.com/scipipe/scipipe"
)

func main() {
    wf := sp.NewWorkflow("hello_world", 4)

    hello := wf.NewProc("hello", "echo 'Hello ' > {o:out|.txt}")
    world := wf.NewProc("world", "echo $(cat {i:in}) World > {o:out|.txt}")

    world.In("in").From(hello.Out("out"))

    wf.Run()
}
