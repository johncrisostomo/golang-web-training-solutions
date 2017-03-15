package main

import (
  "fmt"
  "net"
  "log"
  "bufio"
  "strings"
)

func main() {
  li, err := net.Listen("tcp", ":8080")
  if err != nil {
    log.Fatalln(err)
  }
  defer li.Close()

  for {
    conn, err := li.Accept()
    if err != nil {
      log.Fatalln(err)
    }

    go handle(conn)
  }
}

func handle(conn net.Conn) {
  scanner := bufio.NewScanner(conn)
  scanner.Scan()
  line := scanner.Text()
  msg := strings.Fields(line)
  fmt.Println("Request method is", msg[0])
  conn.Close()
}
