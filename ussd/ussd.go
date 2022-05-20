package ussd

type Request struct {
	Sid        string `json:"sid,omitempty"`
	Subscriber string `json:"subscriber,omitempty"`
	Code       string `json:"code,omitempty"`
	Prompt     string `json:"prompt,omitempty"`
}

type Response struct {
	Sid        *string `json:"sid,omitempty"`
	Subscriber *string `json:"subscriber,omitempty"`
	Code       *string `json:"code,omitempty"`
	Content    *string `json:"content,omitempty"`
	Prompt     *Prompt `json:"input,omitempty"`
	Error      Error   `json:"error"`
	Close      *bool   `json:"close,omitempty"`
}

type Prompt struct {
	None     *bool       `json:"none,omitempty"`
	Requires Requirement `json:"requires,omitempty"`
}

type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type Requirement struct {
	Any     *bool    `json:"options,omitempty"`
	Choices []string `json:"choices,omitempty"`
	Length  *Length  `json:"length,omitempty"`
	Pattern *string  `json:"pattern,omitempty"`
}

type Length struct {
	Eq *int `json:"eq,omitempty"`
	Ne *int `json:"ne,omitempty"`
	Lt *int `json:"lt,omitempty"`
	Gt *int `json:"gt,omitempty"`
	Le *int `json:"le,omitempty"`
	Ge *int `json:"ge,omitempty"`
}

type Config struct {
	Header map[string]string
}

type Service struct {
	CallbackUrl   string
	Authorization struct {
		Basic struct {
			Username string
			Password string
		}
	}
	TLS struct {
		Mutual struct {
			Certificate string
		}
	}
}
