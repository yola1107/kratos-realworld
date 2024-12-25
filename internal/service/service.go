package service

import (
    pb "kratos-realworld/api/realworld/v1"
    "kratos-realworld/internal/biz"

    "github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewRealWorldService)

type RealWorldService struct {
    pb.UnimplementedRealWorldServer
    sc *biz.SocialUsecase
    uc *biz.UserUsecase
}

func NewRealWorldService(sc *biz.SocialUsecase, uc *biz.UserUsecase) *RealWorldService {
    return &RealWorldService{sc: sc, uc: uc}
}
