package repository

import (
	"fmt"

	"go/go/src/github.com/danilocarrarac/desafio.mercado.livre/models"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func GetDBitem(satellite models.DBconsult) models.DBconsult {

	fmt.Printf("\nvalor de satelite no  DENTRO DO FUNC: %#v: \n", satellite)

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1")},
	)
	// Create DynamoDB client
	svc := dynamodb.New(sess)

	tableName := "mercadoLivre"

	result, err := svc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"Name": {
				S: aws.String(satellite.Name),
			},
			"Distance": {
				N: aws.String(fmt.Sprintf("%f", satellite.Distance)),
			},
		},
	})
	// if err != nil {
	// 	log.("Got error calling GetItem: %s", err)
	// }

	item := models.DBconsult{}

	err = dynamodbattribute.UnmarshalMap(result.Item, &item)
	if err != nil {
		panic(fmt.Sprintf("Failed to unmarshal Record, %v", err))
	}

	fmt.Println("Found item:")
	fmt.Println("Name:  ", item.Name)
	fmt.Println("Message: ", item.Message)
	fmt.Println("Distance:  ", item.Distance)

	return item

}
