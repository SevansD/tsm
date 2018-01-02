package tsm

import (
	"net"
	"bufio"
	"gopkg.in/sevansd/tsm.v1/process"
	"gopkg.in/sevansd/tsm.v1/models"
)

func main() {
	ln, _ := net.Listen("tcp", ":8081");
	conn, _ := ln.Accept();
	env := &process.Env{}
	for {
		message, _ := bufio.NewReader(conn).ReadBytes('\n');
		process.HandleMessage(message, env)
	}
}