package main

import (
	"context"
	"fmt"
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
	apiKey := os.Getenv("ARK_API_KEY")
	if apiKey == "" {
		log.Fatal("ARK_API_KEY 未配置，请检查环境变量或 .env 文件")
	}

	modelName := os.Getenv("ARK_MODEL")
	if modelName == "" {
		// 官方默认可用的模型，若需要更高规格请在控制台开通并设置 ARK_MODEL
		modelName = "doubao-1.5-lite-32k"
	}

	ctx := context.Background()

	// 初始化模型
	model, err := ark.NewChatModel(ctx, &ark.ChatModelConfig{
		APIKey: apiKey,
		Model:  modelName,
	})

	input := []*schema.Message{
		schema.SystemMessage("你是一个小助手"),
		schema.UserMessage("帮我写一首关于春天的诗"),
	}
	response, err := model.Generate(ctx, input)
	if err != nil {
		log.Fatal(fmt.Errorf("failed to generate response with model %s: %w", modelName, err))
	}
	print(response.Content)
}
