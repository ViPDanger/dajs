package mapper

import (
	"github.com/ViPDanger/dajs/internal/domain/entity"
	"github.com/ViPDanger/dajs/internal/interfaces/api/dto"
)

// []Status -> []StatusDTO
func ToStatusDTOList(entitys []entity.Status) []dto.StatusDTO {
	result := make([]dto.StatusDTO, len(entitys))
	for i, status := range entitys {
		result[i] = ToStatusDTO(status)
	}
	return result
}

func ToStatusEntityList(dtos []dto.StatusDTO) []entity.Status {
	result := make([]entity.Status, len(dtos))
	for i, dto := range dtos {
		result[i] = ToStatusEntity(dto)
	}
	return result
}

// Status -> StatusDTO
func ToStatusDTO(entity entity.Status) dto.StatusDTO {
	return dto.StatusDTO{
		Description: entity.Description,
		Id:          entity.ID,
		Name:        entity.Name,
	}
}

func ToStatusEntity(dto dto.StatusDTO) entity.Status {
	return entity.Status{
		Description: dto.Description,
		ID:          dto.Id,
		Name:        dto.Name,
	}
}
