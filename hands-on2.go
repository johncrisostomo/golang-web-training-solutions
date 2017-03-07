package main

import (
  "fmt"
)

type person struct {
  name string
  age int
}

func (p person) pSpeak() {
  fmt.Println("This is a person speaking.")
}

type secretAgent struct {
  p person
  skill string
}

func (sa secretAgent) saSpeak() {
  fmt.Println("This is a secret agent speaking.")
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
    "stealth",
  }
  fmt.Println("Person's name is ", p.name)
  p.pSpeak()
  fmt.Println("Secret Agent's skill is ", s.skill)
  s.saSpeak()
  s.p.pSpeak()
}
