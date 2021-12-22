package response

import (
	"encoding/json"
	"fmt"

	"github.com/valyala/fasthttp"
)

type Status string

const (
	OK           Status = "ok"
	TokenExpired Status = "tokenExpired"
	Error        Status = "error"
)

type response struct {
	Status   Status      `json:"status"`
	Result   interface{} `json:"result,omitempty"`
	ErrorMsg string      `json:"errorMsg,omitempty"`
	Trace    string      `json:"trace,omitempty"`
}

type record struct {
	Page  int           `json:"page"`
	Limit int           `json:"limit"`
	Total int           `json:"total"`
	List  []interface{} `json:"list"`
}

func Success(ctx *fasthttp.RequestCtx, result interface{}) {
	resp := response{Status: OK, Result: result}
	send(ctx, resp)
}

func Page(ctx *fasthttp.RequestCtx, page, limit, total int, list []interface{}) {
	resp := response{Status: OK, Result: record{Page: page, Limit: limit, Total: total, List: list}}
	send(ctx, resp)
}

func FormatError(ctx *fasthttp.RequestCtx, errorMsg string, err error) {
	ctx.Response.SetStatusCode(400)
	resp := response{Status: Error, ErrorMsg: errorMsg, Trace: err.Error()}
	send(ctx, resp)
}

func TokenExpiredError(ctx *fasthttp.RequestCtx, errorMsg string, err error) {
	ctx.Response.SetStatusCode(401)
	resp := response{Status: TokenExpired, ErrorMsg: errorMsg, Trace: err.Error()}
	send(ctx, resp)
}

func PermissionError(ctx *fasthttp.RequestCtx, errorMsg string, err error) {
	ctx.Response.SetStatusCode(401)
	resp := response{Status: Error, ErrorMsg: errorMsg, Trace: err.Error()}
	send(ctx, resp)
}

func OperationError(ctx *fasthttp.RequestCtx, errorMsg string, err error) {
	ctx.Response.SetStatusCode(403)
	resp := response{Status: Error, ErrorMsg: errorMsg, Trace: err.Error()}
	send(ctx, resp)
}

func ServerError(ctx *fasthttp.RequestCtx, errorMsg string, err error) {
	ctx.Response.SetStatusCode(500)
	resp := response{Status: Error, ErrorMsg: errorMsg, Trace: err.Error()}
	send(ctx, resp)
}

func send(ctx *fasthttp.RequestCtx, resp response) {
	bytes, _ := json.Marshal(resp)
	fmt.Fprint(ctx, string(bytes))
}
