package service

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"gitlab.com/laboct2021/pnl-game/internal/models"
)

type Jwt struct {
}
type UserClaims struct {
	jwt.StandardClaims
	User models.User
}
type TokenClaims struct {
	jwt.StandardClaims
	Token string
}

func (j Jwt) CreateToken(user models.User) (models.Token, error) {
	var err error
	claims := &UserClaims{
		User: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwt := models.Token{}

	jwt.AccessToken, err = token.SignedString([]byte(os.Getenv("token_password")))
	if err != nil {
		return jwt, err
	}

	return j.createRefreshToken(jwt)
}

func (Jwt) ValidateToken(accessToken string) (models.User, error) {
	claims := &UserClaims{}
	token, err := jwt.ParseWithClaims(accessToken, claims, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("token_password")), nil
	})
	user := models.User{}
	if err != nil {
		return user, err
	}
	if token.Valid {
		user = claims.User //panic
		return user, nil
	}

	return user, errors.New("invalid token")
}

func (Jwt) ValidateRefreshToken(model models.Token) (models.User, error) {
	claimsToken := &TokenClaims{}
	token, err := jwt.ParseWithClaims(model.RefreshToken, claimsToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("token_password")), nil
	})
	user := models.User{}
	if err != nil {
		return user, err
	}
	if !(token.Valid) {
		return user, errors.New("invalid token")
	}

	strToken := claimsToken.Token
	claimsUser := &UserClaims{}
	token, err = jwt.ParseWithClaims(strToken, claimsUser, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("token_password")), nil
	})
	if err != nil {
		return user, err
	}
	user = claimsUser.User

	return user, nil
}

func (Jwt) createRefreshToken(token models.Token) (models.Token, error) {
	var err error
	claims := &TokenClaims{
		Token: token.AccessToken,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(30 * time.Hour * 24).Unix(),
		},
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token.RefreshToken, err = refreshToken.SignedString([]byte(os.Getenv("token_password")))
	if err != nil {
		return token, err
	}

	return token, nil
}
