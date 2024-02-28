package poi

import (
	"context"
	rentalpb "coolcar/rental/api/gen/v1"
	"hash/fnv"

	"google.golang.org/protobuf/proto"
)

var poi = []string{
	"陆家嘴",
	"天安门",
	"广州塔",
	"天河体育中心",
}

type Manager struct {
}

func (*Manager) Resolve(ctx context.Context, loc *rentalpb.Location) (string, error) {
	b, err := proto.Marshal(loc)
	if err != nil {
		return "", err
	}
	h := fnv.New32()
	n, err := h.Write(b)
	if err != nil {
		return "", err
	}
	return poi[n%len(poi)], nil
}
