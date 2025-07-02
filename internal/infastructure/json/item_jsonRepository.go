package json

import (
	"fmt"
	"strings"

	"github.com/ViPDanger/dajs/internal/domain/entity"
	"github.com/ViPDanger/dajs/internal/domain/repository"
	"github.com/ViPDanger/dajs/internal/infastructure/json/jsonDTO"
	"github.com/ViPDanger/dajs/internal/infastructure/json/jsonMapper"
)

type itemJSONRepository struct {
	defaultJSONRepository[entity.Item, jsonDTO.ItemDTO]
}

func NewItemRepository(filepath string) (repository.Repository[entity.Item], error) {
	r := itemJSONRepository{}
	repository, err := NewJSONRepository(filepath, jsonMapper.ToItemDTO, jsonMapper.ToItemEntity, r.itemPathFunc)
	if err != nil {
		return nil, fmt.Errorf("NewItemRepository()/%w", err)
	}
	r.defaultJSONRepository = *repository
	return &r, nil
}

func (r *itemJSONRepository) itemPathFunc(object *entity.Item) string {
	item := (*object).GetSimpleItem()
	var path string
	switch {
	case containsTag(item.Tags, "Лёгкий доспех"):
		path = "Personal items/Базовые предметы/Броня/Легкая броня"
	case containsTag(item.Tags, "Средний доспех"):
		path = "Personal items/Базовые предметы/Броня/Средняя броня"
	case containsTag(item.Tags, "Тяжёлый доспех"):
		path = "Personal items/Базовые предметы/Броня/Тяжёлая броня"
	case containsTag(item.Tags, "Доспех"):
		path = "Personal items/Базовые предметы/Броня/Щиты"
	case containsTag(item.Tags, "Весовые товары"):
		path = "Personal items/Базовые предметы/Весовые товары"
	case containsTag(item.Tags, "Инструменты"):
		path = "Personal items/Базовые предметы/Инструменты"
	case containsTag(item.Tags, "Воинское дальнобойное оружие"):
		path = "Personal items/Базовые предметы/Оружие/Воинское дальнобойное оружие"
	case containsTag(item.Tags, "Воинское рукопашное оружие"):
		path = "Personal items/Базовые предметы/Оружие/Воинское рукопашное оружие"
	case containsTag(item.Tags, "Простое дальнобойное оружие"):
		path = "Personal items/Базовые предметы/Оружие/Воинское дальнобойное оружие"
	case containsTag(item.Tags, "Простое рукопашное оружие"):
		path = "Personal items/Базовые предметы/Оружие/Воинское рукопашное оружие"
	case containsTag(item.Tags, "Снаряжение"):
		path = "Personal items/Базовые предметы/Снаряжение"
	case containsTag(item.Tags, "Легендарный"):
		path = "Personal items/Базовые предметы/DMG/Легендарные"
	case containsTag(item.Tags, "Необычный"):
		path = "Personal items/Базовые предметы/DMG/Необычные"
	case containsTag(item.Tags, "Обычный"):
		path = "Personal items/Базовые предметы/DMG/Обычные магические"
	case containsTag(item.Tags, "Очень редкий"):
		path = "Personal items/Базовые предметы/DMG/Очень редкие"
	case containsTag(item.Tags, "Редкий"):
		path = "Personal items/Базовые предметы/DMG/Редкие"
	case containsTag(item.Tags, "Свиток"):
		path = "Personal items/Базовые предметы/DMG/Свитки"

	default:
		path = "Personal items"
		for i := len(item.Tags); i > 0; i-- {
			path += item.Tags[i] + "/"
		}
	}
	return r.fileDirectory + "/" + path + "/" + item.Name + defaultFileType
}
func containsTag(tags []string, target string) bool {
	for _, t := range tags {
		if strings.TrimSpace(t) == target {
			return true
		}
	}
	return false
}
