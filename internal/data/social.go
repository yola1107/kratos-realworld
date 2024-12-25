package data

import (
    "kratos-realworld/internal/biz"
)

type articleRepo struct {
    data *Data
}

type commentRepo struct {
    data *Data
}

type tagRepo struct {
    data *Data
}

func NewArticleRepo(data *Data, ) biz.ArticleRepo {
    return &articleRepo{data: data}
}

func NewCommentRepo(data *Data) biz.CommentRepo {
    return &commentRepo{data: data}
}

func NewTagRepo(data *Data) biz.TagRepo {
    return &tagRepo{data: data}
}
