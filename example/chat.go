package example

import (
	"context"
	"github.com/gone-io/gone"
	"github.com/gone-io/gpt"
	"github.com/sashabaranov/go-openai"
)

type Chat struct {
	gone.Flag
	gPT gpt.ChatGPT `gone:"gone-gpt"`
}

func (c *Chat) Use(ask string) error {
	response, err := c.gPT.CreateChatCompletion(context.TODO(), openai.ChatCompletionRequest{
		Model: openai.GPT3Dot5Turbo,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: "you are a helpful chatbot",
			},
			{
				Role:    openai.ChatMessageRoleUser,
				Content: ask,
			},
		},
	})
	if err != nil {
		return err
	}

	println(response.Choices[0].Message.Content)
	return nil
}
