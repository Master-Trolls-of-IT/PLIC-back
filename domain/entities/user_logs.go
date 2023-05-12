package entities

type UserLogs struct {
	Level   string `json:"level"`
	Message string `json:"message"`
	Details string `json:"details"`
	Source  string `json:"source"`
	Date    string `json:"date"`
}
