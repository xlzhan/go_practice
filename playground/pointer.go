package main

import "fmt"

type Animal struct {
    Name string
}

type Human struct {
    *Animal
}

func (a *Animal) Speak() {
    fmt.Println(a.Name)
}

func main()  {
    // a := &Animal{Name: "Jude",}
    // b := &Human{a}

    a := new(Animal)
    a.Name = "Jude"
    // b := new(Human)

    var c Animal
    c.Name = "Tom"

    d := &c
    d.Name = "Mike"

    c.Name = "xxx"

    b := a
    b.Name = "Sherry"

    fmt.Printf("a: %v\n", a)
    fmt.Printf("b: %v\n", b)
    fmt.Printf("c: %v\n", c)
    fmt.Printf("d: %v\n", d)
    fmt.Printf("*d: %v\n", *d)

    a.Speak()
    b.Speak()
    c.Speak()
    d.Speak()
}
