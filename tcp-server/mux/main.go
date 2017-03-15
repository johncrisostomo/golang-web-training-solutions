package main

import (
  "fmt"
  "net"
  "log"
  "strings"
  "bufio"
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
  request := strings.Fields(scanner.Text())
  method := request[0]
  url := request[1]

  if method == "GET" && url == "/" {
    serveRoot(conn)
  } else if method == "GET" && url == "/johncrisostomo" {
    serveHomePage(conn)
  } else {
    serveNotFound(conn, url)
  }
}

func serveNotFound(conn net.Conn, url string) {
  body := "<html><head><title>NOT FOUND!</title></head><body>" +
    "<h1>404</h1>" + url + " is not found on this server.</body></html>"

  fmt.Fprint(conn, "HTTP/1.1 200 OK \r\n")
  fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
  fmt.Fprint(conn, "Content-Type: text/html\r\n")
  fmt.Fprint(conn, "\r\n")
  fmt.Fprint(conn, body)
  conn.Close()
}

func serveRoot(conn net.Conn) {
  body := "<html><head><title>MUX ROOT</title></head><body><h1>" +
    "WELCOME TO MUX ROOT</h1></body></html>"

  fmt.Fprint(conn, "HTTP/1.1 200 OK \r\n")
  fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
  fmt.Fprint(conn, "Content-Type: text/html\r\n")
  fmt.Fprint(conn, "\r\n")
  fmt.Fprint(conn, body)
  conn.Close()
}

func serveHomePage(conn net.Conn) {
  body := "<html><head><title>JOHN CRISOSTOMO'S PAGE</title></head>" +
    "<body><h1>JOHN CRISOSTOMO'S HOME PAGE</h1>" +
    "<a href='https://blog.johncrisostomo.com'>BLOG</a></body></html>"

  fmt.Fprint(conn, "HTTP/1.1 200 OK \r\n")
  fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
  fmt.Fprint(conn, "Content-Type: text/html\r\n")
  fmt.Fprint(conn, "\r\n")
  fmt.Fprint(conn, body)
  conn.Close()
}
