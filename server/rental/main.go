package main

import (
	"context"
	rentalpb "coolcar/rental/api/gen/v1"
	"coolcar/rental/profile"
	profDAO "coolcar/rental/profile/dao"
	"coolcar/rental/trip"
	"coolcar/rental/trip/client/car"
	"coolcar/rental/trip/client/poi"
	profClient "coolcar/rental/trip/client/profile"
	tripDAO "coolcar/rental/trip/dao"
	"coolcar/shared/server"
	"log"

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
	db := mo.Database("coolcar")
	tripM := tripDAO.NewMongo(db)
	profM := profDAO.NewMongo(db)
	profService := &profile.Service{
		Logger: logger,
		Mongo:  profM,
	}
	logger.Sugar().Fatal(server.RunGRPCServer(&server.GRPCConfig{
		Name:              "trip service",
		Addr:              ":8082",
		AuthPublicKeyFile: "shared/auth/public.key",
		Logger:            logger,
		RegisterFunc: func(s *grpc.Server) {
			rentalpb.RegisterTripServiceServer(s, &trip.Service{
				Logger: logger,
				Mongo:  tripM,
				ProfileManager: &profClient.Manager{
					Fetcher: profService,
				},
				CarManager: &car.Manager{},
				PoiManager: &poi.Manager{},
			})
			rentalpb.RegisterProfileServiceServer(s, profService)
		},
	}))
}
