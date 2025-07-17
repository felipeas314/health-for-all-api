package repository

import (
	"context"

	"github.com/85labs/health-for-all-api/internal/database"
	"github.com/85labs/health-for-all-api/internal/model"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func SaveExamResult(exam *model.Exam) error {
	item := map[string]types.AttributeValue{
		"id":         &types.AttributeValueMemberS{Value: exam.ID}, // ✅ hash key obrigatória
		"user_email": &types.AttributeValueMemberS{Value: exam.UserEmail},
		"file_name":  &types.AttributeValueMemberS{Value: exam.FileName},
		"type":       &types.AttributeValueMemberS{Value: exam.Type},
		"result":     &types.AttributeValueMemberS{Value: exam.Result},
		"created_at": &types.AttributeValueMemberS{Value: exam.CreatedAt},
	}

	_, err := database.ClientDynamoDB.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String("exams-homolog"),
		Item:      item,
	})

	return err
}
