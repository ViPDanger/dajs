package entity

type Glossary struct {
	ID   string `json:"Id"`
	Text string `json:"Text"`
}

func (c Glossary) GetID() string {
	return c.ID
}
