package dict

// AnswerType 是趣味问答机器人问题类型的字符串
type AnswerType string

const (
	Answer AnswerType = "答案"
)

// GetAnswerType 根据中文字符串获取对应的枚举值
func GetAnswerType(chineseType string) AnswerType {
	switch chineseType {
	case "答案":
		return Answer

	default:
		return "-1"
	}
}
