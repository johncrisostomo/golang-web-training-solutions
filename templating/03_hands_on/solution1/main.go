package main

import (
  "log"
  "os"
  "text/template"
)

type item struct {
  Name string
  Description string
  Meal string
  Price float32
}

var tpl *template.Template

func init() {
  tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
  items := []item{
    {
      "Egg",
      "Choice of either scrambled or poached",
      "Breakfast",
      10.00,
    },
    {
      "Adobo",
      "Popular Filipino Cuisine",
      "Lunch",
      30.00,
    },
    {
      "Angus Steak",
      "Steak with side dish",
      "Dinner",
      40.00,
    },
  }

  err := tpl.Execute(os.Stdout, items)
  if err != nil {
    log.Fatalln(err)
  }
}
