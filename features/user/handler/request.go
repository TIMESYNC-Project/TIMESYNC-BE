package handler

import "timesync-be/features/user"

type RegisterRequest struct {
	Name        string `json:"name" form:"name"`
	BirthOfDate string `json:"birth_of_date" form:"birth_of_date"`
	Nip         string `json:"nip" form:"nip"`
	Email       string `json:"email" form:"email"`
	Gender      string `json:"gender" form:"gender"`
	Position    string `json:"position" form:"position"`
	Phone       string `json:"phone" form:"phone"`
	Address     string `json:"address" form:"address"`
	Password    string `json:"password" form:"password"`
}

type LoginRequest struct {
	Nip      string `json:"nip" form:"nip"`
	Password string `json:"password" form:"password"`
}

func ReqToCore(data interface{}) *user.Core {
	res := user.Core{}

	switch data.(type) {
	case RegisterRequest:
		cnv := data.(RegisterRequest)
		res.Name = cnv.Name
		res.Email = cnv.Email
		res.Phone = cnv.Phone
		res.Address = cnv.Address
		res.Password = cnv.Password
	case LoginRequest:
		cnv := data.(LoginRequest)
		res.Nip = cnv.Nip
		res.Password = cnv.Password
	default:
		return nil
	}

	return &res
}
