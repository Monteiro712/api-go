package handlers

import (
	"log"
	"net/http"
	"encoding/json"
	"github.com/Monteiro712/api-go/models"
)

//list é um handler HTTP para obter todos os registros.
func List(w http.ResponseWriter, r *http.Request) {
	//obter todos os registros do banco de dados.
	todos, err := models.GetAll()

	//verificar se houve um erro ao obter os registros.
	if err != nil {
		log.Printf("erro ao obter registros: %v", err)
	}

	//definir o cabeçalho Content-Type como application/json.
	w.Header().Add("Content-Type", "application/json")

	//codificar os registros como JSON e enviá-los ao cliente.
	json.NewEncoder(w).Encode(todos)
}
