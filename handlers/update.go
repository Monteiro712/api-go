package handlers

import (
	"log"
	"net/http"
	"strconv"
	"encoding/json"
	"github.com/Monteiro712/api-go/models"
	"github.com/go-chi/chi/v5"
)

//update é um handler HTTP para atualizar um registro por ID.
func Update(w http.ResponseWriter, r *http.Request) {
	//obter o parâmetro "id" da URL e converter para um número inteiro.
	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	//verificar se houve um erro ao converter o ID.
	if err != nil {
		log.Printf("erro ao fazer parse do id: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	//criar uma variável para armazenar os dados atualizados.
	var todo models.Todo

	//decodificar o JSON do corpo da requisição para a estrutura Todo.
	err = json.NewDecoder(r.Body).Decode(&todo)

	//verificar se houve um erro ao decodificar o JSON.
	if err != nil {
		log.Printf("erro ao fazer decode do json: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	//chamar a função Update do pacote models para atualizar o registro no banco de dados.
	rows, err := models.Update(int64(id), todo) 

	//verificar se houve um erro ao atualizar o registro.
	if err != nil {
		log.Printf("erro ao atualizar registro: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	//se mais de uma linha foi atualizada, registrar um aviso.
	if rows > 1 {
		log.Printf("erro: foram atualizados %d registros", rows)
	}

	//criar um mapa para armazenar a resposta JSON que será enviada ao cliente.
	resp := map[string]interface{}{
		"Error":   false,
		"Message": "dados atualizados",
	}

	//definir o cabeçalho Content-Type como application/json.
	w.Header().Add("Content-Type", "application/json")

	//codificar a resposta como JSON e enviá-la ao cliente.
	json.NewEncoder(w).Encode(resp)
}
