package server

import (
    "context"

    v1 "kratos-realworld/api/realworld/v1"
    "kratos-realworld/internal/conf"
    "kratos-realworld/internal/service"

    "github.com/go-kratos/kratos/v2/log"
    "github.com/go-kratos/kratos/v2/middleware/auth/jwt"
    "github.com/go-kratos/kratos/v2/middleware/recovery"
    "github.com/go-kratos/kratos/v2/middleware/selector"
    "github.com/go-kratos/kratos/v2/transport/http"
    jwtv5 "github.com/golang-jwt/jwt/v5"
    "github.com/gorilla/handlers"
)

func NewWhiteListMatcher() selector.MatchFunc {
    whiteList := make(map[string]struct{})
    whiteList["/realworld.v1.RealWorld/Login"] = struct{}{}
    whiteList["/realworld.v1.RealWorld/Registration"] = struct{}{}
    return func(ctx context.Context, operation string) bool {
        if _, ok := whiteList[operation]; ok {
            return false
        }
        return true
    }
}

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, RealWorld *service.RealWorldService, logger log.Logger) *http.Server {
    var opts = []http.ServerOption{
        http.Middleware(
            recovery.Recovery(),
            //middleware.JWTAuthMiddleware(),
            //selector.Server(auth.JWTAuth(jwtConf.Token)).Match(NewWhiteListMatcher()).Build(),
            //selector.Server(middleware.JWTAuthMiddleware()).Match(NewWhiteListMatcher()).Build(),

            //jwt.Server(func(token *jwtv4.Token) (interface{}, error) {
            //    return []byte("secret"), nil
            //}),

            selector.Server(jwt.Server(func(token *jwtv5.Token) (interface{}, error) {
                return []byte("secret"), nil
            })).Match(NewWhiteListMatcher()).Build(),
        ),
        http.Filter(handlers.CORS(
            handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
            handlers.AllowedOrigins([]string{"*"}),
            handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
        )),
    }
    if c.Http.Network != "" {
        opts = append(opts, http.Network(c.Http.Network))
    }
    if c.Http.Addr != "" {
        opts = append(opts, http.Address(c.Http.Addr))
    }
    if c.Http.Timeout != nil {
        opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
    }

    srv := http.NewServer(opts...)
    //srv.Handle("/swagger/*any", httpSwagger.WrapHandler(swaggerFiles.Handler)) // 注册 Swagger 文档
    v1.RegisterRealWorldHTTPServer(srv, RealWorld)
    return srv
}
