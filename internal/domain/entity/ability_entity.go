package entity

type Ability struct {
	Name        string
	UserValue   int
	MinValue    int
	Proficiency bool
}

func (c Ability) GetID() string {
	return c.Name
}
