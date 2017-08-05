package hutoma

type hutomaChatResponse struct {
	ChatID    string `json:"chatId"`
	Timestamp int64  `json:"timestamp"`
	Result    struct {
		Score       float64 `json:"score"`
		Query       string  `json:"query"`
		Answer      string  `json:"answer"`
		History     string  `json:"history"`
		ElapsedTime float64 `json:"elapsedTime"`
	} `json:"result"`
	Status struct {
		Code int    `json:"code"`
		Info string `json:"info"`
	} `json:"status"`
}
