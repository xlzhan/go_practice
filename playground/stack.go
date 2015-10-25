// stack.go

package main

import "fmt"
import "strconv"

type stack struct {
    i int // 栈顶位置，0-9
    data [10]int  // 栈内元素
}

func (s *stack) push(i int) {
    if s.i == 9 {
        return
    }
    s.data[s.i] = i
    s.i++
}

func (s *stack) pop() {
    if s.i == 0 {
        return
    }

    s.i--
}

func (s stack) String() (string) {
    var str string
    for i := 0; i<s.i; i++ {
        str = str + "[" + strconv.Itoa(i) + ":" + strconv.Itoa(s.data[i]) + "]"
    }
    return str
}

func main() {
    var s stack
    count := 20
    for j := 0; j < count; j++ {
        s.push(j)
    }
    s.pop()

    fmt.Println(s)
}
