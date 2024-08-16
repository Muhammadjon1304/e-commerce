package views

import "github.com/muhammadjon1304/e-commerce/status"

type R struct {
	Status    string      `json:"status"`
	ErrorCode int         `json:"error_code"`
	ErrorNote string      `json:"error_note"`
	Data      interface{} `json:"data"`
}

func View(data interface{}) R {
	return R{
		Status: status.Success,
		Data:   data,
	}
}

func ErrView(code int, note string) R {
	return R{
		Status:    status.Failure,
		ErrorCode: code,
		ErrorNote: note,
		Data:      nil,
	}
}
