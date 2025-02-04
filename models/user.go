package models

import (
	"errors"
	"fmt"
	"restapi/db"
	"restapi/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u User) Save() error {
	query := "INSERT INTO users(email, password) VALUES (?, ?)"
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()
	hashedPass, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}
	result, err := stmt.Exec(u.Email, hashedPass)

	if err != nil {
		return err
	}

	userId, err := result.LastInsertId()

	u.ID = userId
	return err
}

func GetUserById(id int64) (*User, error) {
	query := "SELECT * FROM users where id = ?"
	row := db.DB.QueryRow(query, id) // QueryRow executes a query that is expected to return at most one row.
	var user User
	err := row.Scan(&user.ID, &user.Email, &user.Password)
	if err != nil {
		fmt.Println("what's error !")
		return nil, err
	}
	return &user, nil
}

// func (u User) ValidateCreddentials() error {
// 	query := `SELECT password FROM users WHERE emial=?`
// 	row := db.DB.QueryRow(query, u.Email)
// 	var retreivePass string
// 	err := row.Scan(&retreivePass)
// 	if err != nil {
// 		return errors.New("Invalid Credentials")
// 	}
// 	passwordIsValid := utils.CompareHashedpassword(retreivePass, u.Password)
// 	if !passwordIsValid {
// 		return errors.New("Invalid Credentials")
// 	}
// 	return nil
// }

func (u *User) ValidateCreddentials() error {
	query := "SELECT  id,password FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, u.Email)

	var retrievedPassword string
	err := row.Scan(&u.ID,&retrievedPassword)

	if err != nil {
		// fmt.Println("can't scan")
		return errors.New("Credentials invalid")
	}
    // fmt.Println(u.Password,retrievedPassword) // u.password is passed by client and retrievepass is hashed 
	passwordIsValid := utils.CompareHashedpassword(retrievedPassword, u.Password)

	if !passwordIsValid {
		// fmt.Println("password not valid")
		return errors.New("Credentials invalid")
	}

	return nil
}
