package main

import (
	"context"
	"os"

	"github.com/cloudwego/eino-ext/components/model/ark"
	"github.com/cloudwego/eino/schema"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Failed to load .env file")
	}
	ctx := context.Background()
	model, err := ark.NewChatModel(ctx, &ark.ChatModelConfig{
		APIKey: os.Getenv("ARK_API_KEY"),
		Model:  os.Getenv("ARK_CHAT_MODEL"),
	})

	input := []*schema.Message{
		schema.SystemMessage("你是一个爱唠叨的老妈"),
		schema.UserMessage("提醒我每天早睡早起"),
	}
	respose, err := model.Generate(ctx, input)
	if err != nil {
		panic(err)
	}
	print(respose.Content)
}
