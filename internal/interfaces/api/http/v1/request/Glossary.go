package request

import (
	"DAJ/internal/domain/entity"
	"DAJ/internal/interfaces/api/dto"
	"DAJ/internal/interfaces/api/mapper"
	"bytes"
	"encoding/json"
	"net/http"
)

// GET Glossary
func (r *HttpRepository) GetGlossary(id string) (Glossary entity.Glossary, err error) {
	req, _ := http.NewRequest("GET", r.Host+"/protected/glossary/get", nil)
	req.Header.Set("id", id)
	resp, err := r.doProtected(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	var GlossaryDTO dto.GlossaryDTO
	if err = json.NewDecoder(resp.Body).Decode(&GlossaryDTO); err != nil {
		return
	}
	Glossary = mapper.ToGlossaryEntity(GlossaryDTO)
	return
}

// POST Glossary
func (r *HttpRepository) NewGlossary(Glossary entity.Glossary) (id string, err error) {
	body, err := json.Marshal(mapper.ToGlossaryDTO(Glossary))
	if err != nil {
		return
	}
	req, _ := http.NewRequest("POST", r.Host+"/protected/glossary/new", bytes.NewBuffer(body))
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

// GET ALL Glossary
func (r *HttpRepository) AllGlossary() (Glossarys []entity.Glossary, err error) {
	req, _ := http.NewRequest("GET", r.Host+"/protected/glossary/", nil)
	resp, err := r.doProtected(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	GlossarysDTO := make([]dto.GlossaryDTO, 0)
	if err = json.NewDecoder(resp.Body).Decode(&GlossarysDTO); err != nil {
		return
	}
	for i := range GlossarysDTO {
		Glossarys = append(Glossarys, mapper.ToGlossaryEntity(GlossarysDTO[i]))
	}
	return
}

// PUT Glossary
func (r *HttpRepository) SetGlossary(Glossary entity.Glossary) (id string, err error) {
	body, err := json.Marshal(mapper.ToGlossaryDTO(Glossary))
	if err != nil {
		return
	}
	req, _ := http.NewRequest("PUT", r.Host+"/protected/glossary/set", bytes.NewBuffer(body))
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

// DELETE Glossary
func (r *HttpRepository) DeleteGlossary(id string) (Glossary entity.Glossary, err error) {
	req, _ := http.NewRequest("DELETE", r.Host+"/protected/glossary/delete", nil)
	req.Header.Set("id", id)
	resp, err := r.doProtected(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	var GlossaryDTO dto.GlossaryDTO
	if err = json.NewDecoder(resp.Body).Decode(&GlossaryDTO); err != nil {
		return
	}
	Glossary = (mapper.ToGlossaryEntity(GlossaryDTO))
	return
}
