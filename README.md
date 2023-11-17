# qq-channel-llm-question-bot
基于 LLM + qq频道 + openai 的智能问答机器人

## 使用说明

将 `config_template.yaml` 复制为 `config.yaml` 将您的配置信息配置上去

```yaml
# 在这个配置文件中补充你的 appid 和 botgo token，并修改文件名为 config.yaml
appid : qq渠道机器人id
token : "qq渠道机器人token"
openAIKey : "open api key"
```

### 环境变量

- [Go](https://golang.google.cn/doc/install)：1.21

### 运行

```bash
go run main.go
```

## 项目架构概述

分为三大模块：LLM、process、ChannelSdk

### LLM（Large Language Model）

LLM是大型语言模型，使用LangChainGo作为主要语言模型开发交互平台。该模型负责处理自然语言输入，执行语言理解和生成相关回复。通过 LLM 提供的 prompt 生成对应的提示词，通过 memory 模块，存储模型历史消息上下文。

### Process模块

Process模块是项目的中间层，负责处理来自QQ频道机器人的信息。它充当信息传递和处理的桥梁，与LLM进行交互。在接收到信息后，Process模块调用LLM生成相应的语言回复，并将处理后的消息返回给QQ频道。

### ChannelSdk模块

ChannelSdk模块是QQ频道的SDK（软件开发工具包），用于与QQ频道的机器人建立连接。它提供了与QQ频道的通信接口，允许项目与QQ频道进行信息交换。

## 设计 && 使用文档 

[频道机器人方案设计文档](https://github.com/Fan-Yu-Feng/qq-channel-llm-question-bot/blob/master/doc/%E8%B6%A3%E5%91%B3%E9%97%AE%E7%AD%94%E9%A2%91%E9%81%93%E6%9C%BA%E5%99%A8%E4%BA%BA%E6%96%B9%E6%A1%88%E8%AE%BE%E8%AE%A1%E6%96%87%E6%A1%A3.md)

[频道机器人使用文档](https://github.com/Fan-Yu-Feng/qq-channel-llm-question-bot/blob/master/doc/%E8%B6%A3%E5%91%B3%E9%97%AE%E7%AD%94%E9%A2%91%E9%81%93%E6%9C%BA%E5%99%A8%E4%BA%BA%E4%BD%BF%E7%94%A8%E8%AF%B4%E6%98%8E%E6%96%87%E6%A1%A3.md)

## feature

思考优化：

- [ ] 答案可以通过指令的方式交互，提供给用户更好的反馈，而不需要通过指令去输入。
- [ ] 机器人的题目类型扩展性很大，不止我提供的枚举。
- [ ] 可以接入更多的模型，在配置文件指定使用的模型，或者多模型一起使用，在单个模型挂掉的时候，使用其他模型，提供更好的容错性





## 参考使用

- [langchainGo](https://tmc.github.io/langchaingo/docs/)
- [开发说明 | QQ机器人文档](https://bot.q.qq.com/wiki/develop/api/)
- [tencent-connect/botgo: QQ频道机器人 GOSDK (github.com)](https://github.com/tencent-connect/botgo)

