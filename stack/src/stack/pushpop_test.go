package stack

import "testing"

func TestPushPop(t *testing.T) {
    c := new(Stack)
    c.Push(5)
    pop_value := c.Pop()
    t.Log(pop_value)
    if pop_value != 5 {
        t.Log("Pop does not give 5!")
        t.Fail()
    }
}
