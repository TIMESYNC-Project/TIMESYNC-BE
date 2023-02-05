package data

import (
	"timesync-be/features/setting"

	"gorm.io/gorm"
)

type Setting struct {
	gorm.Model
	Start       string
	End         string
	Tolerance   int
	AnnualLeave int
}

func DataToCore(data Setting) setting.Core {
	return setting.Core{
		ID:          data.ID,
		Start:       data.Start,
		End:         data.End,
		Tolerance:   data.Tolerance,
		AnnualLeave: data.AnnualLeave,
	}
}

func CoreToData(core setting.Core) Setting {
	return Setting{
		Model:       gorm.Model{ID: core.ID},
		Start:       core.Start,
		End:         core.End,
		Tolerance:   core.Tolerance,
		AnnualLeave: core.AnnualLeave,
	}
}
