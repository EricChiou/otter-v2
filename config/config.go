package config

import "github.com/EricChiou/config"

type Config struct {
	ServerName      string `key:"SERVER_NAME"`
	ServerPort      string `key:"SERVER_PORT"`
	SSLCertFilePath string `key:"SSL_CERT_FILE_PATH"`
	SSLKeyFilePath  string `key:"SSL_KEY_FILE_PATH"`
	JWTKey          string `key:"JWT_KEY"`
	JWTExpire       int    `key:"JWT_EXPIRE"`
	ENV             string `key:"ENV"`
}

var ServerName string
var ServerPort string
var SSLCertFilePath string
var SSLKeyFilePath string
var JWTKey string
var JWTExpire int
var ENV string

func Load(filePath string) error {
	cfg := Config{}
	err := config.Load(filePath, &cfg)
	if err != nil {
		return err
	}

	ServerName = cfg.ServerName
	ServerPort = cfg.ServerPort
	SSLCertFilePath = cfg.SSLCertFilePath
	SSLKeyFilePath = cfg.SSLKeyFilePath
	JWTKey = cfg.JWTKey
	JWTExpire = cfg.JWTExpire
	ENV = cfg.ENV
	return nil
}
