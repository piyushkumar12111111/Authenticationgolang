package models

import "golang.org/x/crypto/bcrypt"

type User struct {
    Username string
    Password []byte //! hashed password
}

//! mock database
var Users = make(map[string]*User)


func CreateUser(username, password string) error {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return err
    }
    Users[username] = &User{Username: username, Password: hashedPassword}
    return nil
}


func AuthenticateUser(username, password string) bool {
    user, exists := Users[username]
    if !exists {
        return false
    }
    return bcrypt.CompareHashAndPassword(user.Password, []byte(password)) == nil
}
