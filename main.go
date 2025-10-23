package main

import (
	"context"
	"github.com/cloudwego/eino-ext/components/model/ark"
	"github.com/cloudwego/eino/schema"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	err := godotenv.Load() // 加载环境变量
	if err != nil {
		log.Fatal("Error loading .env file") // 处理加载错误
	}
	ctx := context.Background()

	// 初始化模型
	model, err := ark.NewChatModel(ctx, &ark.ChatModelConfig{
		APIKey: os.Getenv("ARK_API_KEY"),
		Model:  "doubao-1.5-pro-32k-250115",
	})

	input := []*schema.Message{
		schema.SystemMessage("你是一个小助手"),
		schema.UserMessage("帮我写一首关于春天的诗"),
	}
	response, err := model.Generate(ctx, input)
	if err != nil {
		panic(err)
	}
	print(response.Content)
}
