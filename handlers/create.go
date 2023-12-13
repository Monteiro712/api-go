package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/Monteiro712/api-go/models"
)

//é um handler HTTP para criar uma nova tarefa (Todo).
func Create(w http.ResponseWriter, r *http.Request) {
	//variável para armazenar os dados da nova tarefa.
	var todo models.Todo

	//decodificar o JSON do corpo da requisição para a estrutura Todo.
	err := json.NewDecoder(r.Body).Decode(&todo)

	//verifica se houve um erro ao decodificar o JSON.
	if err != nil {
		log.Printf("Erro ao fazer decode do JSON: %v", err)
		
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	//chama a função Insert do pacote models para inserir a nova tarefa no banco de dados.
	id, err := models.Insert(todo)

	// Criar um mapa para armazenar a resposta JSON que será enviada ao cliente.
	var resp map[string]interface{}

	// Verificar se houve um erro ao inserir a tarefa no banco de dados.
	if err != nil {
		resp = map[string]interface{}{
			"Error":   true,
			"Message": fmt.Sprintf("Ocorreu um erro ao tentar inserir: %v", err),
		}
	} else {
		//se a inserção foi bem-sucedida, criar uma resposta de sucesso.
		resp = map[string]interface{}{
			"Error":   false,
			"Message": fmt.Sprintf("Todo inserido com sucesso. ID: %d", id),
		}
	}

	//definir o cabeçalho Content-Type como application/json.
	w.Header().Add("Content-Type", "application/json")

	//codificar a resposta como JSON e enviá-la ao cliente.
	json.NewEncoder(w).Encode(resp)
}
