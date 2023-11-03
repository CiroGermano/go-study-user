package secretm

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/CiroGermano/go-study-user/awsgo"
	"github.com/CiroGermano/go-study-user/models"
)

func GetSecret(secretName string) (models.SecretRDSJson, error) {
	var secretData models.SecretRDSJson
	fmt.Println("Buscando o secretData: " + secretName)

	svc := secretsmanager.NewFromConfig(awsgo.Cfg)
	clave, err := svc.GetSecretValue(awsgo.Ctx, &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(secretName),
	})
	if err != nil {
		fmt.Println("Erro ao buscar o secretData: " + err.Error())
		return secretData, err
	}

	json.Unmarshal([]byte(*clave.SecretString), &secretData)
	fmt.Println("Read SecretData OK: " + secretData.Username)
	return secretData, nil
}
