package jsonMapper

import (
	"github.com/ViPDanger/dajs/internal/domain/entity"
	"github.com/ViPDanger/dajs/internal/infastructure/json/jsonDTO"
)

// []Status -> []StatusDTO
func ToStatusDTOList(entitys []entity.Status) []jsonDTO.StatusDTO {
	result := make([]jsonDTO.StatusDTO, len(entitys))
	for i, status := range entitys {
		result[i] = ToStatusDTO(status)
	}
	return result
}

func ToStatusEntityList(dtos []jsonDTO.StatusDTO) []entity.Status {
	result := make([]entity.Status, len(dtos))
	for i, dto := range dtos {
		result[i] = ToStatusEntity(dto)
	}
	return result
}

// Status -> StatusDTO
func ToStatusDTO(entity entity.Status) jsonDTO.StatusDTO {
	return jsonDTO.StatusDTO{
		Description: entity.Description,
		Id:          entity.ID,
		Name:        entity.Name,
	}
}

func ToStatusEntity(dto jsonDTO.StatusDTO) entity.Status {
	return entity.Status{
		Description: dto.Description,
		ID:          dto.Id,
		Name:        dto.Name,
	}
}
