package auth

import (
	"context"

	"github.com/Fi44er/PulseMap/backend/auth_service/internal/convert"
	desc "github.com/Fi44er/PulseMap/backend/auth_service/proto/out"
)

func (i *Implementation) Register(ctx context.Context, req *desc.RegisterReq) (*desc.RegisterRes, error) {
  user, err := i.authService.Register(ctx, convert.ToRegisterReqFromDesc(req))
  if err != nil {
    return nil, err
  }

  return &desc.RegisterRes{
    User: convert.ToUserFromService(user),
  }, nil
}
