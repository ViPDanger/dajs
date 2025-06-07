package request

import (
	"DAJ/Internal/domain/entity"
	"DAJ/Internal/interfaces/api/dto"
	"DAJ/Internal/interfaces/api/mapper"
	"bytes"
	"encoding/json"
	"net/http"
)

// GET Character
func (r *HttpRepository) GetCharacter(id string) (character entity.Character, err error) {
	req, _ := http.NewRequest("GET", r.Host+"/protected/character/get", nil)
	req.Header.Set("id", id)
	resp, err := r.doProtected(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	var characterDTO dto.CharacterDTO
	if err = json.NewDecoder(resp.Body).Decode(&characterDTO); err != nil {
		return
	}
	character = mapper.ToCharacterEntity(characterDTO)
	return
}

// POST Character
func (r *HttpRepository) NewCharacter(character entity.Character) (id string, err error) {
	body, err := json.Marshal(mapper.ToCharacterDTO(character))
	if err != nil {
		return
	}
	req, _ := http.NewRequest("POST", r.Host+"/protected/character/new", bytes.NewBuffer(body))
	resp, err := r.doProtected(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	var message dto.Message
	if err = json.NewDecoder(resp.Body).Decode(&message); err != nil {
		return
	}
	return message.Message, nil
}

// GET ALL CHARACTER
func (r *HttpRepository) AllCharacter() (characters []entity.Character, err error) {
	req, _ := http.NewRequest("GET", r.Host+"/protected/character/", nil)
	resp, err := r.doProtected(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	charactersDTO := make([]dto.CharacterDTO, 0)
	if err = json.NewDecoder(resp.Body).Decode(&charactersDTO); err != nil {
		return
	}
	for i := range charactersDTO {
		characters = append(characters, mapper.ToCharacterEntity(charactersDTO[i]))
	}
	return
}

// PUT Character
func (r *HttpRepository) SetCharacter(character entity.Character) (id string, err error) {
	body, err := json.Marshal(mapper.ToCharacterDTO(character))
	if err != nil {
		return
	}
	req, _ := http.NewRequest("PUT", r.Host+"/protected/character/set", bytes.NewBuffer(body))
	resp, err := r.doProtected(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	var message dto.Message
	if err = json.NewDecoder(resp.Body).Decode(&message); err != nil {
		return
	}
	return message.Message, nil
}

// DELETE Character
func (r *HttpRepository) DeleteCharacter(id string) (character entity.Character, err error) {
	req, _ := http.NewRequest("DELETE", r.Host+"/protected/character/delete", nil)
	req.Header.Set("id", id)
	resp, err := r.doProtected(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	var characterDTO dto.CharacterDTO
	if err = json.NewDecoder(resp.Body).Decode(&characterDTO); err != nil {
		return
	}
	character = mapper.ToCharacterEntity(characterDTO)
	return
}
