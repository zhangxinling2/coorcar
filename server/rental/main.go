package main

import (
	"context"
	rentalpb "coolcar/rental/api/gen/v1"
	"coolcar/rental/trip"
	"coolcar/rental/trip/client/car"
	"coolcar/rental/trip/client/poi"
	"coolcar/rental/trip/client/profile"
	"coolcar/rental/trip/dao"
	"coolcar/shared/auth"
	"log"
	"net"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func main() {
	//组装zap日志
	logger, err := NewZapLog()
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
	n, err := net.Listen("tcp", ":8082")
	if err != nil {
		logger.Fatal("cannot listen", zap.Error(err))
	}
	in, err := auth.Interceptor("shared/auth/public.key")
	if err != nil {
		logger.Fatal("Interceptor listen", zap.Error(err))
	}
	s := grpc.NewServer(grpc.UnaryInterceptor(in))
	rentalpb.RegisterTripServiceServer(s, &trip.Service{
		Logger:         logger,
		Mongo:          m,
		ProfileManager: &profile.Manager{},
		CarManager:     &car.Manager{},
		PoiManager:     &poi.Manager{},
	})
	err = s.Serve(n)
	if err != nil {
		logger.Fatal("cannot server", zap.Error(err))
	}
}
func NewZapLog() (*zap.Logger, error) {

	cfg := zap.NewDevelopmentConfig()
	cfg.EncoderConfig.TimeKey = ""
	return cfg.Build()

}
