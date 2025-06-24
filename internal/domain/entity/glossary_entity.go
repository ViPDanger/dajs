package entity

type Glossary struct {
	ID   string
	Text string
}

func (c Glossary) GetID() string {
	return c.ID
}
