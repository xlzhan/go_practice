// defer.go

package main

import "fmt"

func f() (ret int) {
    defer func() {
        ret++
    }()
    return 0
}

func main() {
    ret := f()
    fmt.Println(ret)
}
