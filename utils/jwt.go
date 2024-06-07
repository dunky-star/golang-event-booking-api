package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "eyJhbGciOiJIUzI1NiJ9.eyJSb2xlIjoiQWRtaW4iLCJJc3N1ZXIiOiJkdW5reSIsIlVzZXJuYW1lIjoiR29sYW5nQ29kaW5nRXZlbnQiLCJleHAiOjE3MTc3MjI1NjEsImlhdCI6MTcxNzcyMjU2MX0.NtIypCdQNWHF9Wu9dZWptKClCMzCMt9N1vHLp-WcNv0"

func GenerateToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString([]byte(secretKey))
}
