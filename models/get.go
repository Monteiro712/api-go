package models

import (
	"github.com/Monteiro712/api-go/db"
	//"golang.org/x/tools/go/analysis/passes/defers"
)

func Get(id int64) (todo Todo, err error) {
	conn, err := db.OpenConnection()

	if err != nil {
		return
	}
	//se nao existir erro, faz o defer para fechar a conexão após a execução do get
	defer conn.Close()

	row := conn.QueryRow(`SELECT * FROM todos WHERE id=$1`, id)

	err = row.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Done)

	return
}