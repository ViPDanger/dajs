package mapper

import (
	"github.com/ViPDanger/dajs/internal/domain/entity"
	"github.com/ViPDanger/dajs/internal/interfaces/api/dto"
)

func ToUserEntity(userDTO dto.UserDTO) entity.User {
	return entity.User{
		Name:     userDTO.Username,
		Password: userDTO.Password,
	}
}

func ToUserDTO(user entity.User) dto.UserDTO {
	return dto.UserDTO{
		Username: user.Name,
		Password: user.Password,
	}
}
func ToAccessTokenEntity(accessTokenDTO dto.AccessTokenDTO) entity.AccessToken {
	return entity.AccessToken{
		Str:        accessTokenDTO.Str,
		ExpireTime: accessTokenDTO.ExpireTime,
	}
}

func ToAccessTokenDTO(accessToken entity.AccessToken) dto.AccessTokenDTO {
	return dto.AccessTokenDTO{
		Str:        accessToken.Str,
		ExpireTime: accessToken.ExpireTime,
	}
}
func ToRefreshTokenEntity(refreshTokenDTO dto.RefreshTokenDTO) entity.RefreshToken {
	return entity.RefreshToken{
		Str:        refreshTokenDTO.Str,
		ExpireTime: refreshTokenDTO.ExpireTime,
	}
}

func ToRefreshTokenDTO(refreshToken entity.RefreshToken) dto.RefreshTokenDTO {
	return dto.RefreshTokenDTO{
		Str:        refreshToken.Str,
		ExpireTime: refreshToken.ExpireTime,
	}
}
