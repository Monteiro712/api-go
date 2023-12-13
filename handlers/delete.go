package handlers

import (
	"log"
	"net/http"
	"strconv"
	"encoding/json"
	"github.com/Monteiro712/api-go/models"
	"github.com/go-chi/chi/v5"
)

//handler HTTP para remover um registro por ID.
func Delete(w http.ResponseWriter, r *http.Request) {
	//obter o parâmetro "id" da URL e converter para um número inteiro.
	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	//verificar se houve um erro ao converter o ID.
	if err != nil {
		log.Printf("Erro ao fazer parse do ID: %v", err)
		
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	//chamar a função Delete do pacote models para remover o registro do banco de dados.
	rows, err := models.Delete(int64(id))

	//verificar se houve um erro ao remover o registro.
	if err != nil {
		log.Printf("Erro ao remover registro: %v", err)
		//se houver um erro, retornar um status interno do servidor.
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	//se mais de uma linha foi removida, registrar um aviso.
	if rows > 1 {
		log.Printf("Aviso: Foram removidos %d registros", rows)
	}

	//criar um mapa para armazenar a resposta JSON que será enviada ao cliente.
	resp := map[string]interface{}{
		"Error":   false,
		"Message": "Registro removido",
	}

	//definir o cabeçalho Content-Type como application/json.
	w.Header().Add("Content-Type", "application/json")

	//codificar a resposta como JSON e enviá-la ao cliente.
	json.NewEncoder(w).Encode(resp)
}
