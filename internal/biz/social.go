package biz

import (
    "context"

    "github.com/go-kratos/kratos/v2/log"
)

type ArticleRepo interface {
}

type CommentRepo interface {
}

type TagRepo interface {
}

type SocialUsecase struct {
    ac ArticleRepo
    cr CommentRepo
    tr TagRepo

    log *log.Helper
}

func NewSocialUsecase(ac ArticleRepo, cr CommentRepo, tr TagRepo, logger log.Logger) *SocialUsecase {
    return &SocialUsecase{ac: ac, cr: cr, tr: tr, log: log.NewHelper(logger)}
}

func (uc *SocialUsecase) CreatArticle(ctx context.Context) error {
    return nil
}
