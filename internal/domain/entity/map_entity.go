package entity

type Map struct {
	ID                    string         `json:"Id"`
	Path                  string         `json:"Path"`
	Hold                  bool           `json:"Hold"`
	HoldTime              string         `json:"HoldTime"`
	X                     float64        `json:"X"`
	Y                     float64        `json:"Y"`
	CurrentLayer          int            `json:"CurrentLayer"`
	UserMapSizePercent    int            `json:"UserMapSizePercent"`
	Xmirror               float64        `json:"Xmirror"`
	Ymirror               float64        `json:"Ymirror"`
	RotateMap             int            `json:"RotateMap"`
	FogOfWarPlayerShow    bool           `json:"FogOfWarPlayerShow"`
	VideoVolumeInt        int            `json:"VideoVolumeInt"`
	LastWallId            int            `json:"LastWallId"`
	LockMap               bool           `json:"LockMap"`
	DeepColor             bool           `json:"DeepColor"`
	PlayersFogTranperensy float64        `json:"PlayersFogTranperensy"`
	MasterFogOpacity      float64        `json:"MasterFogOpacity"`
	Fog                   []MapFog       `json:"Fog"`
	Monsters              []MapMonster   `json:"Monsters"`
	Players               []MapPlayer    `json:"Players"`
	Npc                   []MapNpc       `json:"Npc"`
	Props                 []MapProp      `json:"Props"`
	PropsList             []MapPropsList `json:"PropsList"`
}

func (c Map) GetID() string {
	return c.ID
}

type MapFog struct {
	Name               string `json:"Name"`
	ImLight            bool   `json:"ImLight"`
	Layer              int    `json:"Layer"`
	ImLightHalf        bool   `json:"ImLightHalf"`
	ShowFigureToPlayer bool   `json:"ShowFigureToPlayer"`
	ImEraser           bool   `json:"ImEraser"`
	FogDataType        string `json:"FogDataType"`
	Rect               string `json:"Rect"`
}

type MapMonster struct {
	ID                   string      `json:"Id"`
	Name                 string      `json:"Name"`
	X                    float64     `json:"X"`
	Y                    float64     `json:"Y"`
	ShowFigureToPlayer   bool        `json:"ShowFigureToPlayer"`
	EyesightEnabled      bool        `json:"EyesightEnabled"`
	CellsOfEyesight      int         `json:"CellsOfEyesight"`
	IHaveLight           bool        `json:"IHaveLight"`
	TorchValue           int         `json:"TorchValue"`
	TorchValueSecond     int         `json:"TorchValueSecond"`
	HitPoints            int         `json:"HitPoints"`
	CurrentHitPoints     int         `json:"CurrentHitPoints"`
	TempHitPoints        int         `json:"TempHitPoints"`
	TempCurrentHitPoints int         `json:"TempCurrentHitPoints"`
	IniColor             string      `json:"IniColor"`
	LegendaryActions     string      `json:"LegendaryActions"`
	CustomStatuses       []MapStatus `json:"CustomStatuses"`
}

type MapStatus struct {
	ID           string `json:"Id"`
	Name         string `json:"Name"`
	Description  string `json:"Description"`
	IconPath     string `json:"IconPath"`
	TokenPicPath string `json:"TokenPicPath"`
	IsFavorite   bool   `json:"IsFavorite"`
}

func (c MapStatus) GetID() string {
	return c.ID
}

type MapPlayer struct {
	ID                 string  `json:"Id"`
	X                  float64 `json:"X"`
	Y                  float64 `json:"Y"`
	ShowFigureToPlayer bool    `json:"ShowFigureToPlayer"`
	EyesightEnabled    bool    `json:"EyesightEnabled"`
	CellsOfEyesight    int     `json:"CellsOfEyesight"`
	TorchValue         int     `json:"TorchValue"`
	TorchValueSecond   int     `json:"TorchValueSecond"`
	HasAura            bool    `json:"HasAura,omitempty"`
	AuraColor          string  `json:"AuraColor,omitempty"`
	AuraOpacity        float64 `json:"AuraOpacity,omitempty"`
}

func (c MapPlayer) GetID() string {
	return c.ID
}

type MapNpc struct {
	ID                 string  `json:"Id"`
	X                  float64 `json:"X"`
	Y                  float64 `json:"Y"`
	ShowFigureToPlayer bool    `json:"ShowFigureToPlayer"`
	EyesightEnabled    bool    `json:"EyesightEnabled"`
	CellsOfEyesight    int     `json:"CellsOfEyesight"`
}

func (c MapNpc) GetID() string {
	return c.ID
}

type MapProp struct {
	ID                 string  `json:"Id"`
	Name               string  `json:"Name"`
	Uri                string  `json:"Uri"`
	X                  float64 `json:"X"`
	Y                  float64 `json:"Y"`
	Width              float64 `json:"Width"`
	Height             float64 `json:"Height"`
	ZindexElement      int     `json:"ZindexElement"`
	ShowToMaster       bool    `json:"ShowToMaster"`
	ShowFigureToPlayer bool    `json:"ShowFigureToPlayer"`
	Note               string  `json:"Note"`
	LinkedMapId        string  `json:"LinkedMapId,omitempty"`
	Text               string  `json:"Text,omitempty"`
	AudioPath          string  `json:"AudioPath,omitempty"`
}

func (c MapProp) GetID() string {
	return c.ID
}

type MapPropsList struct {
	Name  string    `json:"Name"`
	Props []MapProp `json:"Props"`
}
