package tsm

import (
	"net"
	"bufio"
	"gopkg.in/sevansd/tsm.v1/process"
	"gopkg.in/sevansd/tsm.v1/models"
	"log"
	textTemplate "text/template"
)

func main() {
	ln, _ := net.Listen("tcp", ":8081");
	conn, _ := ln.Accept();
	db, err := models.NewDB("")
	logger := &log.Logger{}
	if (err != nil) {
		logger.Println("Error in database initializing ")
		logger.Println(err)
	}
	template := &textTemplate.Template{}
	env := &process.Env{db, logger, template}
	for {
		message, _ := bufio.NewReader(conn).ReadBytes('\n');
		env.HandleMessage(message)
	}
}