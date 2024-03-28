package models

import "golang.org/x/crypto/bcrypt"

// User represents a user in the system.
type User struct {
    Username string
    Password []byte // hashed password
}

// mock database
var Users = make(map[string]*User)

// CreateUser hashes the password and adds a new user to the mock database.
func CreateUser(username, password string) error {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return err
    }
    Users[username] = &User{Username: username, Password: hashedPassword}
    return nil
}

// AuthenticateUser checks if the user exists and if the password is correct.
func AuthenticateUser(username, password string) bool {
    user, exists := Users[username]
    if !exists {
        return false
    }
    return bcrypt.CompareHashAndPassword(user.Password, []byte(password)) == nil
}
