package chatbot

import "errors"

type Talk interface {
	Hello(userName string) string
	Talk(heard string) (saying string, end bool, err error)
}

type Chatbot interface {
	Name() string
	Begin() (string, error)
	Talk
	ReportError(err error) string
	End() (string, error)
}

//定义error实体
var (
	ErrInvalidChatbotName = errors.New("Invalid chatbot name")
	ErrInvalidChatbot     = errors.New("Invalid chatbot")
	ErrExistingChatbot    = errors.New("Existing chatbot")
)

var chatbotMap = map[string]Chatbot{}

func Register(chatbot Chatbot) error {
	if chatbot == nil {
		return ErrInvalidChatbot
	}
	name := chatbot.Name()
	if name == "" {
		return ErrInvalidChatbotName
	}
	if _, ok := chatbotMap[name]; ok {
		return ErrExistingChatbot
	}
	chatbotMap[name] = chatbot
	return nil
}

func Get(name string) Chatbot {
	return chatbotMap[name]
}
