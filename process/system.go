package process

import (
	"database/sql"
	"log"
	"text/template"
	"encoding/json"
)

type Env struct {
	db        *sql.DB
	logger    *log.Logger
	templates *template.Template
}

func (env *Env) HandleMessage(message []byte) {
	answer := ClientMessage{}
	json.Unmarshal(message, &answer)
	if answer.message.messageType != MessageType_Info {
		// send massage into telegram
	}
	// write to database
}
