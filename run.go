package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
    if len(os.Args) != 2 {
        log.Fatalf("Usage %s host:port", os.Args[1])
    }
    service := os.Args[1]
    conn, err := net.Dial("tcp", service)
    checkError(err)
    result, err := readFully(conn)
    checkError(err)
    fmt.Println(string(result))
}
