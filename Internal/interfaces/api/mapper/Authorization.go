package mapper

import (
	"DAJ/Internal/domain/entity"
	"DAJ/Internal/interfaces/api/dto"
)

func UserDTOtoUser(userDTO dto.UserDTO) entity.User {
	return entity.User{
		Name:     userDTO.Username,
		Password: userDTO.Password,
	}
}

func UsertoUserDTO(user entity.User) dto.UserDTO {
	return dto.UserDTO{
		Username: user.Name,
		Password: user.Password,
	}
}
func DTOtoAccessToken(accessTokenDTO dto.AccessTokenDTO) entity.AccessToken {
	return entity.AccessToken{
		Str:        accessTokenDTO.Str,
		ExpireTime: accessTokenDTO.ExpireTime,
	}
}

func AccessTokenToDTO(accessToken entity.AccessToken) dto.AccessTokenDTO {
	return dto.AccessTokenDTO{
		Str:        accessToken.Str,
		ExpireTime: accessToken.ExpireTime,
	}
}
func DTOtoRefreshToken(refreshTokenDTO dto.RefreshTokenDTO) entity.RefreshToken {
	return entity.RefreshToken{
		Str:        refreshTokenDTO.Str,
		ExpireTime: refreshTokenDTO.ExpireTime,
	}
}

func RefreshTokenToDTO(refreshToken entity.RefreshToken) dto.RefreshTokenDTO {
	return dto.RefreshTokenDTO{
		Str:        refreshToken.Str,
		ExpireTime: refreshToken.ExpireTime,
	}
}
