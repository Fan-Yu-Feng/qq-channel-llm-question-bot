package chatgptServer

import (
	"context"
	"fmt"
	openai1 "github.com/sashabaranov/go-openai"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/openai"
	"github.com/tmc/langchaingo/memory"
	"github.com/tmc/langchaingo/schema"
	"log"
)

var messageHistory schema.ChatMessageHistory = &memory.ChatMessageHistory{}

// var memoryBuffer memory.ConversationBuffer = memory.ConversationBuffer{}
//var memoryBuffer = memory.NewConversationBuffer()

var Token = ""
var client *openai1.Client
var ctx = context.Background()

func InitClient(token string) {
	client = openai1.NewClient(token)
}

// GetMsg 基于openAI 调用/**
func GetMsg(prompt string) (string, error) {

	if client == nil {
		InitClient(Token)
	}

	resp, err := client.CreateChatCompletion(
		ctx,
		openai1.ChatCompletionRequest{
			Model: openai1.GPT3Dot5Turbo,
			Messages: []openai1.ChatCompletionMessage{
				{
					Role:    openai1.ChatMessageRoleUser,
					Content: prompt,
				},
			},
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return "", nil
	}
	//fmt.Println(resp)
	fmt.Println(resp.Choices[0].Message.Content)
	return resp.Choices[0].Message.Content, nil
}

// GetLlmMsg 基于 LLM 模型调用，形成上下文/**
func GetLlmMsg(userMsg string) (string, error) {

	llm, err := openai.NewChat(openai.WithModel("gpt-3.5-turbo-0613"))
	if err != nil {
		log.Println(err)
		return "调用LLM 出现异常了，请及时排查！", err
	}

	humChatMessage := schema.HumanChatMessage{Content: userMsg}
	// 内存缓冲 暂无用到
	//memoryBuffer.ChatHistory.AddMessage(ctx, humChatMessage)
	messageHistory.AddMessage(ctx, humChatMessage)
	messageList, _ := messageHistory.Messages(ctx)
	// 根据上下文进行回答
	var completion, callErr = llm.Call(ctx, messageList,
		llms.WithTemperature(0.5),
		llms.WithStopWords([]string{"Armstrong"}),
	)
	if callErr != nil {
		log.Println(err)
	}
	log.Println(completion)
	// 添加机器人的回复
	messageHistory.AddMessage(ctx, completion)
	//memoryBuffer.ChatHistory.AddMessage(ctx, completion)

	return completion.GetContent(), callErr
}
