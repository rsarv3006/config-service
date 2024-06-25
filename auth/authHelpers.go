package auth

import (
	"RjsConfigService/ent"
	"errors"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

type JWTClaims struct {
	User *ent.User
	jwt.StandardClaims
}

const FiftyYears = 50 * 365 * 24 * time.Hour

var (
	ErrExpired = errors.New("token expired")
	ErrInvalid = errors.New("couldn't parse claims")
)

func GenerateJWT(user *ent.User, ctx *fiber.Ctx) (string, error) {
	jwtSecretString := ctx.Locals("JwtSecret").(string)
	return generateToken(user, jwtSecretString)
}

func GenerateJWTFromSecret(user *ent.User, jwtSecretString string) (string, error) {
	return generateToken(user, jwtSecretString)
}

func generateToken(user *ent.User, jwtSecretString string) (string, error) {
	jwtKey := []byte(jwtSecretString)

	expirationTime := time.Now().Add(FiftyYears)

	claims := &JWTClaims{
		User: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	return tokenString, err
}

func ValidateToken(signedToken string, ctx *fiber.Ctx) (*ent.User, error) {
	jwtSecretString := ctx.Locals("JwtSecret").(string)
	jwtKey := []byte(jwtSecretString)

	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		},
	)

	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*JWTClaims)
	if !ok {
		return nil, ErrInvalid
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		return nil, ErrExpired
	}
	return claims.User, nil
}
