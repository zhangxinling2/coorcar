package auth

import (
	"context"
	authpb "coolcar/auth/api/gen/v1"
	"coolcar/auth/dao"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service struct {
	Logger *zap.Logger
	authpb.UnimplementedAuthServiceServer
	OpenIdResolve OpenIdResolve
	Mongo         *dao.Mongo
}
type OpenIdResolve interface {
	Resolve(code string) (string, error)
}

func (s *Service) Login(ctx context.Context, req *authpb.LoginRequest) (*authpb.LoginResponse, error) {
	s.Logger.Info("received code ", zap.String("code", req.Code))
	openId, err := s.OpenIdResolve.Resolve(req.Code)
	if err != nil {
		return nil, status.Errorf(codes.Unavailable, "cannot resolve openid: %v", err)
	}
	accountId, err := s.Mongo.ResolveAccountID(ctx, openId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot resolve accountid: %v", err)
	}
	return &authpb.LoginResponse{
		AccessToken: "token for accountId " + accountId,
		ExpireIn:    7200,
	}, nil
}
