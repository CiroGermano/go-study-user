package main

import (
	"context"
	"os"
	"fmt"
	"errors"
	"github.com/aws/aws-lambda-go/events"
	lambda "github.com/aws/aws-lambda-go/lambda"
	"github.com/CiroGermano/go-study-user/awsgo"
	"github.com/CiroGermano/go-study-user/models"
	"github.com/CiroGermano/go-study-user/bd"
)

func main () {
	lambda.Start(ExecutarLambda)
}

func ExecutarLambda (ctx context.Context, event events.CognitoEventUserPoolsPostConfirmation) (events.CognitoEventUserPoolsPostConfirmation, error) {
	
	awsgo.InicializarAWS()

	if !ValidarPatrametros() {
		fmt.Println("Erro nos parametros. Deve enviar o SECRET_NAME")
		err := errors.New("Erro nos parametros. Deve enviar o SECRET_NAME")
		return event, err
	}

	var data models.SignUp

	for row, att := range event.Request.UserAttributes {
		switch row {
			case "email":
				data.UserEmail = att
				fmt.Println("Email: " + data.UserEmail)
			case "sub":
				data.UserUUID = att
				fmt.Println("UUID: " + data.UserUUID)
		}
	}

	err := bd.ReadSecret()
	if err != nil {
		fmt.Println("Erro ao buscar o secret: " + err.Error())
		return event, err
	}

	err = bd.SignUp(data)
	return event, err
}

func ValidarPatrametros() bool {
	var tryParam bool
	_, tryParam = os.LookupEnv("SECRET_NAME")
	return tryParam
}