// calc.go
package main

import "os"
import "fmt"
import "simplemath"
import "strconv"

var Usage = func() {
  fmt.Println("USAGE: calc command [arguments] ... ")
}

func main() {
    args := os.Args
    if args == nil || len(args) < 2 {
        Usage()
        return
    }

    fmt.Println(len(args))
    fmt.Println(args[0])
    fmt.Println(args[1])
    fmt.Println(args[2])

    switch args[1] {
        case "add":
            if len(args) != 4 {
                fmt.Println("USAGE: calc add <int1> <int2>")
                return
            }
            v1, err1 := strconv.Atoi(args[2])
            v2, err2 := strconv.Atoi(args[3])
            if err1 != nil || err2 != nil {
                fmt.Println("USAGE: calc add <int1> <int2>")
                return
            }
            ret := simplemath.Add(v1, v2)
            fmt.Println("Result: ", ret)
        default:
            fmt.Println("Unknown command")
    }
}
