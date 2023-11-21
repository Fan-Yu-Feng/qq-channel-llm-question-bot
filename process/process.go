package process

import (
	"channelSdk/dto"
	"channelSdk/dto/message"
	"channelSdk/openapi"
	"context"
	"github.com/tmc/langchaingo/prompts"
	"llm/server/chatgptServer"
	"log"
	"process/dict"
	"strings"
)

type Processor struct {
	Api openapi.OpenAPI
}

// ProcessMessage 消息处理
func (p Processor) ProcessMessage(input string, data *dto.WSATMessageData) error {
	ctx := context.Background()
	// 获取命令
	cmd := strings.Replace(message.ParseCommand(input).Cmd, "/", "", -1)
	var resp string

	if dict.GetAnswerType(cmd) == "-1" && dict.GetQuestionType(cmd) == "-1" {
		resp = "您好，你的输入不符合指令要求，请按照指令要求执行！"
	}
	var answer string
	if strings.Contains(cmd, "答案") {
		answer = strings.Replace(input, "/答案", "", 1) // 使用1表示只替换第一个匹配项
		cmd = "答案"
	}
	if dict.GetAnswerType(cmd) != "-1" {
		// 处理答案和统计信息
		resp = processAnswerMsg(cmd, answer)
	}

	if dict.GetQuestionType(cmd) != "-1" {
		// 处理问题指令
		resp = processQuestion(cmd)
	}

	log.Printf("llm 模型回复结果：%s", resp)
	toCreate := &dto.MessageToCreate{
		Content: resp,
		MessageReference: &dto.MessageReference{
			MessageID:             data.ID,
			IgnoreGetMessageError: true,
		},
	}
	// 发送消息到qq渠道
	p.sendMsg(ctx, data.ChannelID, toCreate)
	return nil
}

func processQuestion(cmd string) string {
	prompt := prompts.NewPromptTemplate("你是一个智能问答机器人，你需要出一道和{{.question}}相关的题目，并给出四个选项，分别为A、B、C、D，其中必须有一个答案是正确选项。"+
		"用户会根据你给出的选项进行选择，回复你提问的问题，如果选项不在范围内（忽略大小写），则需要提醒用户选择正确的选项。"+
		"如果选项在提供的范围内，那么按以下方式处理："+
		"如果用户答错了，你需要回复，'很遗憾，您答错了'。"+
		"如果用户答对了，你需要回复，'恭喜您答对了！！！'。"+
		"在回复用户后，需给出答案解析！。",
		[]string{"question"},
	)
	result, _ := prompt.Format(map[string]any{
		"question": cmd,
	})
	var resp string
	resp, _ = chatgptServer.GetLlmMsg(result)

	return resp
}

func processAnswerMsg(cmd string, answer string) string {
	var resp string
	switch cmd {
	case "答案":
		prompt := prompts.NewPromptTemplate(
			"你好，我的答案是{{.answer}}",
			[]string{"answer"},
		)
		result, _ := prompt.Format(map[string]any{
			"answer": answer,
		})
		resp, _ = chatgptServer.GetLlmMsg(result)
	}
	return resp
}

// 发送消息
func (p Processor) sendMsg(ctx context.Context, channelID string, toCreate *dto.MessageToCreate) {
	_, err := p.Api.PostMessage(ctx, channelID, toCreate)
	if err != nil {
		log.Println(err)
	}
}
