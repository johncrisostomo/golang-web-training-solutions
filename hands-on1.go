// HANDS ON 1
// create a type square
// create a type circle
// attach a method to each that calculates area and returns it
// create a type shape which defines an interface as anything which has the area method
// create a func info which takes type shape and then prints the area
// create a value of type square
// create a value of type circle
// use func info to print the area of square
// use func info to print the area of circle

package main

import (
  "fmt"
  "math"
)

type square struct {
  side float32 
}

type circle struct {
  radius float32
}

func (s square) area() {
  fmt.Println("Square's area is ", s.side * s.side)
}

func (c circle) area() {
  fmt.Println("Circle's area is ", (c.radius * c.radius) * math.Pi)
}

type shape interface {
  area()
}

func info(sh shape) {
  sh.area()
}

func main() {
  sq := square{
    5,
  }

  ci := circle {
    20,
  }

  info(sq)
  info(ci)
}


