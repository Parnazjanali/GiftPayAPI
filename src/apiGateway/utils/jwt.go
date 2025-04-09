package utils

import (
	"apiGateway/internal/model"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("a46f28e07a1655d13f49c579da7fe605af85eec3c5c1e4b4edf0686565c9884751d68bde70993c9dedbd9ebd563f4108e1f2f5089bbfd0ede396fe7c104f02e9e9192338ee5a641b33b5427a9a0deb8d8c753ea50c604e34bb99f4f19416a26f4904be0790a5ef3d4af39498b9872c80a8d925f04766582cd4e557c18324c21455409b22924c0c002abb0e29a0007e2428776480b7de02c7826a3056b8879058a0618e2daea10760f20c2887deefb3fdaf967861a96977511159f32bcb54ff648ae97f2ad3800977a7d68d14531b418f6968985b31107cc93b60d4aade30d0cd5271010bc82c648b70ade546b34c3c018d3005d86c45129af902f1e7df8520ed")

const tokenExpirationTime = time.Second * 3600

func GenerateToken(user model.User) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"exp":      time.Now().Add(tokenExpirationTime).Unix(),
	})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func VerifyToken(tokenstring string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenstring, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Invalid Signing Method")
		}
		return secretKey, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("Invalid Token")
}
