package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"net"
)

func main() {

	ln, err := net.Listen("tcp", ":9000")
	if err != nil {
		panic(err)
	}

	defer ln.Close()
	var num int32

	for {

		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}

		err = binary.Read(conn, binary.LittleEndian, &num)

		res := isPrime(num)

		if res == true {

			io.WriteString(conn, fmt.Sprintf("I'm the server, the number received %v is prime", num))

		} else {

			io.WriteString(conn, fmt.Sprintf("I'm the server, the number received %v is not prime", num))

		}

		conn.Close()
	}
}

func isPrime(n int32) bool {

	var newNum = int(n)

	var divisible int

	divisible = 0

	for i := 2; i <= newNum/2; i++ {

		if newNum%i == 0 {

			divisible++

		}
	}

	if divisible == 0 {

		fmt.Printf("o %d é primo", n)

		return true

	}
	return false
}
