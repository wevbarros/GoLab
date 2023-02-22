package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"net"
)

func main() {

	var num1 int32
	buf := new(bytes.Buffer)

	fmt.Scanln(&num1)

	fmt.Printf("I'm the client, your input was the number: %d", num1)

	//escribir el numero a b, el buffer
	err := binary.Write(buf, binary.LittleEndian, num1)

	//Llama al server
	conn, err := net.Dial("tcp", "localhost:9000")
	if err != nil {
		panic(err)
	}

	//Escribe buffer a la conexion
	buf.WriteTo(conn)

	//Cierra la conexión
	defer conn.Close()

	//Lee la conexión  y la imprime
	bs, _ := io.ReadAll(conn)
	fmt.Println(string(bs))
}
