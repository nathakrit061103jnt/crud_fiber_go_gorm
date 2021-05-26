package configs

import (
	jwtware "github.com/gofiber/jwt/v2"
)

var ConfigAuth = jwtware.Config{
	SigningKey: []byte("secret"),
}
