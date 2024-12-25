package service

import (
    "context"

    pb "kratos-realworld/api/realworld/v1"
)

func (s *RealWorldService) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
    return &pb.HelloReply{Message: "Hello " + req.Name}, nil
}

func (s *RealWorldService) Login(ctx context.Context, req *pb.Authentication) (*pb.AuthenticationReply, error) {
    return s.uc.Login(ctx, req)
}

func (s *RealWorldService) Registration(ctx context.Context, req *pb.RegistrationRequest) (*pb.RegistrationReply, error) {
    return s.uc.Register(ctx, req)
}

func (s *RealWorldService) GetCurrentUser(ctx context.Context, req *pb.GetCurrentUserRequest) (*pb.GetCurrentUserReply, error) {
    return &pb.GetCurrentUserReply{}, nil
}

//func (s *RealWorldService) RefreshToken(ctx context.Context, in *pb.RefreshTokenRequest) (*pb.RefreshTokenReply, error) {
//    return s.uc.UpdateToken(ctx, in)
//}

func (s *RealWorldService) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserReply, error) {
    return &pb.UpdateUserReply{}, nil
}

func (s *RealWorldService) GetProfile(ctx context.Context, req *pb.GetProfileRequest) (*pb.GetProfileReply, error) {
    return &pb.GetProfileReply{}, nil
}

func (s *RealWorldService) FollowUser(ctx context.Context, req *pb.FollowUserRequest) (*pb.FollowUserReply, error) {
    return &pb.FollowUserReply{}, nil
}

func (s *RealWorldService) UnfollowUser(ctx context.Context, req *pb.UnfollowUserRequest) (*pb.UnfollowUserReply, error) {
    return &pb.UnfollowUserReply{}, nil
}

func (s *RealWorldService) ListArticles(ctx context.Context, req *pb.ListArticlesRequest) (*pb.ListArticlesReply, error) {
    return &pb.ListArticlesReply{}, nil
}

func (s *RealWorldService) FeedArticles(ctx context.Context, req *pb.FeedArticlesRequest) (*pb.FeedArticlesReply, error) {
    return &pb.FeedArticlesReply{}, nil
}

func (s *RealWorldService) GetArticle(ctx context.Context, req *pb.GetArticleRequest) (*pb.GetArticleReply, error) {
    return &pb.GetArticleReply{}, nil
}

func (s *RealWorldService) CreateArticle(ctx context.Context, req *pb.CreateArticleRequest) (*pb.CreateArticleReply, error) {
    return &pb.CreateArticleReply{}, nil
}

func (s *RealWorldService) UpdateArticle(ctx context.Context, req *pb.UpdateArticleRequest) (*pb.UpdateArticleReply, error) {
    return &pb.UpdateArticleReply{}, nil
}

func (s *RealWorldService) DeleteArticle(ctx context.Context, req *pb.DeleteArticleRequest) (*pb.DeleteArticleReply, error) {
    return &pb.DeleteArticleReply{}, nil
}

func (s *RealWorldService) AddComments(ctx context.Context, req *pb.AddCommentsRequest) (*pb.AddCommentsReply, error) {
    return &pb.AddCommentsReply{}, nil
}

func (s *RealWorldService) GetComments(ctx context.Context, req *pb.GetCommentsRequest) (*pb.GetCommentsReply, error) {
    return &pb.GetCommentsReply{}, nil
}

func (s *RealWorldService) DeleteComment(ctx context.Context, req *pb.DeleteCommentRequest) (*pb.DeleteCommentReply, error) {
    return &pb.DeleteCommentReply{}, nil
}

func (s *RealWorldService) FavoriteArticle(ctx context.Context, req *pb.FavoriteArticleRequest) (*pb.FavoriteArticleReply, error) {
    return &pb.FavoriteArticleReply{}, nil
}

func (s *RealWorldService) UnFavoriteArticle(ctx context.Context, req *pb.UnFavoriteArticleRequest) (*pb.UnFavoriteArticleReply, error) {
    return &pb.UnFavoriteArticleReply{}, nil
}

func (s *RealWorldService) GetTags(ctx context.Context, req *pb.GetTagsRequest) (*pb.GetTagsReply, error) {
    return &pb.GetTagsReply{}, nil
}
