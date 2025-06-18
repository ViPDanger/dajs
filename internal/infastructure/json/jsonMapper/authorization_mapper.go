package jsonMapper

import (
	"DAJ/internal/domain/entity"
	"DAJ/internal/infastructure/json/jsonDTO"
)

func ToUserEntity(userDTO jsonDTO.UserDTO) entity.User {
	return entity.User{
		Name:     userDTO.Username,
		Password: userDTO.Password,
	}
}

func ToUserDTO(user entity.User) jsonDTO.UserDTO {
	return jsonDTO.UserDTO{
		Username: user.Name,
		Password: user.Password,
	}
}
func ToAccessTokenEntity(accessTokenDTO jsonDTO.AccessTokenDTO) entity.AccessToken {
	return entity.AccessToken{
		Str:        accessTokenDTO.Str,
		ExpireTime: accessTokenDTO.ExpireTime,
	}
}

func ToAccessTokenDTO(accessToken entity.AccessToken) jsonDTO.AccessTokenDTO {
	return jsonDTO.AccessTokenDTO{
		Str:        accessToken.Str,
		ExpireTime: accessToken.ExpireTime,
	}
}
func ToRefreshTokenEntity(refreshTokenDTO jsonDTO.RefreshTokenDTO) entity.RefreshToken {
	return entity.RefreshToken{
		Str:        refreshTokenDTO.Str,
		ExpireTime: refreshTokenDTO.ExpireTime,
	}
}

func ToRefreshTokenDTO(refreshToken entity.RefreshToken) jsonDTO.RefreshTokenDTO {
	return jsonDTO.RefreshTokenDTO{
		Str:        refreshToken.Str,
		ExpireTime: refreshToken.ExpireTime,
	}
}
