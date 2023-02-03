package data

import (
	"errors"
	"log"
	"strconv"
	"timesync-be/features/user"

	"gorm.io/gorm"
)

type userQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) user.UserData {
	return &userQuery{
		db: db,
	}
}

func (uq *userQuery) Register(newUser user.Core) (user.Core, error) {
	if newUser.Email == "" || newUser.Password == "" {
		log.Println("data empty")
		return user.Core{}, errors.New("email or password not allowed empty")
	}
	dupEmail := CoreToData(newUser)
	err := uq.db.Where("email = ?", newUser.Email).First(&dupEmail).Error
	if err == nil {
		log.Println("duplicated")
		return user.Core{}, errors.New("email duplicated")
	}

	cnv := CoreToData(newUser)
	err = uq.db.Create(&cnv).Error
	if err != nil {
		log.Println("query error", err.Error())
		return user.Core{}, errors.New("server error")
	}

	newUser.ID = cnv.ID
	return newUser, nil
}

func (uq *userQuery) Login(nip string) (user.Core, error) {
	if nip == "" {
		log.Println("data empty, query error")
		return user.Core{}, errors.New("nip not allowed empty")
	}
	res := User{}
	if err := uq.db.Where("nip = ?", nip).First(&res).Error; err != nil {
		log.Println("login query error", err.Error())
		return user.Core{}, errors.New("data not found")
	}

	return ToCore(res), nil
}

func (uq *userQuery) Delete(id uint) error {
	qry := uq.db.Delete(&User{}, id)
	rowAffect := qry.RowsAffected
	if rowAffect <= 0 {
		log.Println("no data processed")
		return errors.New("no user has delete")
	}
	err := qry.Error
	if err != nil {
		log.Println("delete query error", err.Error())
		return errors.New("delete account fail")
	}
	return nil
}

// Update implements user.UserData
func (uq *userQuery) Update(employeeID uint, updateData user.Core) (user.Core, error) {
	if updateData.Email != "" {
		dupEmail := User{}
		err := uq.db.Where("email = ?", updateData.Email).First(&dupEmail).Error
		if err == nil {
			log.Println("duplicated")
			return user.Core{}, errors.New("email duplicated")
		}
	}
	cnv := CoreToData(updateData)
	qry := uq.db.Model(&User{}).Where("id = ?", employeeID).Updates(&cnv)
	affrows := qry.RowsAffected
	if affrows == 0 {
		log.Println("no rows affected")
		return user.Core{}, errors.New("no data updated")
	}
	err := qry.Error
	if err != nil {
		log.Println("update user query error", err.Error())
		return user.Core{}, err
	}
	return ToCore(cnv), nil
}

// Csv implements user.UserData
func (uq *userQuery) Csv(newUserBatch []user.Core) error {
	//prepare to make auto increment role

	batch := []User{}
	for i := 0; i < len(newUserBatch); i++ {
		batch = append(batch, CoreToData(newUserBatch[i]))
		batch[i].Role = "employee"
		batch[i].ProfilePicture = "https://cdn.pixabay.com/photo/2015/10/05/22/37/blank-profile-picture-973460_1280.png"
		nipField := []User{}
		err := uq.db.Find(&nipField).Error
		if err != nil {
			log.Println("query error")
			return errors.New("server error")
		}
		if len(nipField) == 0 {
			if len(batch) <= 1 {
				batch[i].Nip = "23001"
			} else {
				temp := batch[i-1].Nip
				// log.Println(temp)
				cnv, _ := strconv.Atoi(temp)
				cnv += 1
				batch[i].Nip = strconv.Itoa(cnv)
			}
		} else {
			if len(batch) <= 1 {
				temp := nipField[len(nipField)-1].Nip
				log.Println(temp)
				cnv, _ := strconv.Atoi(temp)
				cnv += 1
				batch[i].Nip = strconv.Itoa(cnv)
			} else {
				temp := batch[i-1].Nip
				// log.Println(temp)
				cnv, _ := strconv.Atoi(temp)
				cnv += 1
				batch[i].Nip = strconv.Itoa(cnv)
			}

		}
	}
	err := uq.db.Create(&batch).Error
	if err != nil {
		log.Println("query error")
		return errors.New("server error")
	}
	return nil
}
