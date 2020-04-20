package chttp

import (
	"encoding/json"
)

type Res struct {
	Code       string              `json:"code"`
	Message    string              `json:"message"`
	Data       interface{}         `json:"data"`
	Ok         bool                `json:"ok"`
	RespStr    string              `json:"resp_str"`
	Time       int                 `json:"time"`
	StatusCode int                 `json:"status_code"`
	Url        string              `json:"url"`
	Proto      string              `json:"proto"`
	Headers    map[string][]string `json:"headers"`
	Req        *Req                `json:"-"`
	Success    bool                `json:"success"`
	Buffer     []byte              `json:"-"`
}

func NewHttpRes() *Res {
	res := Res{Code: "0", Message: "", Data: nil, Ok: false, StatusCode: 599}
	return &res
}
func (res *Res) String() string {
	str, _ := json.Marshal(res)
	return string(str)
}
