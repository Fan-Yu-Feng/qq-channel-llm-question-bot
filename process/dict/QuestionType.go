package dict

import "fmt"

// QuestionType 是趣味问答机器人问题类型的字符串
type QuestionType string

const (
	BrainTeaser     QuestionType = "脑筋急转弯"
	MathProblem     QuestionType = "数学题"
	HistoryMystery  QuestionType = "历史谜题"
	MovieLiterature QuestionType = "电影与文学"
	Riddles         QuestionType = "谜语与字谜"
	Technology      QuestionType = "技术与科技"
)

// GetQuestionType 根据中文字符串获取对应的枚举值
func GetQuestionType(chineseType string) QuestionType {
	switch chineseType {
	case "脑筋急转弯":
		return BrainTeaser
	case "数学题":
		return MathProblem
	case "历史谜题":
		return HistoryMystery
	case "电影与文学":
		return MovieLiterature
	case "谜语与字谜":
		return Riddles
	case "技术与科技":
		return Technology
	default:
		return "-1"
	}
}

func main() {
	// 示例：使用趣味问答机器人问题类型的字符串
	chineseQuestionType := "数学题"
	questionType := GetQuestionType(chineseQuestionType)

	switch questionType {
	case BrainTeaser:
		fmt.Println("这是一个脑筋急转弯。")
	case MathProblem:
		fmt.Println("这是一个数学题。")
	case HistoryMystery:
		fmt.Println("这是一个历史谜题。")
	case MovieLiterature:
		fmt.Println("这是一个电影与文学问题。")
	case Riddles:
		fmt.Println("这是一个谜语与字谜。")
	case Technology:
		fmt.Println("这是一个技术与科技问题。")
	default:
		fmt.Println("未知问题类型。")
	}
}
