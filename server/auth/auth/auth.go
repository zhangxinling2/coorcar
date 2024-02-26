package auth

import (
	"context"
	authpb "coolcar/auth/api/gen/v1"
	"coolcar/auth/dao"
	"time"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service struct {
	Logger *zap.Logger
	authpb.UnimplementedAuthServiceServer
	TokenGenerator TokenGenerator
	OpenIdResolve  OpenIdResolve
	Mongo          *dao.Mongo
	TokenExpire    time.Duration
}
type OpenIdResolve interface {
	Resolve(code string) (string, error)
}
type TokenGenerator interface {
	GenerateToken(accountId string, expire time.Duration) (string, error)
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
	tk, err := s.TokenGenerator.GenerateToken(accountId.String(), time.Duration(s.TokenExpire))
	if err != nil {
		s.Logger.Fatal("cannot generate token" + err.Error())
		return nil, status.Errorf(codes.Internal, "cannot generate token:%v", err)
	}
	return &authpb.LoginResponse{
		AccessToken: tk,
		ExpireIn:    int32(s.TokenExpire.Seconds()),
	}, nil
}
