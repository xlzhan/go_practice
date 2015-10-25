// stack.go

package stack

import "strconv"

type Stack struct {
    i int // 栈顶位置，0-9
    data [10]int  // 栈内元素
}

func (s *Stack) Push(i int) {
    if s.i == 9 {
        return
    }
    s.data[s.i] = i
    s.i++
}

func (s *Stack) Pop() (ret int) {
    if s.i == 0 {
        return -1
    }
    s.i--
    ret = s.data[s.i]
    return
}

func (s Stack) String() (string) {
    var str string
    for i := 0; i<s.i; i++ {
        str = str + "[" + strconv.Itoa(i) + ":" + strconv.Itoa(s.data[i]) + "]"
    }
    return str
}
