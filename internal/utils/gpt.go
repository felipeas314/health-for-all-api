package utils

import (
	"context"
	"encoding/base64"
	"os"

	openai "github.com/sashabaranov/go-openai"
)

func SendImageToGPT(imageData []byte) (string, error) {
	apiKey := os.Getenv("OPENAI_API_KEY")
	print(apiKey)
	client := openai.NewClient(apiKey)

	b64 := base64.StdEncoding.EncodeToString(imageData)

	req := openai.ChatCompletionRequest{
		Model: openai.GPT4o,
		Messages: []openai.ChatCompletionMessage{
			{
				Role: "user",
				MultiContent: []openai.ChatMessagePart{
					{
						Type: openai.ChatMessagePartTypeImageURL,
						ImageURL: &openai.ChatMessageImageURL{
							URL:    "data:image/jpeg;base64," + b64,
							Detail: openai.ImageURLDetailHigh,
						},
					},
					{
						Type: openai.ChatMessagePartTypeText,
						Text: "Extraia os dados clínicos relevantes desse exame em formato JSON.",
					},
				},
			},
		},
		MaxTokens: 800,
	}

	resp, err := client.CreateChatCompletion(context.Background(), req)
	if err != nil {
		print(err)
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}

func SendTextToGPT(text string) (string, error) {
	apiKey := os.Getenv("OPENAI_API_KEY")

	client := openai.NewClient(apiKey)

	req := openai.ChatCompletionRequest{
		Model: openai.GPT4o,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    "user",
				Content: "Extraia os dados clínicos desse exame em formato JSON:\n\n" + text,
			},
		},
		MaxTokens: 800,
	}

	resp, err := client.CreateChatCompletion(context.Background(), req)
	if err != nil {
		return "", err
	}
	print(resp.Choices[0].Message.Content)
	return resp.Choices[0].Message.Content, nil
}
