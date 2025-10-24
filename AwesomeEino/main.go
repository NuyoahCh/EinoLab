package main

import (
	"context"
	"os"

	"github.com/cloudwego/eino-ext/components/model/ark"
	"github.com/cloudwego/eino/components/prompt"
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

	template := prompt.FromMessages(schema.FString,
		schema.SystemMessage("你是一个{role}"),
		&schema.Message{
			Role:    schema.User,
			Content: "提醒我每天{task}",
		},
	)

	params := map[string]any{
		"role": "爱睡觉的老爸",
		"task": "好好学习，天天向上",
	}

	// input := []*schema.Message{
	// 	schema.SystemMessage("你是一个爱唠叨的老妈"),
	// 	schema.UserMessage("提醒我每天早睡早起"),
	// }

	message, err := template.Format(ctx, params)

	response, err := model.Generate(ctx, message)
	if err != nil {
		panic(err)
	}
	print(response.Content)

	// respose, err := model.Generate(ctx, input)
	// if err != nil {
	// 	panic(err)
	// }
	// print(respose.Content)

	// reader, err := model.Stream(ctx, input)
	// if err != nil {
	// 	panic(err)
	// }
	// defer reader.Close()

	// for {
	// 	chunk, err := reader.Recv()
	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	print(chunk.Content)
	// }
}
