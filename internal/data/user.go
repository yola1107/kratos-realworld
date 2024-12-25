package data

import (
    "context"
    "fmt"

    pb "kratos-realworld/api/realworld/v1"
    "kratos-realworld/internal/biz"
    "kratos-realworld/pkg/middleware"
    "kratos-realworld/pkg/passwords"

    "github.com/go-kratos/kratos/v2/errors"
    "github.com/go-kratos/kratos/v2/log"
)

type userRepo struct {
    data *Data
    log  *log.Helper
}

type profileRepo struct {
    data *Data
}

func NewUserRepo(data *Data) biz.UserRepo {
    return &userRepo{data: data}
}

func NewProfileRepo(data *Data) biz.ProfileRepo {
    return &commentRepo{data: data}
}

func (r *userRepo) CreateUser(ctx context.Context, name, email, password string) (*biz.User, error) {
    if r.data.db.Where("email = ?", email).First(&biz.User{}).RowsAffected >= 1 {
        return nil, errors.New(int(pb.ErrorReason_VALIDATIONS), "user", fmt.Sprintf("email(%+v) registered", email))
    }
    token, err := middleware.GenerateToken(email)
    if err != nil {
        return nil, err
    }
    u := &biz.User{
        Email:        email,
        Token:        token,
        Username:     name,
        Bio:          "",
        Image:        "",
        PasswordHash: passwords.HashPassword(password),
    }
    if r.data.db.Create(u).Error != nil {
        return nil, errors.New(int(pb.ErrorReason_VALIDATIONS), "user", fmt.Sprintf("create err by email(%+v)", email))
    }
    return u, nil
}

func (r *userRepo) GetUserByEmail(ctx context.Context, email string) (*biz.User, error) {
    u := &biz.User{}
    if err := r.data.db.Where("email = ?", email).First(u).Error; err != nil {
        return nil, errors.New(int(pb.ErrorReason_VALIDATIONS), "user", "not find")
    }
    return u, nil
}

func (r *userRepo) UpdateToken(ctx context.Context, email, newToken string) error {
    result := r.data.db.Model(&biz.User{}).
        Where("email = ?", email). // Specify the condition for the update
        Update("token", newToken) // Update the token field
    return result.Error           // Handle the error accordingly
}
