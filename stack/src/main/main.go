package main

import "fmt"
import "stack"


func main()  {
    var s stack.Stack
    for i := 0; i<10; i++ {
        s.Push(i)
    }

    fmt.Println(s)
}
