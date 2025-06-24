package dto

type MapDTO struct {
	ID                    string            `json:"Id"`
	Path                  string            `json:"Path"`
	Hold                  bool              `json:"Hold"`
	HoldTime              string            `json:"HoldTime"`
	X                     float64           `json:"X"`
	Y                     float64           `json:"Y"`
	CurrentLayer          int               `json:"CurrentLayer"`
	IsTeleportIsVisible   bool              `json:"IsTeleportIsVisible"`
	IsZoomerMap           bool              `json:"IsZoomerMap"`
	GridType              int               `json:"GridType"`
	UseGridMask           bool              `json:"UseGridMask"`
	Snap                  bool              `json:"Snap"`
	GridSize              float64           `json:"GridSize"`
	GridThickness         float64           `json:"GridThickness"`
	GridOffsetX           float64           `json:"GridOffsetX"`
	GridOffsetY           float64           `json:"GridOffsetY"`
	ColorGrid             string            `json:"ColorGrid"`
	MapColor              string            `json:"MapColor"`
	UserMapSizePercent    int               `json:"UserMapSizePercent"`
	Xmirror               float64           `json:"Xmirror"`
	Ymirror               float64           `json:"Ymirror"`
	RotateMap             int               `json:"RotateMap"`
	FogOfWarPlayerShow    bool              `json:"FogOfWarPlayerShow"`
	VideoVolumeInt        int               `json:"VideoVolumeInt"`
	LastWallId            int               `json:"LastWallId"`
	LockMap               bool              `json:"LockMap"`
	DeepColor             bool              `json:"DeepColor"`
	PlayersFogTranperensy float64           `json:"PlayersFogTranperensy"`
	MasterFogOpacity      float64           `json:"MasterFogOpacity"`
	GlobalEyeEnabled      bool              `json:"GlobalEyeEnabled"`
	MatrixForPlayers      string            `json:"MatrixForPlayers"`
	SelectedPattern       *string           `json:"SelectedPattern"`
	Layers                []interface{}     `json:"Layers"`
	Fog                   []MapFogDTO       `json:"Fog"`
	GridMask              []interface{}     `json:"GridMask"`
	InNodes               []string          `json:"InNodes"`
	Monsters              []MapMonsterDTO   `json:"Monsters"`
	Players               []MapPlayerDTO    `json:"Players"`
	Npc                   []MapNpcDTO       `json:"Npc"`
	Figures               []interface{}     `json:"Figures"`
	Props                 []MapPropDTO      `json:"Props"`
	PropsList             []MapPropsListDTO `json:"PropsList"`
}

type MapFogDTO struct {
	Name               string `json:"Name"`
	ImLight            bool   `json:"ImLight"`
	Layer              int    `json:"Layer"`
	ImLightHalf        bool   `json:"ImLightHalf"`
	ShowFigureToPlayer bool   `json:"ShowFigureToPlayer"`
	ImEraser           bool   `json:"ImEraser"`
	FogDataType        string `json:"FogDataType"`
	Rect               string `json:"Rect"`
}

type MapMonsterDTO struct {
	ID                   string         `json:"Id"`
	Name                 string         `json:"Name"`
	X                    float64        `json:"X"`
	Y                    float64        `json:"Y"`
	ShowFigureToPlayer   bool           `json:"ShowFigureToPlayer"`
	EyesightEnabled      bool           `json:"EyesightEnabled"`
	CellsOfEyesight      int            `json:"CellsOfEyesight"`
	IHaveLight           bool           `json:"IHaveLight"`
	TorchValue           int            `json:"TorchValue"`
	TorchValueSecond     int            `json:"TorchValueSecond"`
	HitPoints            int            `json:"HitPoints"`
	CurrentHitPoints     int            `json:"CurrentHitPoints"`
	TempHitPoints        int            `json:"TempHitPoints"`
	TempCurrentHitPoints int            `json:"TempCurrentHitPoints"`
	IniColor             string         `json:"IniColor"`
	LegendaryActions     string         `json:"LegendaryActions"`
	CustomStatuses       []MapStatusDTO `json:"CustomStatuses"`
}

type MapStatusDTO struct {
	ID           string `json:"Id"`
	Name         string `json:"Name"`
	Description  string `json:"Description"`
	IconPath     string `json:"IconPath"`
	TokenPicPath string `json:"TokenPicPath"`
	IsFavorite   bool   `json:"IsFavorite"`
}

type MapPlayerDTO struct {
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

type MapNpcDTO struct {
	ID                 string  `json:"Id"`
	X                  float64 `json:"X"`
	Y                  float64 `json:"Y"`
	ShowFigureToPlayer bool    `json:"ShowFigureToPlayer"`
	EyesightEnabled    bool    `json:"EyesightEnabled"`
	CellsOfEyesight    int     `json:"CellsOfEyesight"`
}

type MapPropDTO struct {
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

type MapPropsListDTO struct {
	Name  string       `json:"Name"`
	Props []MapPropDTO `json:"Props"`
}
