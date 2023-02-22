package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"net"
)

func main() {
	//Escuchar un request
	ln, err := net.Listen("tcp", ":9000")
	if err != nil {
		panic(err)
	}

	defer ln.Close()
	var num int32

	for {

		//Recibir un request
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}

		err = binary.Read(conn, binary.LittleEndian, &num)

		io.WriteString(conn, fmt.Sprintf("I'm the server, the number I received is: %v", num))

		conn.Close()
	}
}
