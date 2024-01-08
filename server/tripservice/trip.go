package tripservice

import (
	"context"
	trippb "coorcar/proto/gen/go"
)

type Service struct {
	*trippb.UnimplementedTripServiceServer
}

func (*Service) GetTrip(ctx context.Context, req *trippb.GetTripRequest) (*trippb.GetTripResponse, error) {
	return &trippb.GetTripResponse{
		Id: req.Id,
		Trip: &trippb.Trip{
			Start:       "abc",
			End:         "def",
			DurationSec: 3600,
			FeeCent:     10000,
			StartPos: &trippb.Location{
				Lagitude:  30,
				Longitude: 120,
			},
			Status: trippb.TripStatus_IN_PROGRESS,
		},
	}, nil
}
