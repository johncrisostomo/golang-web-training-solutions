package main

import (
  "os"
  "text/template"
  "log"
)

type restaurant struct {
  Name string
  Address string
  Region string 
  Meals []meal
}

type item struct {
  Name string
  Description string
  Price float32
}

type meal struct {
  Name string
  Items []item
}

var tpl *template.Template

func init() {
  tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
  meals := []meal{
    {
      "Breakfast",
      []item{
        {
          "Eggs",
          "Just eggs.",
          10,
        },
        {
          "Bacon",
          "Eggs sold separately",
          15,
        },
      },
    },
    {
      "Lunch",
      []item{
        {
          "Nilagang Baka",
          "Filipino beef stew",
          45,
        },
        {
          "Pork Adobo",
          "Popular Filipino dish",
          50,
        },
      },
    },
    {
      "Dinner",
     []item{
        {
          "Rosemary Steak",
          "Comes with side dish",
          40,
        },
      },
    },
  }

  restaurants := []restaurant{
    {
      "Restaurant 1",
      "Restaurant 1 Address",
      "Restaurant 1 Region",
      meals, 
    },
    {
      "Restaurant 2",
      "Restaurant 2 Address",
      "Restaurant 2 Region",
      meals, 
    },
    {
      "Restaurant 3",
      "Restaurant 3 Address",
      "Restaurant 3 Region",
      meals, 
    },
  }

  err := tpl.Execute(os.Stdout, restaurants)
  if err != nil {
    log.Fatalln(err) 
  }
}
