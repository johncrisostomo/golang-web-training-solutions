package main

import (
  "text/template"
  "os"
  "log"
)

type Hotel struct {
  Name string
  Address string
  City string
  Zip int
  Region string
}

var tpl *template.Template

func init() {
  tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
  hotels := []Hotel{
    { "Fictitious Hotel", "Surreal St.", "Dali City", 6000, "Southern" }, 
    { "Dolphin Hotel", "Sheepman St.", "Murakami City", 2281, "Northern" }, 
  }

  err := tpl.Execute(os.Stdout, hotels)
  if err != nil {
    log.Fatalln(err)
  }
}
