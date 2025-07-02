package json

import (
	"fmt"

	"github.com/ViPDanger/dajs/internal/domain/entity"
	"github.com/ViPDanger/dajs/internal/domain/repository"
	"github.com/ViPDanger/dajs/internal/infastructure/json/jsonDTO"
	"github.com/ViPDanger/dajs/internal/infastructure/json/jsonMapper"
)

type UserRepository struct {
	defaultJSONRepository[entity.User, jsonDTO.UserDTO]
}

func NewUserRepository(filepath string) (repository.Repository[entity.User], error) {
	r := UserRepository{}
	defaultRepository, err := NewJSONRepository(filepath, jsonMapper.ToUserDTO, jsonMapper.ToUserEntity, r.UserPathFunc)
	if err != nil {
		return nil, fmt.Errorf("NewUserRepository()/%w", err)
	}
	r.defaultJSONRepository = *defaultRepository
	return &r, nil
}
func (r *UserRepository) UserPathFunc(u *entity.User) string {
	return r.fileDirectory + u.Name + defaultFileType
}
