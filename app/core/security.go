package security

import (
	"fmt"
	"os"
	"time"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

//работа с паролем
func Password_hashing(password string) (string, error){
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func Password_checking(hashpassword, password string) (bool){
	marker := bcrypt.CompareHashAndPassword([]byte(hashpassword), []byte(password))
	return (marker == nil)
}

//работа с токеном
func GenerateToken(userID uint) (string, error) {
	var secret_key = []byte(os.Getenv("SECRET_KEY"))
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(time.Minute*15).Unix(),
		"iat": time.Now().Unix(),
	})
	token, err := claims.SignedString(secret_key)
	if err != nil {
		return "", err
	}
	return token, nil
}

func VarifyToken(token string) error{
	tok, err := jwt.Parse(token, func(Token *jwt.Token) (interface{}, error){
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if err != nil {
		return err
	}
	if !tok.Valid {
		return fmt.Errorf("invalid token")
	}
	return nil
}	
