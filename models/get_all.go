package models

import (
	"github.com/Monteiro712/api-go/db"
	//"golang.org/x/tools/go/analysis/passes/defers"
)
//slice de todos
func GetAll() (todos []Todo, err error) {
	conn, err := db.OpenConnection()

	if err != nil {
		return
	}
	//se nao existir erro, faz o defer para fechar a conexão após a execução do get
	defer conn.Close()

	//seleciona tudo no banco 
	rows, err := conn.Query(`SELECT * FROM todos`)

	if err != nil {
		return
	}

	//percorrer todos os itens retornados
	for rows.Next() {
	var todo Todo
	//fazer scan para passar os valores para o objeto todo
	err = rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Done)

	//se der erro, ignora e passa para o proximo item
	if err != nil {
		continue
	}

	todos = append(todos, todo)

	}
	return
}