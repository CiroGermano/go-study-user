
package bd

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql" // Não sei por que preciso do _ no início da linha, mas sem ele dá erro
	"github.com/CiroGermano/go-study-user/models"
	"github.com/CiroGermano/go-study-user/tools"
)

func SignUp(sig models.SignUp) error {
	fmt.Println("Iniciando o SignUp")

	err := DbConnect()
	if err != nil {
		return err
	}
	defer Db.Close()

	statement := "INSERT INTO users (user_email, user_uuid, user_dataadd) VALUES ('"+sig.UserEmail+"', '"+sig.UserUUID+"', '"+tools.CloseMySQL()+"')"
	fmt.Println("Statement: " + statement)

	_, err = Db.Exec(statement)
	if err != nil {
		fmt.Println("Erro ao inserir o usuário: " + err.Error())
		return err
	}

	fmt.Println("Usuário inserido com sucesso")
	return nil
}