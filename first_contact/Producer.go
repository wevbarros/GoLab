package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"math/rand"
	"net"
	"time"
)

func main() {

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	num1 := r.Int31n(100)

	buf := new(bytes.Buffer)

	fmt.Printf("I'm the client, the generate number is: %d \n", num1)

	err := binary.Write(buf, binary.LittleEndian, num1)

	conn, err := net.Dial("tcp", "localhost:9000")
	if err != nil {
		panic(err)
	}

	buf.WriteTo(conn)

	defer conn.Close()

	bs, _ := io.ReadAll(conn)
	fmt.Println(string(bs))

}
