package provision

import (
	"context"

	"github.com/X0JIO/nebula-api/internal/platform/xray"
)

func (s *Service) addReality(
	ctx context.Context,
	uuid string,
	email string,
) error {

	inbound, err := s.xray.FindInboundByProtocol(ctx, "reality")
	if err != nil {
		return err
	}

	return s.xray.AddUser(ctx, xray.AddUserRequest{
		InboundID: inbound.ID,
		UUID:      uuid,
		Email:     email,
		Protocol:  "reality",
	})
}
