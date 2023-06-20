package entity

type Tool struct {
	Id   int    `json:"id"`
	Name string `json:"tool_name"`
	Icon []byte `json:"icon"`
}
