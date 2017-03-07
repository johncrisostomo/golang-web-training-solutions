/*
  // HANDS ON 3
  create an interface type that both person and secretAgent implement
  declare a func with a parameter of the interface’s type
  call that func in main and pass in a value of type person
  call that func in main and pass in a value of type secretAgent
*/
package main

import (
  "fmt"
)

type person struct {
  name string
  age int
}

func (p person) speak() {
  fmt.Println("My name is", p.name)
}

type secretAgent struct {
  p person
  skill string
}

func (sa secretAgent) speak() {
  fmt.Println("I'm a secret agent. Call me", sa.p.name)
}

type hooman interface {
  speak()
}

func introduce(h hooman) {
  h.speak()
}

func main() {
  p := person{
    "John",
    20,
  }

  s := secretAgent{
    person{
      "Crisostomo",
      25,
    },
    "brute force",
  }

  introduce(p)
  introduce(s)
}
