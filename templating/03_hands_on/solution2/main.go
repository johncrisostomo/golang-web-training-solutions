package main

import (
  "os"
  "text/template"
  "log"
)

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

  err := tpl.Execute(os.Stdout, meals)
  if err != nil {
    log.Fatalln(err) 
  }
}
