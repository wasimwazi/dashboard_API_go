package user

import "github.com/dgrijalva/jwt-go"

// User : User Struct
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}

// Login : User login struct
type Login struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// JWT : JWT token struct
type JWT struct {
	ID          int    `json:"id,omitempty"`
	AccessToken string `json:"access_token"`
}

// Claims : Details required to identify an Admin
type Claims struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	jwt.StandardClaims
}
