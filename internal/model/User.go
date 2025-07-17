package model

type User struct {
	ID       string `dynamodbav:"id"`
	Name     string `dynamodbav:"name"`
	Email    string `dynamodbav:"email"`
	Password string `dynamodbav:"password"`
}
