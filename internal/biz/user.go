package biz

import (
    "context"

    pb "kratos-realworld/api/realworld/v1"
    "kratos-realworld/pkg/middleware"
    "kratos-realworld/pkg/passwords"

    "github.com/go-kratos/kratos/v2/errors"
    "github.com/go-kratos/kratos/v2/log"
    "gorm.io/gorm"
)

type User struct {
    gorm.Model
    Email        string
    Token        string
    Username     string
    Bio          string
    Image        string
    PasswordHash string
}

type UserRepo interface {
    CreateUser(ctx context.Context, name, email, password string) (*User, error)
    GetUserByEmail(ctx context.Context, email string) (*User, error)
    UpdateToken(ctx context.Context, email, token string) error
}

type ProfileRepo interface {
}

type UserUsecase struct {
    ur UserRepo
    pr ProfileRepo

    log *log.Helper
}

func NewUserUsecase(ur UserRepo, pr ProfileRepo, logger log.Logger) *UserUsecase {
    return &UserUsecase{
        ur: ur,
        pr: pr,

        log: log.NewHelper(logger),
    }
}

func (uc *UserUsecase) Register(ctx context.Context, in *pb.RegistrationRequest) (*pb.RegistrationReply, error) {
    if in.User == nil || in.User.Email == "" || in.User.Password == "" {
        return nil, errors.New(int(pb.ErrorReason_VALIDATIONS), "user", "email or password is empty")
    }

    u, err := uc.ur.CreateUser(ctx, in.User.Username, in.User.Email, in.User.Password)
    if err != nil {
        return nil, err
    }
    return &pb.RegistrationReply{User: &pb.User{
        Email:    u.Email,
        Token:    u.Token,
        Username: u.Username,
        Bio:      u.Bio,
        Image:    u.Image,
    }}, nil
}

func (uc *UserUsecase) Login(ctx context.Context, in *pb.Authentication) (*pb.AuthenticationReply, error) {
    if in.User == nil || in.User.Email == "" || in.User.Password == "" {
        return nil, errors.New(int(pb.ErrorReason_VALIDATIONS), "user", "email or password is empty")
    }

    u, err := uc.ur.GetUserByEmail(ctx, in.User.Email)
    if err != nil {
        return nil, err
    }
    if !passwords.VerifyPassword(u.PasswordHash, in.User.Password) {
        return nil, errors.Newf(int(pb.ErrorReason_VALIDATIONS), "user", "password invalid")
    }

    // 判断token是否要刷新
    if middleware.IsTokenExpired(u.Token) {
        newToken, err := middleware.GenerateToken(u.Email)
        if err != nil {
            return nil, err
        }
        if err = uc.ur.UpdateToken(ctx, in.User.Email, newToken); err != nil {
            return nil, err
        }
        u.Token = newToken
        uc.log.Warnf("u:(%+v) refresh token. %+v", u.Email, u.Token)
    }

    return &pb.AuthenticationReply{User: &pb.User{
        Email:    u.Email,
        Token:    u.Token,
        Username: u.Username,
        Bio:      u.Bio,
        Image:    u.Image,
    }}, nil
}
