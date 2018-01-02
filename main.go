package tsm

import (
	"net"
	"bufio"
)

func main() {
	ln, _ := net.Listen("tcp", ":8081");
	conn, _ := ln.Accept();
	for {
		message, _ := bufio.NewReader(conn).ReadBytes('\n');
	}
}