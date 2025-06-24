package jsonDTO

type ModeDTO struct {
	Name        string   `json:"Name"`
	Mode        *int     `json:"Mode,omitempty"`
	RunTimeType *int     `json:"RunTimeType,omitempty"`
	Value       *float64 `json:"Value,omitempty"`
	Formula     string   `json:"Formula,omitempty"`
	Ability     string   `json:"Ability,omitempty"`
	AutoFail    *bool    `json:"AutoFail,omitempty"`
}

type StatusDTO struct {
	IsFavorite            bool        `json:"IsFavorite,omitempty"`
	InnerStatusCollection []StatusDTO `json:"InnerStatusCollection,omitempty"`
	SelectedModes         []ModeDTO   `json:"SelectedModes"`
	Rounds                string      `json:"Rounds,omitempty"`
	BeforeTurn            bool        `json:"BeforeTurn,omitempty"`
	Description           string      `json:"Description,omitempty"`
	TokenPicPath          string      `json:"TokenPicPath,omitempty"`
	IconPath              string      `json:"IconPath,omitempty"`
	IsDefault             bool        `json:"IsDefault,omitempty"`
	Id                    string      `json:"Id"`
	Name                  string      `json:"Name"`
}

type AbilityDTO struct {
	Name        string `json:"Name,omitempty"`
	UserValue   int    `json:"userValue"`
	MinValue    int    `json:"minValue"`
	Proficiency bool   `json:"proficiency"`
}
