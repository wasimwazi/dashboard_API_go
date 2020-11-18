package user

import (
	"dashboard-api/utils"
	"errors"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// GenerateToken : to Generate JWT Access Token
func GenerateToken(id int, username string) (*JWT, error) {
	token, err := getAccessToken(id, username)
	if err != nil {
		return nil, errors.New(utils.JWTError)
	}
	var jwtToken JWT
	jwtToken.ID = id
	jwtToken.AccessToken = token
	return &jwtToken, nil
}

func getAccessToken(id int, username string) (string, error) {
	var jwtKey = []byte(utils.GetEnvKey("JWTKEY"))
	expirationTime := time.Now().Add(utils.AccessTokenInterval * time.Minute)
	claims := &Claims{
		ID:       id,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(jwtKey)
}

// VerifyToken : to Verify JWT Token
func VerifyToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bearerToken := r.Header.Get("Authorization")
		authHeader := strings.Split(bearerToken, " ")
		if len(authHeader) != 2 {
			utils.Fail(w, utils.BadRequest, utils.InvalidTokenError)
			return
		}
		jwtToken := authHeader[1]
		claims := &Claims{}
		token, err := jwt.ParseWithClaims(jwtToken, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(utils.GetEnvKey("JWTKEY")), nil
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				utils.Fail(w, utils.BadRequest, utils.InvalidTokenError)
				return
			}
			utils.Fail(w, utils.BadRequest, utils.InvalidTokenError)
			return
		}
		if !token.Valid {
			utils.Fail(w, utils.BadRequest, utils.InvalidTokenError)
			return
		}
		r.Header.Add("id", strconv.Itoa(claims.ID))
		r.Header.Add("username", claims.Username)
		next.ServeHTTP(w, r)
	})
}
