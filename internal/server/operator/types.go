package operator

type NewSMS struct {
	Messages []Message `json:"messages,omitempty"`
}

type Message struct {
	Content string `json:"content,omitempty"`
}

type Operator struct {
	SMS  SMSOperator  `json:"sms"`
	USSD USSDOperator `json:"ussd"`
}

type SMSOperator struct {
	Provider string `json:"provider,omitempty"`
	Config   string `json:"config,omitempty"`
}

type USSDOperator struct {
	Provider string `json:"provider,omitempty"`
	Config   string `json:"config,omitempty"`
}
