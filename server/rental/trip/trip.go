package trip

import (
	"context"
	"math/rand"
	"time"

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
	Verify(ctx context.Context, id id.AccountId) (id.IdentityId, error)
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
	err = s.CarManager.Verify(ctx, id.CarId(req.CarId), req.Start)
	if err != nil {
		return nil, status.Error(codes.FailedPrecondition, "")
	}
	start := s.calcCurrentStatus(ctx, &rentalpb.LocationStatus{
		Location:     req.Start,
		TimestampSec: nowFunc(),
	}, req.Start)
	tr, err := s.Mongo.CreateTrip(ctx, &rentalpb.Trip{
		AccountId:  aid.String(),
		CarId:      req.CarId,
		Start:      start,
		Current:    start,
		Status:     rentalpb.TripStatus_IN_PROGRESS,
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
	aid, err := auth.AccountIdFromContext(ctx)
	if err != nil {
		return nil, err
	}
	tr, err := s.Mongo.GetTrip(ctx, id.TripId(req.Id), aid)
	if err != nil {
		s.Logger.Error("cannot get trip", zap.Error(err))
		return nil, status.Error(codes.NotFound, "")
	}
	return tr.Trip, nil
}
func (s *Service) GetTrips(ctx context.Context, req *rentalpb.GetTripsRequest) (*rentalpb.GetTripsReponse, error) {
	aid, err := auth.AccountIdFromContext(ctx)
	s.Logger.Info("get trips for:" + aid.String())
	if err != nil {
		return nil, err
	}
	trs, err := s.Mongo.GetTrips(ctx, aid, req.Status)
	if err != nil {
		s.Logger.Error("cannot get trips", zap.Error(err))
		return nil, status.Error(codes.Internal, "")
	}
	var res []*rentalpb.TripEntity
	for _, tr := range trs {
		ent := &rentalpb.TripEntity{
			Id:   tr.ID.String(),
			Trip: tr.Trip,
		}
		res = append(res, ent)
	}
	return &rentalpb.GetTripsReponse{
		Trips: res,
	}, nil
}
func (s *Service) UpdateTrip(ctx context.Context, req *rentalpb.UpdateTripRequest) (*rentalpb.Trip, error) {
	aid, err := auth.AccountIdFromContext(ctx)
	if err != nil {
		return nil, err
	}
	tid := id.TripId(req.Id)
	tr, err := s.Mongo.GetTrip(ctx, tid, aid)
	if err != nil {
		return nil, status.Error(codes.NotFound, "")
	}
	if tr.Trip.Status == rentalpb.TripStatus_FINISHED {
		return nil, status.Error(codes.FailedPrecondition, "cannot update a finished trip")
	}
	if tr.Trip.Current == nil {
		s.Logger.Error("trip without current set", zap.String("id", tid.String()))
		return nil, status.Error(codes.Internal, "")
	}
	cur := tr.Trip.Current.Location
	if req.Current != nil {
		cur = req.Current
	}

	tr.Trip.Current = s.calcCurrentStatus(ctx, tr.Trip.Current, cur)
	if req.EndTrip {
		tr.Trip.End = tr.Trip.Current
		tr.Trip.Status = rentalpb.TripStatus_FINISHED
	}
	err = s.Mongo.UpdateTrip(ctx, id.TripId(req.Id), aid, tr.UpdatedAt, tr.Trip)
	if err != nil {
		return nil, status.Error(codes.Aborted, "")
	}
	return tr.Trip, nil
}

const (
	centsPerSec = 0.7
	kmPerSec    = 0.02
)

var nowFunc = func() int64 {
	return time.Now().Unix()
}

func (s *Service) calcCurrentStatus(ctx context.Context, last *rentalpb.LocationStatus, cur *rentalpb.Location) *rentalpb.LocationStatus {
	now := nowFunc()
	elaspedSec := float64(now - last.TimestampSec)
	poi, err := s.PoiManager.Resolve(ctx, cur)
	if err != nil {
		s.Logger.Info("cannot resolve location to name", zap.Stringer("location", cur))
	}
	return &rentalpb.LocationStatus{
		Location:     cur,
		KmDriven:     last.KmDriven + kmPerSec*elaspedSec*2*rand.Float64(),
		FeeCent:      last.FeeCent + int32(centsPerSec*elaspedSec*2*rand.Float64()),
		PoiName:      poi,
		TimestampSec: now,
	}
}
