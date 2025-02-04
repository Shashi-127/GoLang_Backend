package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string,error){
	bytes,err:=bcrypt.GenerateFromPassword([]byte(password),14) // generates hash of length 14 from password
	return string(bytes),err
}

func CompareHashedpassword(hashedPass, password string) bool{
	err:=bcrypt.CompareHashAndPassword([]byte(hashedPass),[]byte(password))
	return err==nil
}