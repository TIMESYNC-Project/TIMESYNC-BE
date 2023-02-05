package data

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"
	"timesync-be/features/attendance"

	"gorm.io/gorm"
)

type attendanceQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) attendance.AttendanceData {
	return &attendanceQuery{
		db: db,
	}
}

// ClockIn implements attendance.AttendanceData
func (aq *attendanceQuery) ClockIn(employeeID uint, latitudeData string, longitudeData string) (attendance.Core, error) {
	//====================================================================
	// cari data location dan url map nya
	//====================================================================
	latitude, _ := strconv.ParseFloat(latitudeData, 64)
	longitude, _ := strconv.ParseFloat(longitudeData, 64)
	url := fmt.Sprintf("https://nominatim.openstreetmap.org/reverse?format=json&lat=%f&lon=%f&zoom=18&addressdetails=1", latitude, longitude)
	response, err := http.Get(url)
	if err != nil {
		log.Println("server error, location not found", err)
		return attendance.Core{}, errors.New("server error, location not found")
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Println("server error, location not found", err)
		return attendance.Core{}, errors.New("server error, location not found")
	}
	var nominatimResponse NominatimResponse
	err = json.Unmarshal(body, &nominatimResponse)
	if err != nil {
		log.Println("server error, location not found", err)
		return attendance.Core{}, errors.New("server error, location not found")
	}
	city := nominatimResponse.Address.City
	street := nominatimResponse.Address.Road
	postcode := nominatimResponse.Address.Postcode
	state := nominatimResponse.Address.State
	country := nominatimResponse.Address.Country
	urlLocation := fmt.Sprintf("https://www.openstreetmap.org/#map=19/%f/%f", latitude, longitude)
	loc := ""
	if len(street) == 0 {
		loc = fmt.Sprintf("%s,%s,%s,%s", city, state, postcode, country)
	} else {
		loc = fmt.Sprintf("%s %s,%s,%s,%s", street, city, state, postcode, country)
	}
	//====================================================================
	// cari Hour dan Minute kemudian convert ke string
	//====================================================================
	t := time.Now()
	hour := strconv.Itoa(t.Hour())
	minute := strconv.Itoa(t.Minute())
	year := strconv.Itoa(t.Year())
	monthConv := t.Month()
	monthInt := int(monthConv)
	month := fmt.Sprintf("%d", monthInt)

	day := strconv.Itoa(t.Day())
	if len(hour) == 1 {
		hour = "0" + hour
	}
	if len(minute) == 1 {
		minute = "0" + minute
	}
	if len(month) == 1 {
		month = "0" + month
	}
	if len(day) == 1 {
		day = "0" + day
	}
	clockInTime := fmt.Sprintf("%s:%s", hour, minute)
	clockInDate := fmt.Sprintf("%s-%s-%s", year, month, day)
	//====================================================================
	// Cek apakah user udah clockin ?
	//====================================================================
	check := Attendance{}
	err = aq.db.Where("attendance_date = ? AND user_id = ?", clockInDate, employeeID).First(&check).Error
	if err == nil {
		return attendance.Core{}, errors.New("you already clock in today")
	}

	//====================================================================
	// Sekarang tinggal diinput ke variable input
	//====================================================================
	input := Attendance{}
	input.ClockIn = clockInTime
	input.ClockInLocation = loc
	input.AttendanceDate = clockInDate
	input.ClockInOSM = urlLocation
	input.Attendance = "present"

	//====================================================================
	// Lakukan query untuk menentukan attendance status &
	// Cek apakah user udah memenuhi tepat waktu login ?
	//====================================================================
	stg := Setting{}
	err = aq.db.First(&stg).Error
	if err != nil {
		log.Println("setting at query error", err.Error())
		return attendance.Core{}, errors.New("server error, setting not found")
	}
	workStart, _ := strconv.Atoi(stg.Start[:2])
	hr, _ := strconv.Atoi(input.ClockIn[:2])
	mnt, _ := strconv.Atoi(input.ClockIn[3:])
	log.Println(hr, mnt)
	if hr < (workStart - 1) {
		log.Println("time not match")
		return attendance.Core{}, errors.New("you cannot clock in now")
	} else if hr == (workStart - 1) {
		input.AttendanceStatus = "ontime"
	} else if hr == workStart {
		if mnt > stg.Tolerance {
			input.AttendanceStatus = "late"
		} else {
			input.AttendanceStatus = "ontime"
		}
	} else {
		input.AttendanceStatus = "late"
	}
	hrEnd, _ := strconv.Atoi(stg.End[:2])
	if hr > hrEnd {
		log.Println("expired clockin time")
		return attendance.Core{}, errors.New("clockin time was expired")
	}

	input.UserId = employeeID
	err = aq.db.Create(&input).Error
	if err != nil {
		log.Println("query error, cannot insert to database", err.Error())
		return attendance.Core{}, errors.New("server error, cannot insert to database")
	}

	return DataToCore(input), nil
}

// ClockOut implements attendance.AttendanceData
func (aq *attendanceQuery) ClockOut(employeeID uint, latitudeData string, longitudeData string) (attendance.Core, error) {

	//====================================================================
	// Cek apakah user udah clockin atau clockout ?
	//====================================================================
	//inisialisasi untuk mencari time nya
	t := time.Now()
	hour := strconv.Itoa(t.Hour())
	minute := strconv.Itoa(t.Minute())
	year := strconv.Itoa(t.Year())
	monthConv := t.Month()
	monthInt := int(monthConv)
	month := fmt.Sprintf("%d", monthInt)

	day := strconv.Itoa(t.Day())
	if len(hour) == 1 {
		hour = "0" + hour
	}
	if len(minute) == 1 {
		minute = "0" + minute
	}
	if len(month) == 1 {
		month = "0" + month
	}
	if len(day) == 1 {
		day = "0" + day
	}
	clockOutTime := fmt.Sprintf("%s:%s", hour, minute)
	clockInDate := fmt.Sprintf("%s-%s-%s", year, month, day)
	//proses pengecekan apakah user sudah melakukan clock in atau belum ?
	check := Attendance{}
	err := aq.db.Where("attendance_date = ? AND user_id = ?", clockInDate, employeeID).First(&check).Error
	if err != nil {
		log.Println("query not found", err)
		return attendance.Core{}, errors.New("you dont have clock in data today,you must clock in first")
	}
	// cek apakah data clockout sudah terisi berarti user sudah melakukan clock in dan clock out
	if len(check.ClockOut) != 0 || check.ClockOut != "" {
		log.Println("query not found", err.Error())
		return attendance.Core{}, errors.New("user already clock out today")
	}

	//====================================================================
	// cari data location dan url map nya
	//====================================================================
	latitude, _ := strconv.ParseFloat(latitudeData, 64)
	longitude, _ := strconv.ParseFloat(longitudeData, 64)
	url := fmt.Sprintf("https://nominatim.openstreetmap.org/reverse?format=json&lat=%f&lon=%f&zoom=18&addressdetails=1", latitude, longitude)
	response, err := http.Get(url)
	if err != nil {
		log.Println("server error, location not found", err)
		return attendance.Core{}, errors.New("server error, location not found")
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Println("server error, location not found", err)
		return attendance.Core{}, errors.New("server error, location not found")
	}
	var nominatimResponse NominatimResponse
	err = json.Unmarshal(body, &nominatimResponse)
	if err != nil {
		log.Println("server error, location not found", err)
		return attendance.Core{}, errors.New("server error, location not found")
	}
	city := nominatimResponse.Address.City
	street := nominatimResponse.Address.Road
	postcode := nominatimResponse.Address.Postcode
	state := nominatimResponse.Address.State
	country := nominatimResponse.Address.Country
	urlLocation := fmt.Sprintf("https://www.openstreetmap.org/#map=19/%f/%f", latitude, longitude)
	loc := ""
	if len(street) == 0 {
		loc = fmt.Sprintf("%s,%s,%s,%s", city, state, postcode, country)
	} else {
		loc = fmt.Sprintf("%s %s,%s,%s,%s", street, city, state, postcode, country)
	}

	//====================================================================
	// Sekarang tinggal diinput ke variable input
	//====================================================================
	input := Attendance{}
	input.ClockOut = clockOutTime
	input.ClockOutLocation = loc
	input.ClockOutOSM = urlLocation
	// cari jumlah working time terlebih dahulu
	clockInHour, _ := strconv.Atoi(check.ClockIn[:2])
	clockOutHour, _ := strconv.Atoi(input.ClockOut[:2])
	clockInMinute, _ := strconv.Atoi(check.ClockIn[3:])
	clockOutMinute, _ := strconv.Atoi(input.ClockOut[3:])
	log.Println(clockInMinute, clockOutMinute)
	sum := clockOutHour - clockInHour
	if clockOutMinute < clockInMinute {
		sum -= 1
	}
	input.WorkTime = sum
	//====================================================================
	// Cek apakah clock out time sudah melebihi batas waktu clockout yang diberikan
	//====================================================================
	stg := Setting{}
	err = aq.db.First(&stg).Error
	if err != nil {
		log.Println("setting at query error", err.Error())
		return attendance.Core{}, errors.New("server error, setting not found")
	}
	workingHourEnd, _ := strconv.Atoi(stg.End[:2])
	if clockOutHour > (workingHourEnd + 1) {
		log.Println("time clock out not match")
		return attendance.Core{}, errors.New("clock out time expired")
	}
	//====================================================================
	// Saatnya update gaskeun
	//====================================================================
	upd := aq.db.Where("attendance_date = ? AND user_id = ?", clockInDate, employeeID).Updates(&input)
	affrows := upd.RowsAffected
	if affrows == 0 {
		log.Println("no rows affected")
		return attendance.Core{}, errors.New("no data updated")
	}
	err = upd.Error
	if err != nil {
		log.Println("update error", err.Error())
		return attendance.Core{}, errors.New("update fail, server error")
	}
	input.ClockIn = check.ClockIn
	input.ClockInLocation = check.ClockInLocation
	input.ClockInOSM = check.ClockInOSM
	input.AttendanceDate = check.AttendanceDate
	input.Attendance = check.Attendance
	input.AttendanceStatus = check.AttendanceStatus

	return DataToCore(input), nil
}

// AttendanceFromAdmin implements attendance.AttendanceData
func (aq *attendanceQuery) AttendanceFromAdmin(adminID uint, dateStart string, dateEnd string, attendanceType string, employeeID uint) error {
	if adminID != 1 {
		log.Println("user is not admin")
		return errors.New("user is not admin")
	}
	yFr, _ := strconv.Atoi(dateStart[:4])
	yTo, _ := strconv.Atoi(dateEnd[:4])
	mFr, _ := strconv.Atoi(dateStart[5:7])
	mTo, _ := strconv.Atoi(dateEnd[5:7])
	dFr, _ := strconv.Atoi(dateStart[8:])
	dTo, _ := strconv.Atoi(dateEnd[8:])
	if dTo < dFr || yTo < yFr || mTo < mFr {
		log.Println("wrong input format")
		return errors.New("wrong input format")
	}
	isfalse := true
	y, _ := strconv.Atoi(dateStart[:4])
	mm, _ := strconv.Atoi(dateStart[5:7])
	d, _ := strconv.Atoi(dateStart[8:])
	m := time.Month(mm)
	// log.Println(y, m, d)
	fD := time.Date(y, m, d, 0, 0, 0, 0, time.UTC)
	x := 1
	data := []Attendance{}
	date := dateStart
	isfalse = true
	i := 0
	for isfalse {
		createAt := date
		temp := Attendance{}
		data = append(data, temp)
		data[i].AttendanceDate = createAt
		data[i].Attendance = attendanceType
		data[i].UserId = employeeID
		fmt.Println(createAt)
		if createAt == dateEnd {
			isfalse = false
		}
		tomorrow := fD.AddDate(0, 0, x)
		year := strconv.Itoa(tomorrow.Year())
		monthCnv := int(tomorrow.Month())
		month := strconv.Itoa(monthCnv)
		day := strconv.Itoa(tomorrow.Day())
		if len(month) == 1 {
			month = "0" + month
		}
		if len(day) == 1 {
			day = "0" + day
		}
		date = fmt.Sprintf("%s-%s-%s", year, month, day)
		x++
		i++
	}
	err := aq.db.Create(&data).Error
	if err != nil {
		log.Println("creating data error", err.Error())
		return errors.New("creating data fail, server error")
	}
	return nil
}

// Record implements attendance.AttendanceData
func (aq *attendanceQuery) Record(employeeID uint, dateFrom string, dateTo string) ([]attendance.Core, error) {
	yFr, _ := strconv.Atoi(dateFrom[:4])
	yTo, _ := strconv.Atoi(dateTo[:4])
	mFr, _ := strconv.Atoi(dateFrom[5:7])
	mTo, _ := strconv.Atoi(dateTo[5:7])
	dFr, _ := strconv.Atoi(dateFrom[8:])
	dTo, _ := strconv.Atoi(dateTo[8:])
	log.Println(yFr, yTo)
	if dTo < dFr || yTo < yFr || mTo < mFr {
		log.Println("wrong input format")
		return []attendance.Core{}, errors.New("wrong input format")
	}
	data := []Attendance{}
	err := aq.db.Where("user_id = ?", employeeID).Find(&data).Error //.Order("attendance_date desc")
	if err != nil {
		log.Println("query error data not found", err.Error())
		return []attendance.Core{}, errors.New("data not found")
	}
	result := []attendance.Core{}
	for i := 0; i < len(data); i++ {
		result = append(result, DataToCore(data[i]))
	}
	//Cek apakah result
	date := dateFrom
	isfalse := true
	y, _ := strconv.Atoi(date[:4])
	mm, _ := strconv.Atoi(date[5:7])
	d, _ := strconv.Atoi(date[8:])
	m := time.Month(mm)
	// log.Println(y, m, d)
	fD := time.Date(y, m, d, 0, 0, 0, 0, time.UTC)
	x := 1
	i := 0
	val := 0
	response := []attendance.Core{}
	for isfalse {

		createAt := date

		// if response[i].AttendanceDate == dateTo {
		// 	isfalse = false
		// }
		log.Println(result[val].AttendanceDate, createAt)
		crtInt, _ := strconv.Atoi(createAt[8:])
		rs, _ := strconv.Atoi(result[val].AttendanceDate[8:])
		log.Println(rs, crtInt)
		if val > 0 {
			if rs > crtInt {
				val -= 1
			}
		}
		if rs >= crtInt || result[val].AttendanceDate == createAt {
			if createAt != result[val].AttendanceDate {
				crtInt, _ := strconv.Atoi(createAt[8:])
				rs, _ := strconv.Atoi(result[val].AttendanceDate[8:])
				log.Println(rs, crtInt)
				if rs > crtInt && result[val].AttendanceDate == result[0].AttendanceDate {
					temp := attendance.Core{}
					//cari data attendance date
					if len(response) == 0 {
						temp.AttendanceDate = date
					} else {
						conv, _ := strconv.Atoi(response[len(response)-1].AttendanceDate[8:])
						conv += 1
						cnvS := strconv.Itoa(conv)
						if len(cnvS) == 1 {
							cnvS = "0" + cnvS
						}
						final := (response[len(response)-1].AttendanceDate[:8]) + cnvS
						temp.AttendanceDate = final
					}
					temp.Attendance = "NO DATA"
					response = append(response, temp)
				} else {
					temp := attendance.Core{}
					//cari data attendance date
					if len(response) == 0 {
						temp.AttendanceDate = date
					} else {
						conv, _ := strconv.Atoi(response[len(response)-1].AttendanceDate[8:])
						conv += 1
						cnvS := strconv.Itoa(conv)
						if len(cnvS) == 1 {
							cnvS = "0" + cnvS
						}
						final := (response[len(response)-1].AttendanceDate[:8]) + cnvS
						temp.AttendanceDate = final
					}
					temp.Attendance = "Absent"
					response = append(response, temp)
				}
			} else {
				response = append(response, result[val])
			}

			if createAt == dateTo {
				isfalse = false
			}
			tomorrow := fD.AddDate(0, 0, x)
			year := strconv.Itoa(tomorrow.Year())
			monthCnv := int(tomorrow.Month())
			month := strconv.Itoa(monthCnv)
			day := strconv.Itoa(tomorrow.Day())
			if len(month) == 1 {
				month = "0" + month
			}
			if len(day) == 1 {
				day = "0" + day
			}
			date = fmt.Sprintf("%s-%s-%s", year, month, day)
			x++
			i++
		}
		if result[val].AttendanceDate == result[len(result)-1].AttendanceDate {
			log.Println("data reach limit")
			break
			// return []attendance.Core{}, errors.New("data reach limit")
		}
		val++
	}
	// log.Println(response)

	return response, nil
}
