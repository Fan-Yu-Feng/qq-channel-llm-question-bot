# qq-channel-llm-question-bot
基于 LLM + qq频道 + openai 的智能问答机器人

## 使用说明

### 配置文件

将 `config_template.yaml` 复制为 `config.yaml` 将您的配置信息配置上去

```yaml
# 在这个配置文件中补充你的 appid 和 botgo token，并修改文件名为 config.yaml
appid : qq渠道机器人id
token : "qq渠道机器人token"
openAIKey : "open api key"
```

### 开发环境

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

## 流程设计

**用户交互流程：*

1. **用户选择题目类型：**
   - 用户可以通过与机器人进行交互，选择自己感兴趣的题目类型。例如，用户可以在qq频道发送指令，选择题目类型既可：如“脑筋急转弯。”
2. **机器人生成问题：**
   - 根据用户选择的题目类型，使用 LLM 计算后，生成一个问题，并提供答案选项，将其呈现给用户。问题的形式应当符合所选题目类型的特点，以确保用户体验的趣味性和挑战性。
3. **用户发起答题指令：**
   - 用户收到问题后，可以发送指令给机器人，表达自己的答案意向。例如，用户可以发送：”答案A”
4. **机器人评估答案：**
   - 机器人收到用户的答题指令后，使用 LLM 对用户的答案进行评估。这可能涉及到使用事先定义的规则、模型或算法来判断答案的正确性或相关性。
5. **机器人生成回答：**
   - 如果用户的答案是正确的，机器人将生成一个富有趣味性的回答，可以包含赞扬或进一步的解释。如果用户的答案不正确，机器人可以提供正确答案并进行解释。
6. **用户互动与反馈：**
   - 机器人与用户可以继续进行互动。用户可以选择继续答题、更换题目类型，或向机器人提出其他相关问题。机器人应当能够根据用户的反馈灵活调整互动流程。

​	![QQ频道机器人模块流程图](https://fanyohong-blog.oss-cn-shenzhen.aliyuncs.com/image/QQ%E9%A2%91%E9%81%93%E6%9C%BA%E5%99%A8%E4%BA%BA%E6%A8%A1%E5%9D%97%E6%B5%81%E7%A8%8B%E5%9B%BE.png)



## 设计 && 使用文档 

[频道机器人方案设计文档](https://github.com/Fan-Yu-Feng/qq-channel-llm-question-bot/blob/master/doc/%E8%B6%A3%E5%91%B3%E9%97%AE%E7%AD%94%E9%A2%91%E9%81%93%E6%9C%BA%E5%99%A8%E4%BA%BA%E6%96%B9%E6%A1%88%E8%AE%BE%E8%AE%A1%E6%96%87%E6%A1%A3.md)

[频道机器人使用文档](https://github.com/Fan-Yu-Feng/qq-channel-llm-question-bot/blob/master/doc/%E8%B6%A3%E5%91%B3%E9%97%AE%E7%AD%94%E9%A2%91%E9%81%93%E6%9C%BA%E5%99%A8%E4%BA%BA%E4%BD%BF%E7%94%A8%E8%AF%B4%E6%98%8E%E6%96%87%E6%A1%A3.md)

## feature

思考优化：
- [ ] 答案可以通过可按键按钮的的方式交互，提供给用户更好的反馈，而不需要通过指令去输入。
- [ ] 机器人的题目类型扩展性很大，不止我提供的枚举，因为我用了 llm 的prompt 模板，理论上任何类型的题目都可以咨询。
- [ ] 可以接入更多的模型，在配置文件指定使用的模型，或者多模型一起使用，在单个模型挂掉的时候，使用其他模型，提供更好的容错性。
- [ ] 多人互动可以提供答题排行榜，以答题数量和答题正确率算法排名。



## 运行截图

![image-20231117203454855](https://fanyohong-blog.oss-cn-shenzhen.aliyuncs.com/image/image-20231117203454855.png)





## 参考使用

- [langchainGo](https://tmc.github.io/langchaingo/docs/)
- [开发说明 | QQ机器人文档](https://bot.q.qq.com/wiki/develop/api/)
- [tencent-connect/botgo: QQ频道机器人 GOSDK (github.com)](https://github.com/tencent-connect/botgo)

