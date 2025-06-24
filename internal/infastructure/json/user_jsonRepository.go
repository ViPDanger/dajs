package json

import (
	"DAJ/internal/domain/entity"
	"DAJ/internal/domain/repository"
	"DAJ/internal/infastructure/json/jsonDTO"
	"DAJ/internal/infastructure/json/jsonMapper"
)

type UserRepository struct {
	defaultJSONRepository[entity.User, jsonDTO.UserDTO]
}

func NewUserRepository(filepath string) (repository.Repository[entity.User], error) {
	r := UserRepository{}
	defaultRepository, err := NewJSONRepository(filepath, jsonMapper.ToUserDTO, jsonMapper.ToUserEntity, r.UserPathFunc)
	if err != nil {
		return nil, err
	}
	r.defaultJSONRepository = *defaultRepository
	return &r, nil
}
func (r *UserRepository) UserPathFunc(u *entity.User) string {
	return r.fileDirectory + u.Name + defaultFileType
}
