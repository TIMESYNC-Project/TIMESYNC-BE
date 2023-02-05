package handler

import "timesync-be/features/setting"

type EditSetting struct {
	Start       string `json:"working_hour_start" form:"working_hour_start"`
	End         string `json:"working_hour_end" form:"working_hour_end"`
	Tolerance   int    `json:"tolerance" form:"tolerance"`
	AnnualLeave int    `json:"annual_leave" form:"annual_leave"`
}

func ReqToCore(data interface{}) *setting.Core {
	res := setting.Core{}
	switch data.(type) {
	case EditSetting:
		cnv := data.(EditSetting)
		res.Start = cnv.Start
		res.End = cnv.End
		res.Tolerance = cnv.Tolerance
		res.AnnualLeave = cnv.AnnualLeave
	default:
		return nil
	}

	return &res
}
