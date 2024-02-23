package auth

import (
	"context"
	"coolcar/shared/auth/token"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// Interceptor 传入公钥文件生成UnaryServerInterceptor
func Interceptor(publicKeyFile string) (grpc.UnaryServerInterceptor, error) {
	file, err := os.Open(publicKeyFile)
	if err != nil {
		return nil, fmt.Errorf("connot open public key file:%v", err)
	}
	pkBytes, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("connot read public key file:%v", err)
	}
	pk, err := jwt.ParseRSAPublicKeyFromPEM(pkBytes)
	if err != nil {
		return nil, fmt.Errorf("connot parse public key :%v", err)
	}
	interceptor := interceptor{
		verifier: &token.JWTTokenVerify{
			PublicKey: pk,
		},
	}
	return interceptor.HandleReq, nil
}

type tokenVerfier interface {
	Verify(token string) (string, error)
}
type interceptor struct {
	verifier tokenVerfier
}
type accountIdKey struct {
}

// 处理请求
func (i *interceptor) HandleReq(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
	tk, err := tokenFromCtx(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "")
	}
	aid, err := i.verifier.Verify(tk)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "token not valid : %v", err)
	}
	ctx = ContextWithAccountId(ctx, aid)
	return handler(ctx, req)
}
func tokenFromCtx(ctx context.Context) (string, error) {
	unauthenticated := status.Error(codes.Unauthenticated, "")
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", unauthenticated
	}
	tkn := ""
	for _, v := range md["authorization"] {
		if strings.HasPrefix(v, "Bearer ") {
			tkn = v[len("Bearer "):]
		}
	}
	if tkn == "" {
		return "", unauthenticated
	}
	return tkn, nil
}
func ContextWithAccountId(ctx context.Context, accountId string) context.Context {
	return context.WithValue(ctx, accountIdKey{}, accountId)
}
func AccountIdFromContext(ctx context.Context) (string, error) {
	v := ctx.Value(accountIdKey{})
	aid, ok := v.(string)
	if !ok {
		return "", status.Error(codes.Unauthenticated, "")
	}
	return aid, nil
}
