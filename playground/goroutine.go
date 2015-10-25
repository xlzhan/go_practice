package main

import "fmt"
import "time"

var c chan int

func ready(w string, sec int){
    time.Sleep(time.Duration(sec) * time.Second)
    fmt.Println(w, "is ready!")
    c <- 1
}

func main() {
    c = make(chan int)
    go ready("Coffee", 2)
    go ready("Tea", 4)
    fmt.Println("I am waitting, but not too long")
    i := 0
    L: for {
        select {
        case <- c:
            i++
            if i > 1 {
                break L
            }
        }
    }
    fmt.Println("Done")
}
