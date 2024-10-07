package auth

import (
	"github.com/Fi44er/PulseMap/backend/auth_service/internal/repository"
	def "github.com/Fi44er/PulseMap/backend/auth_service/internal/service"
)

var _ def.IAuthService = (*service)(nil)

type service struct {
  userRepository repository.IAuthRepository
}

func NewService(userRepository repository.IAuthRepository) *service {
  return &service{
    userRepository: userRepository,
  }
}
