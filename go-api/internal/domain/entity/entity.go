package entity

type Identifiable interface {
	GetID() ID
}

type ID string

func (id ID) GetID() ID {
	return id
}

func (id ID) String() string {
	return string(id)
}
