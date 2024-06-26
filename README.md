# gpt
将github.com/sashabaranov/go-openai封装为Goner提供给Gone框架使用



## 使用和编写示例
1. 在配置文件中，写入配置项，参考[通过内置Goners支持配置文件](https://goner.fun/zh/guide/config.html)
2. 在需要使用的结构体中注入ChatGPT并调用相关接口，代码示例如下：

```go
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
```