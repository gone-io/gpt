package gpt

import (
	"context"
	"github.com/gone-io/gone"
	"github.com/gone-io/gone/goner/config"
	"github.com/sashabaranov/go-openai"
)

type ChatGPT interface {
	CreateCompletion(
		ctx context.Context,
		request openai.CompletionRequest,
	) (response openai.CompletionResponse, err error)

	CreateChatCompletion(
		ctx context.Context,
		request openai.ChatCompletionRequest,
	) (response openai.ChatCompletionResponse, err error)
}

// NewChatGPTClient returns a new Goner which is ChatGPT client.
func NewChatGPTClient() (gone.Goner, gone.GonerId) {
	return &chatGPTClient{}, "gone-gpt"
}

// Priest 用于埋葬chatGPTClient和其依赖的Goner
func Priest(cemetery gone.Cemetery) error {

	//使用config.Priest来埋葬Gone配置模块相关的Goner
	_ = config.Priest(cemetery)

	//埋葬chatGPTClient
	cemetery.Bury(NewChatGPTClient())
	return nil
}

type chatGPTClient struct {
	gone.Flag
	*openai.Client

	//从配置文件中注入`openai.base`配置项
	openaiBase string `gone:"config,openai.base"`

	//从配置文件中注入`openai.token`配置项
	openaiToken string `gone:"config,openai.token"`
}

func (g *chatGPTClient) AfterRevive() error {
	conf := openai.DefaultConfig(g.openaiToken)
	conf.BaseURL = g.openaiBase
	g.Client = openai.NewClientWithConfig(conf)
	return nil
}
