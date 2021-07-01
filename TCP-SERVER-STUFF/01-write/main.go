package main

import (
	"fmt"
	"io"
	"log"
	"net"
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
			log.Println(err)
			continue
		}

		io.WriteString(conn, "\n TCP server says what up\n")
		fmt.Fprintln(conn, "How u doing homie?")
		fmt.Fprintf(conn, "%v", "good i hhope")

		conn.Close()
	}

}
