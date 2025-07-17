package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"

	dynamodb "github.com/85labs/health-for-all-api/internal/database"
	"github.com/85labs/health-for-all-api/internal/handler"
	"github.com/85labs/health-for-all-api/internal/middleware"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️  Nenhum .env encontrado, variáveis de ambiente devem estar definidas no sistema")
	}
	// 1. Inicializa o cliente do DynamoDB
	dynamodb.InitDynamo()

	// 2. Inicializa o servidor HTTP
	app := fiber.New()

	// 3. Define as rotas
	api := app.Group("/api")
	api.Post("/auth/register", handler.Register)
	api.Post("/auth/login", handler.Login)

	apiProtected := api.Group("", middleware.AuthMiddleware)
	apiProtected.Post("/exam/upload/photo", handler.UploadExam)
	apiProtected.Post("/exam/upload/pdf", handler.UploadExamPDF)

	// 4. Sobe o servidor
	log.Println("🚀 Servidor rodando em http://localhost:8080")
	if err := app.Listen(":8080"); err != nil {
		log.Fatal(err)
	}
}
