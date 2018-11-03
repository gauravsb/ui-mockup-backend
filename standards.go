package root

type Controls struct {
	ControlName string `json:"controlName"`
	ControlInfo ControlInfo `json:"controlInfo"`
}

type ControlInfo struct {
	Family          string  `json:"family"`
	Name     		string  `json:"name"`
	Description     string  `json:"desc"`
}

type Standard struct {
	StandardName string `json:"standardName"`
	Controls[] Controls `json:"controls"`
}

type StandardService interface {
	CreateStandard(u *Standard) error
	GetStandardInfo(standardname string) (error, Standard)
}

