package data

import (
	"errors"
	"log"
	"timesync-be/features/company"

	"gorm.io/gorm"
)

type companyQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) company.CompanyData {
	return &companyQuery{
		db: db,
	}
}

// EditProfile implements company.CompanyData
func (cq *companyQuery) EditProfile(adminID uint, updateData company.Core) (company.Core, error) {
	if adminID != 1 {
		log.Println("cannot modifed admin data")
		return company.Core{}, errors.New("cannot modifed admin data")
	}
	cnv := CoreToData(updateData)
	qry := cq.db.Model(&Company{}).Where("id = ?", 1).Updates(&cnv)
	affrows := qry.RowsAffected
	if affrows == 0 {
		log.Println("no rows affected")
		return company.Core{}, errors.New("no data updated")
	}
	err := qry.Error
	if err != nil {
		log.Println("update user query error", err.Error())
		return company.Core{}, err
	}
	result := DataToCore(cnv)
	return result, nil
}

// GetProfile implements company.CompanyData
func (cq *companyQuery) GetProfile() (company.Core, error) {
	data := Company{}
	err := cq.db.Find(&data).Error
	if err != nil {
		log.Println("data not found")
		return company.Core{}, errors.New("query error, problem with server")
	}
	return DataToCore(data), nil
}
