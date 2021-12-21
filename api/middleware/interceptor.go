package middleware

import (
	"otter-v2/api/http/jwt"
	"otter-v2/api/http/response"

	"github.com/EricChiou/httprouter"
)

type WebInput struct {
	Context *httprouter.Context
	Payload jwt.Payload
}

func Set(context *httprouter.Context, needToken bool, run func(WebInput)) {
	webInput := WebInput{
		Context: context,
	}

	// check token
	payload, err := verifyToken(context.Ctx)
	webInput.Payload = payload
	if needToken && err != nil {
		response.Error(context.Ctx, response.TokenError, "", err)
		return
	}

	run(webInput)
}
