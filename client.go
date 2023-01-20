package main

import (
	"bytes"
	"io"
	"net"
)

/*
@param conn, the connection object
@return slice of bytes, with error status
*/
func readFully(conn net.Conn) ([]byte, error)  {
    result := bytes.NewBuffer(nil)
    var buf [512]byte
    for {
        n, err := conn.Read(buf[:])
        result.Write(buf[:n])
        if err != nil {
            if err == io.EOF {
                break
            }
            return nil, err
        }
    }
    return result.Bytes(), nil
}
