package mapper

import (
	"DAJ/internal/domain/entity"
	"DAJ/internal/interfaces/api/dto"
)

// Map -> MapDTO
func ToMapEntity(d dto.MapDTO) entity.Map {
	return entity.Map{
		ID:                    d.ID,
		Path:                  d.Path,
		Hold:                  d.Hold,
		HoldTime:              d.HoldTime,
		X:                     d.X,
		Y:                     d.Y,
		CurrentLayer:          d.CurrentLayer,
		UserMapSizePercent:    d.UserMapSizePercent,
		Xmirror:               d.Xmirror,
		Ymirror:               d.Ymirror,
		RotateMap:             d.RotateMap,
		FogOfWarPlayerShow:    d.FogOfWarPlayerShow,
		VideoVolumeInt:        d.VideoVolumeInt,
		LastWallId:            d.LastWallId,
		LockMap:               d.LockMap,
		DeepColor:             d.DeepColor,
		PlayersFogTranperensy: d.PlayersFogTranperensy,
		MasterFogOpacity:      d.MasterFogOpacity,
		Fog:                   ToMapFogEntitySlice(d.Fog),
		Monsters:              ToMapMonsterEntitySlice(d.Monsters),
		Players:               ToMapPlayerEntitySlice(d.Players),
		Npc:                   ToMapNpcEntitySlice(d.Npc),
		Props:                 ToMapPropEntitySlice(d.Props),
		PropsList:             ToMapPropsListEntitySlice(d.PropsList),
	}
}

func ToMapDTO(e entity.Map) dto.MapDTO {
	return dto.MapDTO{
		ID:                    e.ID,
		Path:                  e.Path,
		Hold:                  e.Hold,
		HoldTime:              e.HoldTime,
		X:                     e.X,
		Y:                     e.Y,
		CurrentLayer:          e.CurrentLayer,
		UserMapSizePercent:    e.UserMapSizePercent,
		Xmirror:               e.Xmirror,
		Ymirror:               e.Ymirror,
		RotateMap:             e.RotateMap,
		FogOfWarPlayerShow:    e.FogOfWarPlayerShow,
		VideoVolumeInt:        e.VideoVolumeInt,
		LastWallId:            e.LastWallId,
		LockMap:               e.LockMap,
		DeepColor:             e.DeepColor,
		PlayersFogTranperensy: e.PlayersFogTranperensy,
		MasterFogOpacity:      e.MasterFogOpacity,
		Fog:                   ToMapFogDTOSlice(e.Fog),
		Monsters:              ToMapMonsterDTOSlice(e.Monsters),
		Players:               ToMapPlayerDTOSlice(e.Players),
		Npc:                   ToMapNpcDTOSlice(e.Npc),
		Props:                 ToMapPropDTOSlice(e.Props),
		PropsList:             ToMapPropsListDTOSlice(e.PropsList),
	}
}

// MapFog -> MapFogDTO
func ToMapFogEntitySlice(dtos []dto.MapFogDTO) []entity.MapFog {
	result := make([]entity.MapFog, len(dtos))
	for i, d := range dtos {
		result[i] = entity.MapFog(d)
	}
	return result
}

func ToMapFogDTOSlice(entities []entity.MapFog) []dto.MapFogDTO {
	result := make([]dto.MapFogDTO, len(entities))
	for i, e := range entities {
		result[i] = dto.MapFogDTO(e)
	}
	return result
}

// MapMonster -> MapMonsterDTO
func ToMapMonsterEntitySlice(dtos []dto.MapMonsterDTO) []entity.MapMonster {
	result := make([]entity.MapMonster, len(dtos))
	for i, d := range dtos {
		result[i] = entity.MapMonster{
			ID:                   d.ID,
			Name:                 d.Name,
			X:                    d.X,
			Y:                    d.Y,
			ShowFigureToPlayer:   d.ShowFigureToPlayer,
			EyesightEnabled:      d.EyesightEnabled,
			CellsOfEyesight:      d.CellsOfEyesight,
			IHaveLight:           d.IHaveLight,
			TorchValue:           d.TorchValue,
			TorchValueSecond:     d.TorchValueSecond,
			HitPoints:            d.HitPoints,
			CurrentHitPoints:     d.CurrentHitPoints,
			TempHitPoints:        d.TempHitPoints,
			TempCurrentHitPoints: d.TempCurrentHitPoints,
			IniColor:             d.IniColor,
			LegendaryActions:     d.LegendaryActions,
			CustomStatuses:       ToMapStatusEntitySlice(d.CustomStatuses),
		}
	}
	return result
}

func ToMapMonsterDTOSlice(entities []entity.MapMonster) []dto.MapMonsterDTO {
	result := make([]dto.MapMonsterDTO, len(entities))
	for i, e := range entities {
		result[i] = dto.MapMonsterDTO{
			ID:                   e.ID,
			Name:                 e.Name,
			X:                    e.X,
			Y:                    e.Y,
			ShowFigureToPlayer:   e.ShowFigureToPlayer,
			EyesightEnabled:      e.EyesightEnabled,
			CellsOfEyesight:      e.CellsOfEyesight,
			IHaveLight:           e.IHaveLight,
			TorchValue:           e.TorchValue,
			TorchValueSecond:     e.TorchValueSecond,
			HitPoints:            e.HitPoints,
			CurrentHitPoints:     e.CurrentHitPoints,
			TempHitPoints:        e.TempHitPoints,
			TempCurrentHitPoints: e.TempCurrentHitPoints,
			IniColor:             e.IniColor,
			LegendaryActions:     e.LegendaryActions,
			CustomStatuses:       ToMapStatusDTOSlice(e.CustomStatuses),
		}
	}
	return result
}

// MapStatus -> MapStatusDTO
func ToMapStatusEntitySlice(dtos []dto.MapStatusDTO) []entity.MapStatus {
	result := make([]entity.MapStatus, len(dtos))
	for i, d := range dtos {
		result[i] = entity.MapStatus(d)
	}
	return result
}

func ToMapStatusDTOSlice(entities []entity.MapStatus) []dto.MapStatusDTO {
	result := make([]dto.MapStatusDTO, len(entities))
	for i, e := range entities {
		result[i] = dto.MapStatusDTO(e)
	}
	return result
}

// MapPlayer -> MapPlayerDTO
func ToMapPlayerEntitySlice(dtos []dto.MapPlayerDTO) []entity.MapPlayer {
	result := make([]entity.MapPlayer, len(dtos))
	for i, d := range dtos {
		result[i] = entity.MapPlayer(d)
	}
	return result
}

func ToMapPlayerDTOSlice(entities []entity.MapPlayer) []dto.MapPlayerDTO {
	result := make([]dto.MapPlayerDTO, len(entities))
	for i, e := range entities {
		result[i] = dto.MapPlayerDTO(e)
	}
	return result
}

// MapNPC -> MapNPCDTO
func ToMapNpcEntitySlice(dtos []dto.MapNpcDTO) []entity.MapNpc {
	result := make([]entity.MapNpc, len(dtos))
	for i, d := range dtos {
		result[i] = entity.MapNpc(d)
	}
	return result
}

func ToMapNpcDTOSlice(entities []entity.MapNpc) []dto.MapNpcDTO {
	result := make([]dto.MapNpcDTO, len(entities))
	for i, e := range entities {
		result[i] = dto.MapNpcDTO(e)
	}
	return result
}

// --- Prop ---
func ToMapPropEntitySlice(dtos []dto.MapPropDTO) []entity.MapProp {
	result := make([]entity.MapProp, len(dtos))
	for i, d := range dtos {
		result[i] = entity.MapProp(d)
	}
	return result
}

func ToMapPropDTOSlice(entities []entity.MapProp) []dto.MapPropDTO {
	result := make([]dto.MapPropDTO, len(entities))
	for i, e := range entities {
		result[i] = dto.MapPropDTO(e)
	}
	return result
}

// --- PropsList ---
func ToMapPropsListEntitySlice(dtos []dto.MapPropsListDTO) []entity.MapPropsList {
	result := make([]entity.MapPropsList, len(dtos))
	for i, d := range dtos {
		result[i] = entity.MapPropsList{
			Name:  d.Name,
			Props: ToMapPropEntitySlice(d.Props),
		}
	}
	return result
}

func ToMapPropsListDTOSlice(entities []entity.MapPropsList) []dto.MapPropsListDTO {
	result := make([]dto.MapPropsListDTO, len(entities))
	for i, e := range entities {
		result[i] = dto.MapPropsListDTO{
			Name:  e.Name,
			Props: ToMapPropDTOSlice(e.Props),
		}
	}
	return result
}
