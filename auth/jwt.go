package auth

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/mnuddindev/betterkeep/db"
	"github.com/mnuddindev/betterkeep/models"
	"github.com/mnuddindev/betterkeep/utils"
)

func GenerateAT(user models.Users) (string, error) {
	et := time.Now().Add(1 * time.Minute)
	claims := &models.AClaim{
		ID:           user.ID.String(),
		Email:        user.Email,
		ProfilePhoto: user.ProfilePhoto,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    user.ID.String(),
			ExpiresAt: jwt.NewNumericDate(et),
		},
	}
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	access_token, err := at.SignedString([]byte(utils.Config("APP_SECRET")))
	if err != nil {
		return "", err
	}
	return "Bearer " + access_token, err
}

func GenerateRT(user models.Users) (string, error) {
	et := time.Now().Add(7 * 24 * time.Minute)
	claims := &models.RClaim{
		ID: user.ID.String(),
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    user.ID.String(),
			ExpiresAt: jwt.NewNumericDate(et),
		},
	}
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	refresh_token, err := rt.SignedString([]byte(utils.Config("APP_SECRET")))
	if err != nil {
		return "", err
	}
	return "Bearer " + refresh_token, err
}

func GenerateTokens(user models.Users) (access_token, refresh_token string, err error) {
	access_token, err = GenerateAT(user)
	if err != nil {
		return "", "", err
	}
	refresh_token, err = GenerateRT(user)
	if err != nil {
		return "", "", err
	}
	return access_token, refresh_token, nil
}

func VerifyToken(tr string) (*jwt.Token, error) {
	token, err := jwt.Parse(tr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("secret")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, err
}

func ExtractToken(token string) (jwt.MapClaims, error) {
	tr, err := VerifyToken(token)
	if err != nil {
		return nil, err
	}
	if claims, ok := tr.Claims.(jwt.MapClaims); ok && tr.Valid {
		return claims, nil
	}
	return nil, err
}

func ExtractTokenAuth(token string) (string, error) {
	tr, err := VerifyToken(token)
	if err != nil {
		return "", err
	}
	if claims, ok := tr.Claims.(jwt.MapClaims); ok && tr.Valid {
		userid, err := uuid.Parse(claims["user_id"].(string))
		if err != nil {
			return "", err
		}
		user, err := db.UserById(userid)
		if err != nil {
			return "", err
		}
		newAT, err := GenerateAT(user)
		if err != nil {
			return "", err
		}
		return newAT, nil
	}
	return "", err
}
