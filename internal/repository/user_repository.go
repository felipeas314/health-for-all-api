package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/85labs/health-for-all-api/internal/database"
	"github.com/85labs/health-for-all-api/internal/model"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"

	smithy "github.com/aws/smithy-go"
)

const userTable = "users-homolog"

func SaveUser(user *model.User) error {
	item, err := attributevalue.MarshalMap(user)
	if err != nil {
		return err
	}

	_, err = database.ClientDynamoDB.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName:           aws.String(userTable),
		Item:                item,
		ConditionExpression: aws.String("attribute_not_exists(email)"),
	})

	if err != nil {
		var oe smithy.APIError
		if errors.As(err, &oe) && oe.ErrorCode() == "ConditionalCheckFailedException" {
			return errors.New("usuário já existe")
		}

		return fmt.Errorf("erro ao salvar usuário: %w", err)
	}

	return nil
}

func GetUserByEmail(email string) (*model.User, error) {
	out, err := database.ClientDynamoDB.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: aws.String(userTable),
		Key: map[string]types.AttributeValue{
			"email": &types.AttributeValueMemberS{Value: email},
		},
	})
	if err != nil {
		return nil, err
	}

	if out.Item == nil {
		return nil, errors.New("usuário não encontrado")
	}

	var user model.User
	err = attributevalue.UnmarshalMap(out.Item, &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
