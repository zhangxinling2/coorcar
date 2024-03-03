package car

import (
	"context"
	rentalpb "coolcar/rental/api/gen/v1"
	"coolcar/shared/id"
)

type Manager struct {
}

func (c *Manager) Verify(ctx context.Context, id id.CarId, loc *rentalpb.Location) error {
	return nil
}
func (c *Manager) UnLock(ctx context.Context, id id.CarId) error {
	return nil
}
