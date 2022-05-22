package message

type NewSMS struct {
	Messages []Message `json:"messages,omitempty"`
}

type Message struct {
	Content string `json:"content,omitempty"`
}
