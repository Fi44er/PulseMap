package convert

import (
	"github.com/Fi44er/PulseMap/backend/auth_service/internal/dto"
	"github.com/Fi44er/PulseMap/backend/auth_service/internal/model"
	desc "github.com/Fi44er/PulseMap/backend/auth_service/proto/out"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToUserFromService(user *model.User) *desc.User {
  var updatedAt *timestamppb.Timestamp
  if user.UpdatedAt != nil {
    updatedAt = timestamppb.New(*user.UpdatedAt)
  }

  return &desc.User{
    Id: user.ID,
    Username: user.Username,
    Email: user.Email,
    PasswordHash: user.PasswordHash,
    CreatedAt: timestamppb.New(user.CreatedAt),
    UpdatedAt: updatedAt,
  }
}

func ToRegisterReqFromDesc(req *desc.RegisterReq) *dto.RegisterDTO {
  return &dto.RegisterDTO{
    Username: req.Username,
    Email: req.Email,
    Password: req.Password,
  }
}
