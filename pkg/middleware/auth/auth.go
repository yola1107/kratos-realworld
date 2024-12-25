package auth

import (
    "context"
    "errors"
    "fmt"
    "strings"
    "time"

    "github.com/go-kratos/kratos/v2/middleware"
    "github.com/go-kratos/kratos/v2/transport"
    "github.com/golang-jwt/jwt/v4"
)

func GenerateToken(secretKey, username string) string {
    // Create a new token object, specifying signing method and the claims
    // you would like it to contain.
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "username": username,
        "nbf":      time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
    })

    // Sign and get the complete encoded token as a string using the secret
    tokenString, err := token.SignedString([]byte(secretKey))
    if err != nil {
        return err.Error()
    }
    return tokenString
}

func JWTAuth(secretKey string) middleware.Middleware {
    return func(handler middleware.Handler) middleware.Handler {
        return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
            if tr, ok := transport.FromServerContext(ctx); ok {
                tokenString := tr.RequestHeader().Get("Authorization")

                //Token XXXXXXXXXXXXXXXXX
                auths := strings.SplitN(tokenString, " ", 2)
                if len(auths) != 2 || !strings.EqualFold(auths[0], "Token") {
                    return nil, errors.New("jwt Token missing")
                }

                // Parse takes the token string and a function for looking up the key. The latter is especially
                // useful if you use multiple keys for your application.  The standard is to use 'kid' in the
                // head of the token to identify which key to use, but the parsed token (head and claims) is provided
                // to the callback, providing flexibility.
                token, err := jwt.Parse(auths[1], func(token *jwt.Token) (interface{}, error) {
                    // Don't forget to validate the alg is what you expect:
                    if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
                        return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
                    }

                    // hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
                    return []byte(secretKey), nil
                })
                if err != nil {
                    return nil, err
                }

                if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
                    fmt.Println(claims["username"], claims["nbf"])
                } else {
                    return nil, errors.New("InValid Token")
                }

            }
            return handler(ctx, req)
        }
    }
}
