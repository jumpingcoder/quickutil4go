package entity

type TestBO struct {
	Uid    int    `json:"uid"`
	Input  string `json:"input"`
	Output string `json:"output"`
}

type ResultVO struct {
	Success bool        `json:"success"`
	Code    int         `json:"code"`
	Message interface{} `json:"message"`
}
