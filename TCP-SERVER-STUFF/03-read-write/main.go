package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("err")
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			//fix error
			continue
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		fmt.Fprint(conn, "I heard u say:\n", ln)
	}

	defer conn.Close()

	fmt.Println("code wont get here")
}
