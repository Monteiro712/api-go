package handlers

import (
	"log"
	"net/http"
	"strconv"
	"encoding/json"
	"github.com/Monteiro712/api-go/models"
	"github.com/go-chi/chi/v5"
)

//handler HTTP para obter uma tarefa por ID.
func Get(w http.ResponseWriter, r *http.Request) {
	//obter o parâmetro "id" da URL e converter para um número inteiro.
	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	//verificar se houve um erro ao converter o ID.
	if err != nil {
		log.Printf("Erro ao fazer parse do ID: %v", err)
		//se houver um erro, retornar um status interno do servidor.
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	//chamar a função Get do pacote models para obter a tarefa do banco de dados.
	todo, err := models.Get(int64(id))

	//verificar se houve um erro ao obter a tarefa.
	if err != nil {
		log.Printf("Erro ao obter registro: %v", err)
		//se houver um erro, retornar um status interno do servidor.
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	//definir o cabeçalho Content-Type como application/json.
	w.Header().Add("Content-Type", "application/json")

	//codificar a tarefa como JSON e enviá-la ao cliente.
	json.NewEncoder(w).Encode(todo)
}
