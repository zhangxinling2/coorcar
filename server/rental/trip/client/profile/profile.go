package profile

import (
	"context"
	"coolcar/shared/id"
)

type Manager struct {
}

func (p *Manager) Verify(ctx context.Context, aid id.AccountId) (id.IdentityId, error) {
	return id.IdentityId("identity1"), nil
}
