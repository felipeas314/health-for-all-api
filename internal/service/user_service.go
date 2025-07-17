package service

import (
	"errors"

	"github.com/85labs/health-for-all-api/internal/dto"
	"github.com/85labs/health-for-all-api/internal/model"
	"github.com/85labs/health-for-all-api/internal/repository"
	"github.com/85labs/health-for-all-api/internal/utils"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(input dto.RegisterInputDTO) (*model.User, error) {
	hashed, err := utils.HashPassword(input.Password)
	if err != nil {
		return nil, err
	}

	user := &model.User{
		ID:       uuid.NewString(),
		Name:     input.Name,
		Email:    input.Email,
		Password: hashed,
	}

	err = repository.SaveUser(user)
	if err != nil {
		if err.Error() == "ConditionalCheckFailedException" {
			return nil, errors.New("usuário já existe")
		}
		return nil, err
	}

	return user, nil
}

func LoginUser(req dto.LoginRequestDTO) (string, *model.User, error) {
	user, err := repository.GetUserByEmail(req.Email)
	if err != nil {
		return "", nil, errors.New("usuário ou senha inválidos")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return "", nil, errors.New("usuário ou senha inválidos")
	}

	token, err := utils.GenerateJWT(user.Email)
	if err != nil {
		return "", nil, errors.New("erro ao gerar token")
	}

	// retorna sem o hash da senha
	user.Password = ""

	return token, user, nil
}
