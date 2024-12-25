package middleware

import (
    "context"
    "strings"
    "time"

    "github.com/go-kratos/kratos/v2/middleware"
    "github.com/go-kratos/kratos/v2/transport"
    "github.com/golang-jwt/jwt/v4"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
)

const (
    _secretKey   = "secret"
    _tokenExpiry = 7 * 24 * time.Hour // Token 默认有效期 24 小时
)

// GenerateToken 生成 JWT Token
func GenerateToken(userID string) (string, error) {
    claims := jwt.MapClaims{
        "userID": userID,
        "iat":    time.Now().Unix(),
        "exp":    time.Now().Add(_tokenExpiry).Unix(), // 设置有效期为 7*24 小时
        "nbf":    time.Now().Unix(),                   // 设置 Token 生效时间为当前时间
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

    // Sign and get the complete encoded token as a string using the secret
    return token.SignedString([]byte(_secretKey))
}

//// RefreshToken 检查是否需要刷新token
//func RefreshToken(expireToken string) (string, error) {
//    // 解析token
//    token, err := parseToken(strings.TrimPrefix(expireToken, "Bearer "))
//    if err != nil {
//        return "", err
//    }
//
//    // 验证 token 是否有效并提取 Claims
//    claims, ok := token.Claims.(jwt.MapClaims)
//    if !ok || !token.Valid {
//        return "", status.Errorf(codes.Unauthenticated, "Invalid Token: %v", err.Error())
//    }
//
//    // 校验 token 的过期时间 (exp)
//    if exp, ok := claims["exp"].(float64); ok {
//        if time.Now().Unix() < int64(exp) {
//            return "", status.Errorf(codes.Unauthenticated, "Token has no expired")
//        }
//    }
//
//    // 如果 Token 过期，重新生成一个新的 Token
//    return GenerateToken(claims["userID"].(string))
//}

func IsTokenExpired(expireToken string) bool {
    // 解析token
    token, err := parseToken(strings.TrimPrefix(expireToken, "Bearer "))
    if err != nil {
        return true
    }

    // 验证 token 是否有效并提取 Claims
    claims, ok := token.Claims.(jwt.MapClaims)
    if !ok || !token.Valid {
        return true
    }

    // 校验 token 的过期时间 (exp)
    if exp, ok := claims["exp"].(float64); ok {
        if time.Now().Unix() >= int64(exp) {
            return true
        }
        return false
    }

    return false
}

// JWTAuthMiddleware JWT 认证中间件
func JWTAuthMiddleware() middleware.Middleware {
    return func(handler middleware.Handler) middleware.Handler {
        return func(ctx context.Context, req interface{}) (reply interface{}, err error) {

            if tr, ok := transport.FromServerContext(ctx); ok {
                // 从请求头中获取 Bearer Token
                authHeader := tr.RequestHeader().Get("Authorization")
                if authHeader == "" {
                    return nil, status.Errorf(codes.Unauthenticated, "Authorization header missing")
                }

                // 获取 token 字符串
                tokenString := strings.TrimPrefix(authHeader, "Bearer ")
                if tokenString == authHeader {
                    return nil, status.Errorf(codes.Unauthenticated, "Authorization header must start with Bearer")
                }

                // 解析并验证 JWT
                token, err := parseToken(tokenString)
                if err != nil {
                    return nil, err
                }

                // 验证 token 是否有效
                if err = isTokenValid(token); err != nil {
                    return nil, err
                }

                // 将用户名和其他信息放到 context 中
                claims, _ := token.Claims.(jwt.MapClaims)
                ctx = context.WithValue(ctx, "userID", claims["userID"])
            }

            // 调用下一个处理函数
            return handler(ctx, req)
        }
    }
}

// parseToken 解析并验证 JWT Token
func parseToken(tokenString string) (*jwt.Token, error) {
    return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, status.Errorf(codes.Unauthenticated, "unexpected signing method: %v", token.Header["alg"])
        }
        return []byte(_secretKey), nil
    })
}

// isTokenValid 检查 Token 是否有效（包括 nbf、exp 等）
func isTokenValid(token *jwt.Token) error {
    claims, ok := token.Claims.(jwt.MapClaims)
    if !ok || !token.Valid {
        return status.Errorf(codes.Unauthenticated, "Invalid Token")
    }

    // 校验 token 的过期时间 (exp)
    if isTokenExpired(claims) {
        return status.Errorf(codes.Unauthenticated, "Token has expired")
    }

    // 校验 token 的生效时间 (nbf)
    if isTokenNotYetValid(claims) {
        return status.Errorf(codes.Unauthenticated, "Token is not yet valid")
    }

    return nil
}

// isTokenExpired 检查 Token 是否过期
func isTokenExpired(claims jwt.MapClaims) bool {
    if exp, ok := claims["exp"].(float64); ok {
        return time.Now().Unix() > int64(exp)
    }
    return false
}

// isTokenNotYetValid 检查 Token 是否还未生效
func isTokenNotYetValid(claims jwt.MapClaims) bool {
    if nbf, ok := claims["nbf"].(float64); ok {
        return time.Now().Unix() < int64(nbf)
    }
    return false
}
