package entity

import (
	"be-classroom/app/config"
	"be-classroom/pkg/errs"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id        int
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

var invalidToken = errs.NewUnauthenticatedError("Invalid token")

func (u *User) parseToken(tokenString string) (*jwt.Token, errs.Error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {

		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, invalidToken
		}

		return []byte(config.AppConfig().JwtSecretKey), nil
	})

	if err != nil {
		return nil, invalidToken
	}

	return token, nil
}

func (u *User) bindTokenToUserEntity(claims jwt.MapClaims) errs.Error {

	if name, ok := claims["name"].(string); ok {
		u.Name = name
	} else {
		return invalidToken
	}

	if email, ok := claims["email"].(string); ok {
		u.Email = email
	} else {
		return invalidToken
	}

	return nil
}

func (u *User) ValidateToken(bearerToken string) errs.Error {
	isBearer := strings.HasPrefix(bearerToken, "Bearer ")

	if !isBearer {
		return invalidToken
	}

	spiltToken := strings.Split(bearerToken, " ")

	if len(spiltToken) != 2 {
		return invalidToken
	}

	tokenString := spiltToken[1]
	token, err := u.parseToken(tokenString)

	if err != nil {
		return err
	}

	var mapClaims jwt.MapClaims

	if claims, ok := token.Claims.(jwt.MapClaims); !ok || !token.Valid {
		return invalidToken
	} else {
		mapClaims = claims
	}

	err = u.bindTokenToUserEntity(mapClaims)

	return err
}

func (u *User) tokenClaim() jwt.MapClaims {
	return jwt.MapClaims{
		"name":  u.Name,
		"email": u.Email,
	}
}

func (u *User) signToken(claims jwt.MapClaims) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString([]byte(config.AppConfig().JwtSecretKey))

	return tokenString
}

func (u *User) GenerateToken() string {
	claims := u.tokenClaim()

	return u.signToken(claims)
}

func (u *User) HashPassword() errs.Error {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)

	if err != nil {
		return errs.NewInternalServerError("Something went wrong")
	}

	u.Password = string(hashPassword)

	return nil
}

func (u *User) ComparePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))

	return err == nil
}
