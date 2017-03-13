package main

import (
  "os"
  "log"
  "encoding/csv"
  "html/template"
)


type table struct {
  Date string
  Open string 
}

var tpl *template.Template

func init() {
  tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
  file, err := os.Open("table.csv")
  if err != nil {
    log.Fatalln(err)
  }
  defer file.Close()

  r := csv.NewReader(file)
  records, err := r.ReadAll()
  if err != nil {
    log.Fatalln(err)
  }

  data := []table{}

  for i := 1; i < len(records); i++ {
    data = append(data, table{ records[i][0], records[i][1] })
  }

  err = tpl.Execute(os.Stdout, data)
  if err != nil {
    log.Fatalln(err)
  }
}
