package db

import (
	"database/sql"
	"fmt"
	"github.com/Monteiro712/api-go/configs"
	_ "github.com/lib/pq"
)
//abre uma conexão com o banco de dados e retorna um ponteiro para sql.DB.
func OpenConnection() (*sql.DB, error) {
	conf := configs.GetDB()
	//construir a string de conexão usando as configurações do banco de dados.
	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
	conf.Host, conf.Port, conf.User, conf.Pass, conf.Database)
	//abre uma conexão com o banco de dados PostgreSQL.
	conn, err := sql.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}
	//verifica se a conexão com o banco de dados está ativa.
	err = conn.Ping()

	return conn, err	
}