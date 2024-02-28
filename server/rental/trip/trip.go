package trip

import (
	"context"

	rentalpb "coolcar/rental/api/gen/v1"
	"coolcar/rental/trip/dao"
	"coolcar/shared/auth"
	"coolcar/shared/id"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service struct {
	Logger         *zap.Logger
	Mongo          *dao.Mongo
	ProfileManager ProfileManager
	CarManager     CarManager
	PoiManager     PoiManager
	rentalpb.UnimplementedTripServiceServer
}
type ProfileManager interface {
	Verify(ctx context.Context, id id.AccountId) (id.IndetityId, error)
}
type CarManager interface {
	Verify(ctx context.Context, id id.CarId, loc *rentalpb.Location) error
	UnLock(ctx context.Context, id id.CarId) error
}
type PoiManager interface {
	Resolve(ctx context.Context, loc *rentalpb.Location) (string, error)
}

func (s *Service) CreateTrip(ctx context.Context, req *rentalpb.CreateTripRequest) (*rentalpb.TripEntity, error) {
	aid, err := auth.AccountIdFromContext(ctx)
	if err != nil {
		return nil, err
	}
	iId, err := s.ProfileManager.Verify(ctx, aid)
	if err != nil {
		return nil, status.Error(codes.FailedPrecondition, "")
	}
	err = s.CarManager.Verify(ctx, id.CarId(req.CarId), req.Start.Location)
	if err != nil {
		return nil, status.Error(codes.FailedPrecondition, "")
	}
	poi, err := s.PoiManager.Resolve(ctx, req.Start.Location)
	if err != nil {
		s.Logger.Info("cannot resolve location to name", zap.Stringer("location", req.Start.Location))
	}
	start := &rentalpb.LocationStatus{
		Location: req.Start.Location,
		PoiName:  poi,
	}
	tr, err := s.Mongo.CreateTrip(ctx, &rentalpb.Trip{
		AccountId:  aid.String(),
		CarId:      req.CarId,
		Start:      start,
		Current:    start,
		IdentityId: iId.String(),
	})
	if err != nil {
		return nil, status.Error(codes.FailedPrecondition, "failed to create trip")
	}
	go func() {
		err := s.CarManager.UnLock(context.Background(), id.CarId(req.CarId))
		if err != nil {
			s.Logger.Warn("cannot unlock car", zap.Error(err))
		}
	}()
	return &rentalpb.TripEntity{
		Id:   tr.ID.Hex(),
		Trip: tr.Trip,
	}, nil
}
func (s *Service) GetTrip(ctx context.Context, req *rentalpb.GetTripRequest) (*rentalpb.Trip, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTrip not implemented")
}
func (s *Service) GetTrips(ctx context.Context, req *rentalpb.GetTripsRequest) (*rentalpb.GetTripsReponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTrips not implemented")
}
func (s *Service) UpdateTrip(ctx context.Context, req *rentalpb.UpdateTripRequest) (*rentalpb.Trip, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTrip not implemented")
}
