package server

type Message struct {
	From    string `json:"from"`
	Type    string `json:"type"`
	Message any    `json:"message"`
}
