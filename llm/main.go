package main

import (
	"llm/server/baidusever"
	"llm/server/messagesever"
	"llm/server/wssever"
)

func main() {
	baidusever.InitEnv()
	dataChan := wssever.GetWsResMessage()
	for data := range dataChan {
		go messagesever.SendMsg(data)
	}

}
