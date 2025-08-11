package mapper

import (
	"github.com/ViPDanger/dajs/go-api/internal/domain/entity"
	"github.com/ViPDanger/dajs/go-api/internal/interfaces/dto"
)

func ToUserEntity(userDTO dto.UserDTO) entity.User {
	return entity.User{
		Username: userDTO.Username,
		Password: userDTO.Password,
	}
}

func ToUserDTO(user entity.User) dto.UserDTO {
	return dto.UserDTO{
		Username: user.Username,
		Password: user.Password,
	}
}
