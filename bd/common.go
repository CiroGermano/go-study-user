package bd

import (
	"database/sql"
	"os"
	"fmt"
	_ "github.com/go-sql-driver/mysql" // Não sei por que preciso do _ no início da linha, mas sem ele dá erro
	"github.com/CiroGermano/go-study-user/secretm"
	"github.com/CiroGermano/go-study-user/models"
)

var SecretModel models.SecretRDSJson
var err error
var Db *sql.DB

func ReadSecret() error {
	var err error
	SecretModel, err = secretm.GetSecret(os.Getenv("SECRET_NAME"))
	return err
}

func DbConnect() error {
	Db, err = sql.Open("mysql", ConnStr(SecretModel))
	if err != nil {
		fmt.Println("Erro ao conectar no banco de dados: " + err.Error())
		return err
	}

	err = Db.Ping()
	if err != nil {
		fmt.Println("Erro ao pingar no banco de dados: " + err.Error())
		return err
	}
	fmt.Println("Conectado com êxito no banco de dados")
	return nil
}

func ConnStr(keys models.SecretRDSJson) string {
	var dbUser, authToken, dbEndpoint, dbName string
	dbUser = keys.Username
	authToken = keys.Password
	dbEndpoint = keys.Host
	dbName = "gostudy"
	connStr :=  fmt.Sprintf("%s:%s@tcp(%s)/%s?allowCleartextPassword=true", dbUser, authToken, dbEndpoint, dbName)
	return connStr
}