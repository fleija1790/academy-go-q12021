package model

type Joke struct {
	ID        int    `json:id`
	Setup     string `json:setup`
	Punchline string `json:punchline`
}

type ApiError struct {
	Error   string `json:error`
	Message string `json:message`
}

type UpdateOk struct {
	Message string `json:message`
}
