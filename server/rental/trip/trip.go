package trip

import (
	"context"

	rentalpb "coolcar/rental/api/gen/v1"

	"go.uber.org/zap"
)

type Service struct {
	Logger *zap.Logger
	rentalpb.UnimplementedTripServiceServer
}

func (s *Service) CreateTrip(ctx context.Context, req *rentalpb.CreateTripRequest) (*rentalpb.CreateTripResponse, error) {
	s.Logger.Info("create trip", zap.String("start", req.Start))
	return &rentalpb.CreateTripResponse{}, nil
}
