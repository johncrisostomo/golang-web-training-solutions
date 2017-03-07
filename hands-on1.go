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


