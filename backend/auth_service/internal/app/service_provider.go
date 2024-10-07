package app

import (
	"log"

	"github.com/Fi44er/PulseMap/backend/auth_service/internal/api/auth"
	"github.com/Fi44er/PulseMap/backend/auth_service/internal/config"
	"github.com/Fi44er/PulseMap/backend/auth_service/internal/repository"
	"github.com/Fi44er/PulseMap/backend/auth_service/internal/service"
)

type serviceProvider struct {
  grpcConfig config.IGRPCConfig
  authRepository repository.IUserRepository
  authService service.IAuthService
  authImpl *auth.Implementation
}

func newServiceProvider() *serviceProvider {
  return &serviceProvider{}
}

func (s *serviceProvider) GRPCConfig() config.IGRPCConfig {
  if s.grpcConfig == nil {
    cfg, err := config.NewGRPCConfig()
    if err != nil {
      log.Fatalf("failed to get grpc config: %s", err.Error())
    }

    s.grpcConfig = cfg
  }

  return s.grpcConfig
} 

func (s *serviceProvider) AuthRepository() repository.IUserRepository {
  if s.authRepository == nil {
  
  }
}
