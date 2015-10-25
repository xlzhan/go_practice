// callback.go

package main

import "fmt"

func print_me(i int) {
    fmt.Println(i)
}

func callback(i int, f func(int)) {
    fmt.Println("I am going to print the value...")
    f(i)
}

func main() {
    callback(10, print_me)
}
