package database

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

var ClientDynamoDB *dynamodb.Client

func InitDynamo() {
	cfg, err := config.LoadDefaultConfig(context.TODO())

	if err != nil {

		log.Fatalf("erro ao carregar configuração AWS: %v", err)
	}

	ClientDynamoDB = dynamodb.NewFromConfig(cfg)
}
