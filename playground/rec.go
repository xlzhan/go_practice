// rec.go
package main

import "fmt"

func rec(i int) {
    if i == 10 {
        return
    }
    rec(i+1)
    fmt.Println(i)
}

func main() {
    rec(0)
}
