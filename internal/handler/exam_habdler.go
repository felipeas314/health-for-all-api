package handler

import (
	"io"

	"github.com/85labs/health-for-all-api/internal/service"
	"github.com/gofiber/fiber/v2"
)

func UploadExam(c *fiber.Ctx) error {
	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "arquivo não enviado"})
	}

	f, err := file.Open()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "falha ao abrir o arquivo"})
	}
	defer f.Close()

	buf, err := io.ReadAll(f)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "falha ao ler o arquivo"})
	}

	userEmail := c.Locals("user_email").(string) // pego do token
	exam, err := service.ProcessExamImage(buf, file.Filename, userEmail)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(exam)
}

func UploadExamPDF(c *fiber.Ctx) error {
	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "arquivo não enviado"})
	}

	src, err := file.Open()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "erro ao abrir o arquivo"})
	}
	defer src.Close()

	userEmail := c.Locals("user_email").(string)

	// Passa para o service
	exam, err := service.ProcessExamPDF(src, file.Size, file.Filename, userEmail)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(exam)
}
