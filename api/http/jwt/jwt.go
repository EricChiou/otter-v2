package jwt

import (
	"otter-calendar-ws/config"
	"otter-calendar-ws/pkg/jwt"
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
	payload := Payload{
		ID:       id,
		Acc:      acc,
		Name:     name,
		RoleCode: rolaeCode,
		Exp:      time.Now().Unix() + int64(config.JWTExpire*86400),
	}

	return jwt.Generate(payload, config.JWTKey, alg)
}

func Verify(j, k string) (Payload, error) {
	var payload Payload
	err := jwt.Verify(&payload, j, config.JWTKey, alg)

	return payload, err
}
