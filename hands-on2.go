// HANDS ON 2
// create a struct that holds person fields
// create a struct that holds secret agent fields and embeds person type
// attach a method to person: pSpeak
// attach a method to secret agent: saSpeak
// create a variable of type person
// create a variable of type secret agent
// print a field from person
// run pSpeak attached to the variable of type person
// print a field from secret agent
// run saSpeak attached to the variable of type secret agent
// run pSpeak attached to the variable of type secret agent

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
