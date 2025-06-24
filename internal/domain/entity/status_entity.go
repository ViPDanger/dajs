package entity

type Status struct {
	ID                    string
	Name                  string
	IsHided               bool
	InnerStatusCollection []interface{}
	SelectedModes         []StatusMode
	Description           string
	TokenPicPath          string
	IconPath              string
}

func (c Status) GetID() string {
	return c.ID
}

type StatusMode struct {
	Name        string
	Mode        int
	RunTimeType int
	Value       float64
	Formula     string
	Ability     string
}
