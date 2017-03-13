package main

import (
  "os"
  "log"
  "bufio"
  "fmt"
)

type table struct {
  Date string
  Open float32
}

func main() {
  file, err := os.Open("table.csv")
  if err != nil {
    log.Fatalln(err)
  }
  defer file.Close()
  
  reader := bufio.NewReader(file)
  scanner := bufio.NewScanner(reader)

  scanner.Split(bufio.ScanLines)

  lines := []string{}

  for scanner.Scan() {
    lines = append(lines, scanner.Text())
  }

  data := []table{}

  for i := 1; i < len(lines); i++ {
    
  }

  fmt.Printf("%s\n", lines[0])
}
