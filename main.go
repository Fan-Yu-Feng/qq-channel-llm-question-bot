package main

import (
	"channelSdk/botgo"
	"channelSdk/dto"
	"channelSdk/dto/message"
	"channelSdk/event"
	"channelSdk/token"
	"channelSdk/websocket"
	"context"
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"os"
	"path"
	"process"
	"runtime"
	"strings"
	"time"
)

var processor process.Processor

func openAiEnv(openAIKey string) {
	err := os.Setenv("OPENAI_API_KEY", openAIKey)
	if err != nil {
		return
	}
}

// 入口
func main() {
	configName := "config.yaml"
	// 获取配置文件中的 appId 和 token 信息
	appId, tokenStr, openAIKey, err := getConfigInfo(configName)
	if err != nil {
		log.Fatal(err)
	}
	openAiEnv(openAIKey)
	botToken := token.BotToken(appId, tokenStr)

	// 沙箱
	//api := botgo.NewSandboxOpenAPI(botToken).WithTimeout(3 * time.Second)
	// 正式
	api := botgo.NewOpenAPI(botToken).WithTimeout(3 * time.Second)

	ctx := context.Background()
	// 获取 websocket 信息
	wsInfo, err := api.WS(ctx, nil, "")
	if err != nil {
		log.Fatal(err)
	}

	processor = process.Processor{Api: api}

	//websocket.RegisterResumeSignal(syscall.SIGUSR1)
	// 根据不同的回调，生成 intents
	intent := websocket.RegisterHandlers(
		// at 机器人事件
		ATMessageEventHandler(),
	)

	err = botgo.NewSessionManager().Start(wsInfo, botToken, &intent)
	if err != nil {
		log.Fatal(err)
	}

}

// 获取配置文件中的信息
func getConfigInfo(fileName string) (uint64, string, string, error) {
	// 获取当前go程调用栈所执行的函数的文件和行号信息
	// 忽略pc和line
	_, filePath, _, ok := runtime.Caller(1)

	if !ok {
		log.Fatal("runtime.Caller(1) 读取失败")
	}
	file := fmt.Sprintf("%s/%s", path.Dir(filePath), fileName)
	var conf struct {
		AppID     uint64 `yaml:"appid"`
		Token     string `yaml:"token"`
		OpenAIKey string `yaml:"openAIKey"`
	}
	content, err := ioutil.ReadFile(file)
	if err != nil {
		log.Print("ioutil.ReadFile() 读取失败")
		return 0, "", "", err
	}
	if err = yaml.Unmarshal(content, &conf); err != nil {
		log.Print("yaml.Unmarshal(content, &conf) 读取失败")
		return 0, "", "", err
	}
	return conf.AppID, conf.Token, conf.OpenAIKey, nil
}

// ATMessageEventHandler 实现处理 at 消息的回调
func ATMessageEventHandler() event.ATMessageEventHandler {
	return func(event *dto.WSPayload, data *dto.WSATMessageData) error {
		input := strings.ToLower(message.ETLInput(data.Content))
		return processor.ProcessMessage(input, data)
	}
}
