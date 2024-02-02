package main

import (
	"fmt"
	
	 "net"
	 "os"
)

func main() {
	
	fmt.Println("Logs from your program will appear here!")

	
	
	 l, err := net.Listen("tcp", "0.0.0.0:4221")
	 if err != nil {
	 	fmt.Println("Failed to bind to port 4221")
	 	os.Exit(1)
	 }
	defer l.Close()
	for {
		connection, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}
		handle_connection(connection)

	}
}
func handle_connection(connection net.Conn) {
	defer connection.Close()
	var buf bytes.Buffer
	var header = "HTTP/1.1 200 OK\r\n\r\n"
	buf.Write([]byte(header))

	_, err := connection.Write(buf.Bytes())
	
	 if err != nil {
	 	fmt.Println("Failed to write to socket")
	 	os.Exit(1)
	 }
}
