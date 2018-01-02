package process

import (
	"database/sql"
	"log"
	"text/template"
)

type Env struct {
	db        *sql.DB
	logger    *log.Logger
	templates *template.Template
}
