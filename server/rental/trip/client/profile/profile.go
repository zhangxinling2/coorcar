package profile

import (
	"context"
	rentalpb "coolcar/rental/api/gen/v1"
	"coolcar/shared/id"
	"encoding/base64"
	"fmt"

	"google.golang.org/protobuf/proto"
)

type Fetcher interface {
	GetProfile(ctx context.Context, req *rentalpb.GetProfileRequest) (*rentalpb.Profile, error)
}
type Manager struct {
	Fetcher Fetcher
}

func (m *Manager) Verify(ctx context.Context, aid id.AccountId) (id.IdentityId, error) {
	nilId := id.IdentityId("")
	p, err := m.Fetcher.GetProfile(ctx, &rentalpb.GetProfileRequest{})
	if err != nil {
		return nilId, fmt.Errorf("cannot get profile : %v", err)
	}
	if p.IdentityStatus != rentalpb.IdentityStatus_VERIFIED {
		return nilId, fmt.Errorf("invalid identity status")
	}
	b, err := proto.Marshal(p.Identity)
	if err != nil {
		return nilId, fmt.Errorf("cannot marshal identity:%v", err)
	}
	return id.IdentityId(base64.StdEncoding.EncodeToString(b)), nil
}
