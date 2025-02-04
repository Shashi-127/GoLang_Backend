package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)
const secretKey="supersecret"
func GenerateToken(email string, id int64) (string,error){
	token:=jwt.NewWithClaims(jwt.SigningMethodHS256,jwt.MapClaims{
		"email":email,
		"userId":id,
		"exp":time.Now().Add(time.Hour*2).Unix(),
	})
	return token.SignedString([]byte(secretKey))
}

func VerifyToken(token string) (int64,error){
	parsedToken,err:=jwt.Parse(token,func (token *jwt.Token) (interface{},error) {
        _,ok:=token.Method.(*jwt.SigningMethodHMAC)

		if!ok{
			return nil,errors.New("Unexpected signing Method")
		}
		return []byte(secretKey),nil 
	})
	 if err!=nil{
		return 0,errors.New("Couldn't parse Token")
	 }
	 tokenIsValid:=parsedToken.Valid
	 if !tokenIsValid{
		return 0,errors.New("Invalid Token")
	 }

	 claims,ok:=parsedToken.Claims.(jwt.MapClaims)
	 if !ok{
		return 0,errors.New("Invalid token claims")
	 }
	 userId:=int64(claims["userId"].(float64))
	 return userId,nil 
}

// JWT (JSON Web Token) is a compact and self-contained way of securely transmitting information between parties as a JSON object. It is commonly used for authentication and authorization in web applications and APIs.

// A JWT consists of three parts, separated by dots(.)
// Header(metadata about token ): Specifies the type of token (JWT) and the signing algorithm (e.g., HS256, RS256).
// payload :Contains the actual data or claims (e.g., user ID, roles, expiration time).
// signature:Created by signing the header and payload with a secret key.