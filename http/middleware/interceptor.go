package middleware

import (
	"otter-v2/http/jwt"
	"otter-v2/http/response"

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
		response.TokenVerifyError(context.Ctx, "", err)
		return
	}

	run(webInput)
}
