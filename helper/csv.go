package helper

import (
	"timesync-be/features/user"
)

type CsvRequest struct {
	Name        string `json:"name" form:"name"`
	BirthOfDate string `json:"birth_of_date" form:"birth_of_date"`
	Email       string `json:"email" form:"email"`
	Gender      string `json:"gender" form:"gender"`
	Position    string `json:"position" form:"position"`
	Phone       string `json:"phone" form:"phone"`
	Address     string `json:"address" form:"address"`
	Password    string `json:"password" form:"password"`
}

func ConvertToCore(data CsvRequest) user.Core {
	return user.Core{
		Name:        data.Name,
		BirthOfDate: data.BirthOfDate,
		Email:       data.Email,
		Gender:      data.Gender,
		Position:    data.Position,
		Phone:       data.Phone,
		Address:     data.Address,
		Password:    data.Password,
	}
}

func ConvertCSV(data [][]string) []user.Core {

	res := []CsvRequest{}
	for i := 1; i < len(data); i++ {
		//inisialisasi data pertama kosong
		init := CsvRequest{}
		res = append(res, init)
		// log.Println(data[i])
		// pindahkan isi csv ke slice
		for j := 0; j < len(data[0]); j++ {
			if j == 0 {
				res[i-1].Name = data[i][j]
			}
			if j == 1 {
				res[i-1].BirthOfDate = data[i][j]
			}
			if j == 2 {
				res[i-1].Email = data[i][j]
			}
			if j == 3 {
				res[i-1].Gender = data[i][j]
			}
			if j == 4 {
				res[i-1].Position = data[i][j]
			}
			if j == 5 {
				res[i-1].Phone = data[i][j]
			}
			if j == 6 {
				res[i-1].Address = data[i][j]
			}
			if j == 7 {
				res[i-1].Password = data[i][j]
			}
		}
	}
	result := []user.Core{}
	for i := 0; i < len(res); i++ {
		result = append(result, ConvertToCore(res[i]))
	}
	return result
}
