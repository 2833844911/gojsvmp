package http

import (
	"myvmp/object"
	"myvmp/parsejson"
)

type Response struct {
	IsError    bool
	ErrMessage string
	Status     string
	Text       string
	ReHeaders  string
	Content    []byte
}

func New_response() *Response {
	dt := &Response{IsError: true}
	return dt
}

func (rp *Response) ToJSON() object.Object {
	dt := parsejson.ParseStrToJson(rp.Text)

	return dt
}
