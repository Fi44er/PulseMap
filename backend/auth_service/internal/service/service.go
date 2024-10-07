package service

import (
	"context"

	"github.com/Fi44er/PulseMap/backend/auth_service/internal/dto"
	"github.com/Fi44er/PulseMap/backend/auth_service/internal/model"
)

type IAuthService interface {
  Register(ctx context.Context, req *dto.RegisterDTO) (*model.User, error)
}
