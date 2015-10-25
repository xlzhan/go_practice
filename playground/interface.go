package main

import "fmt"

type Integer int

func (a *Integer) Less(b Integer) bool {
    if *a < b {
        return true
    }
    return false
}

func (a *Integer) Add(b Integer) {
    *a += b
}

type LessAdder interface {
    Less(b Integer) bool
    Add(b Integer)
}

func main() {
    var a Integer = 1
    // b 是一个更笼统的类型，任何实现了Less和Add方法的类型都可以称之为b
    var b LessAdder = &a
    a.Add(2)
    b.Add(2)
    fmt.Println(a)
    fmt.Println(b)
}
