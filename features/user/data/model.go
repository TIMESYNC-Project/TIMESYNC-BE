package data

import (
	"timesync-be/features/user"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ProfilePicture string
	Name           string
	BirthOfDate    string
	Nip            string `gorm:"not null"`
	Email          string `gorm:"unique"`
	Gender         string
	Position       string
	Phone          string
	Address        string
	Password       string
	Role           string
	AnnualLeave    int
}

type Setting struct {
	gorm.Model
	Start       string
	End         string
	Tolerance   int
	AnnualLeave int
}

func ToCore(data User) user.Core {
	return user.Core{
		ID:             data.ID,
		ProfilePicture: data.ProfilePicture,
		Name:           data.Name,
		BirthOfDate:    data.BirthOfDate,
		Nip:            data.Nip,
		Email:          data.Email,
		Gender:         data.Gender,
		Position:       data.Position,
		Phone:          data.Phone,
		Address:        data.Address,
		Password:       data.Password,
		Role:           data.Role,
		AnnualLeave:    data.AnnualLeave,
	}
}

func CoreToData(data user.Core) User {
	return User{
		Model:          gorm.Model{ID: data.ID},
		ProfilePicture: data.ProfilePicture,
		Name:           data.Name,
		BirthOfDate:    data.BirthOfDate,
		Nip:            data.Nip,
		Email:          data.Email,
		Gender:         data.Gender,
		Position:       data.Position,
		Phone:          data.Phone,
		Address:        data.Address,
		Password:       data.Password,
		Role:           data.Role,
		AnnualLeave:    data.AnnualLeave,
	}
}
