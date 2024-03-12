package main

import (
	"context"
	authpb "coolcar/auth/api/gen/v1"
	"coolcar/auth/auth"
	"coolcar/auth/dao"
	"coolcar/auth/token"
	"coolcar/auth/wechat"
	"coolcar/shared/server"
	"io"
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func main() {
	//组装zap日志
	logger, err := server.NewZapLog()
	if err != nil {
		log.Fatal("can not create zap log", logger)
	}
	//组装mongoDB
	c := context.Background()
	mo, err := mongo.Connect(c, options.Client().ApplyURI("mongodb://localhost:27017/coolcar?readPreference=primary&ssl=false"))
	if err != nil {
		logger.Fatal("cannot connect db", zap.Error(err))
	}
	m := dao.NewMongo(mo.Database("coolcar"))
	//组装privateKey
	file, err := os.Open("auth/private.key")
	if err != nil {
		logger.Fatal("cannot open private.key", zap.Error(err))
	}
	pkBytes, err := io.ReadAll(file)
	if err != nil {
		logger.Fatal("cannot read private.key", zap.Error(err))
	}
	pk, err := jwt.ParseRSAPrivateKeyFromPEM(pkBytes)
	if err != nil {
		logger.Fatal("cannot parse private.key", zap.Error(err))
	}
	logger.Sugar().Fatal(server.RunGRPCServer(&server.GRPCConfig{
		Name:   "auth service",
		Addr:   ":8081",
		Logger: logger,
		RegisterFunc: func(s *grpc.Server) {
			authpb.RegisterAuthServiceServer(s, &auth.Service{
				OpenIdResolve: &wechat.Service{
					AppId:     "wx9fc227d95b260fb3",
					AppSecret: "602e30ce220a6732503f85eea2807f7a",
				},
				Mongo:          m,
				Logger:         logger,
				TokenGenerator: token.NewJwtTokenGen("coolcar/auth", pk),
				TokenExpire:    10 * time.Second,
			})
		},
	}))
}
