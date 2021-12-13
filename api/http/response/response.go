package response

import (
	"encoding/json"
	"fmt"

	"github.com/valyala/fasthttp"
)

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

func Success(ctx *fasthttp.RequestCtx, result interface{}) {
	resp := response{Status: OK, Result: result}
	send(ctx, resp)
}

func Page(ctx *fasthttp.RequestCtx, page, limit, total uint, list []interface{}) {
	resp := response{Status: OK, Result: record{Page: page, Limit: limit, Total: total, List: list}}
	send(ctx, resp)
}

func Error(ctx *fasthttp.RequestCtx, status Status, errorMsg string, err error) {
	resp := response{Status: status, ErrorMsg: errorMsg, Trace: err.Error()}
	send(ctx, resp)
}

func send(ctx *fasthttp.RequestCtx, resp response) {
	bytes, _ := json.Marshal(resp)
	fmt.Fprint(ctx, string(bytes))
}
