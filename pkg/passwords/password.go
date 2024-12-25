package passwords

import (
    "golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) string {
    b, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return ""
    }
    return string(b)
}

func VerifyPassword(hashPwd, password string) bool {
    return bcrypt.CompareHashAndPassword([]byte(hashPwd), []byte(password)) == nil
}
