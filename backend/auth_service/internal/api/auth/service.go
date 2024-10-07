package auth

import (
	"github.com/Fi44er/PulseMap/backend/auth_service/internal/service"
	"github.com/Fi44er/PulseMap/backend/auth_service/proto/out"
)

type Implementation struct {
  auth.UnimplementedAuthServer
  authService service.IAuthService
}

func NewImplementation(authService service.IAuthService) *Implementation {
  return &Implementation{
    authService: authService,
  }
}

