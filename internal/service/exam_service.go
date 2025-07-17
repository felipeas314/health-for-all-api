package service

import (
	"mime/multipart"
	"time"

	"github.com/85labs/health-for-all-api/internal/model"
	"github.com/85labs/health-for-all-api/internal/repository"
	"github.com/85labs/health-for-all-api/internal/utils"
)

func ProcessExamImage(fileBytes []byte, filename, userEmail string) (*model.Exam, error) {
	text, err := utils.SendImageToGPT(fileBytes)
	if err != nil {
		return nil, err
	}

	// Aqui você pode fazer parsing + validação do resultado
	// Simulando resposta
	exam := &model.Exam{
		ID:        utils.GenerateUUID(),
		UserEmail: userEmail,
		FileName:  filename,
		Type:      "Hemograma",
		Result:    text,
		CreatedAt: time.Now().Format(time.RFC3339),
	}

	if err := repository.SaveExamResult(exam); err != nil {
		return nil, err
	}

	return exam, nil
}

func ProcessExamPDF(file multipart.File, size int64, filename, userEmail string) (*model.Exam, error) {
	text, err := utils.ExtractTextFromPDF(file, size)
	if err != nil {
		return nil, err
	}

	gptResult, err := utils.SendTextToGPT(text)
	if err != nil {
		return nil, err
	}

	exam := &model.Exam{
		ID:        utils.GenerateUUID(), // ESSENCIAL
		UserEmail: userEmail,
		FileName:  filename,
		Type:      "PDF",
		Result:    gptResult,
		CreatedAt: time.Now().Format(time.RFC3339),
	}

	if err := repository.SaveExamResult(exam); err != nil {
		return nil, err
	}

	return exam, nil
}
