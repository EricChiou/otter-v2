package service

import (
	"encoding/json"
	"fmt"

	"github.com/valyala/fasthttp"
)

var Response = response{}

type Status string

const (
	OK               Status = "ok"
	AccInactive      Status = "accInactive"
	Duplicate        Status = "duplicate"
	TokenError       Status = "tokenError"
	PermissionDenied Status = "permissionDenied"
	FormatError      Status = "formatError"
	ServerError      Status = "serverError"
	Unknown          Status = "unknown"
)

type response struct {
	Status   Status      `json:"status"`
	Result   interface{} `json:"result,omitempty"`
	ErrorMsg string      `json:"errorMsg,omitempty"`
	Trace    string      `json:"trace,omitempty"`
}

type record struct {
	Page  uint          `json:"page"`
	Limit uint          `json:"limit"`
	Total uint          `json:"total"`
	List  []interface{} `json:"list"`
}

func (r response) OK(ctx *fasthttp.RequestCtx, result interface{}) {
	resp := response{Status: OK, Result: result}
	r.send(ctx, resp)
}

func (r response) Page(ctx *fasthttp.RequestCtx, page, limit, total uint, list []interface{}) {
	resp := response{Status: OK, Result: record{Page: page, Limit: limit, Total: total, List: list}}
	r.send(ctx, resp)
}

func (r response) Error(ctx *fasthttp.RequestCtx, errorMsg string, err error) {
	resp := response{Status: OK, ErrorMsg: errorMsg, Trace: err.Error()}
	r.send(ctx, resp)
}

func (r response) send(ctx *fasthttp.RequestCtx, resp response) {
	bytes, _ := json.Marshal(resp)
	fmt.Fprint(ctx, string(bytes))
}
