package server

import (
	"coolcar/shared/auth"
	"net"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type GRPCConfig struct {
	Name              string
	Addr              string
	AuthPublicKeyFile string
	RegisterFunc      func(*grpc.Server)
	Logger            *zap.Logger
}

func RunGRPCServer(conf *GRPCConfig) error {
	n, err := net.Listen("tcp", conf.Addr)
	if err != nil {
		conf.Logger.Fatal("cannot listen", zap.Error(err))
	}
	var opts []grpc.ServerOption
	if conf.AuthPublicKeyFile != "" {
		in, err := auth.Interceptor(conf.AuthPublicKeyFile)
		if err != nil {
			conf.Logger.Fatal("Interceptor listen", zap.Error(err))
		}
		opts = append(opts, grpc.UnaryInterceptor(in))
	}
	s := grpc.NewServer(opts...)
	conf.RegisterFunc(s)
	conf.Logger.Info("server started", zap.String("addr", conf.Addr))
	return s.Serve(n)
}
