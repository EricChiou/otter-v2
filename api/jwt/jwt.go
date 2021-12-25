package jwt

import (
	"otter-v2/config"
	"otter-v2/pkg/jwt"
	"time"
)

var alg = jwt.Alg.HS256

type Payload struct {
	ID       int    `json:"id"`
	Acc      string `json:"acc"`
	Name     string `json:"name"`
	RoleCode string `json:"roleCode"`
	Exp      int64  `json:"exp"`
}

func Generate(id int, acc, name, rolaeCode string) (string, error) {
	cfg := config.Get()

	payload := Payload{
		ID:       id,
		Acc:      acc,
		Name:     name,
		RoleCode: rolaeCode,
		Exp:      time.Now().Unix() + int64(cfg.JWTExpire*86400),
	}

	return jwt.Generate(payload, cfg.JWTKey, alg)
}

func Verify(j, k string) (Payload, error) {
	var payload Payload
	err := jwt.Verify(&payload, j, config.Get().JWTKey, alg)

	return payload, err
}
