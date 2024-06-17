package models

import "github.com/golang-jwt/jwt/v5"

type AClaim struct {
	ID           string `json:"user_id"`
	Email        string `json:"email"`
	ProfilePhoto string `json:"profile_photo"`
	jwt.RegisteredClaims
}

type RClaim struct {
	ID string `json:"user_id"`
	jwt.RegisteredClaims
}
