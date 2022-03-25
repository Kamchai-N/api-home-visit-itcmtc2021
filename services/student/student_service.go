package services

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"

	"github.com/google/uuid"
	"github.com/kamchai-n/api-student-home-visit/domain"
	"github.com/kamchai-n/api-student-home-visit/middlewares"
	"github.com/kamchai-n/api-student-home-visit/model"
	repoSetting "github.com/kamchai-n/api-student-home-visit/repository/setting"
	repoStudent "github.com/kamchai-n/api-student-home-visit/repository/student"
	"github.com/xuri/excelize/v2"
)

type studentService struct {
	stdRepo     repoStudent.StudentRepository
	settingRepo repoSetting.SettingRepository
}

func NewStudentService(stdRepo repoStudent.StudentRepository, settingRepo repoSetting.SettingRepository) StudentService {
	return studentService{stdRepo: stdRepo, settingRepo: settingRepo}
}

func (s studentService) NewStudentFromExcel(file *multipart.FileHeader, createdBy uuid.UUID) ([]domain.ResponseCreateStudentFromExcel, error) {
	locationFile := fmt.Sprintf("./static/temp/%s", file.Filename)
	worksheet := "example_form"

	// Move File To Temp
	resFile, err := os.Create(locationFile)
	if err != nil {
		return nil, fmt.Errorf("1303")
	}
	defer resFile.Close()

	multipartFile, err := file.Open()
	if err != nil {
		return nil, fmt.Errorf("1304")
	}

	io.Copy(resFile, multipartFile)
	defer os.Remove(locationFile)

	// Read Data Form File
	f, err := excelize.OpenFile(locationFile)
	if err != nil {
		return nil, fmt.Errorf("1305")

	}

	rows, err := f.GetRows(worksheet)
	if err != nil {
		return nil, fmt.Errorf("1306")
	}

	var dataForm []domain.StudentForm
	for _, row := range rows[4:] {

		var student_level uint

		if row[5] == "ป.ตรี" || row[5] == "ตรี" || row[5] == "ปริญญาตรี" || row[5] == "49" {
			student_level = 49
		} else if row[5] == "ปวส." || row[5] == "ปวส" || row[5] == "ประกาศนียบัตรวิชาชีพชั้นสูง" || row[5] == "39" {
			student_level = 39
		} else if row[5] == "ปวช." || row[5] == "ปวช" || row[5] == "ประกาศนียบัตรวิชาชีพ" || row[5] == "29" {
			student_level = 29
		} else {
			student_level = 0
		}

		dataForm = append(dataForm, domain.StudentForm{
			GroupId:            &row[1],
			StudentId:          row[2],
			StudentName:        row[3],
			StudentPhone:       &row[4],
			StudentLevel:       &student_level,
			StudentImage:       &row[6],
			AddressNo:          &row[7],
			AddressLane:        &row[8],
			AddressRoad:        &row[9],
			AddressSubDistrict: row[10],
			AddressDistrict:    row[11],
			AddressProvince:    row[12],
			Latitude:           row[13],
			Longitude:          row[14],
			ParentName:         row[15],
			ParentPhone:        row[16],
			CreatedBy:          createdBy,
		})
	}

	// Send to database
	response, err := s.stdRepo.CreateStudentFromExcel(dataForm)
	if err != nil {
		return nil, fmt.Errorf("1302")
	}
	return response, err
}

func (s studentService) NewStudent(req domain.StudentForm) (domain.ResponseCreateStudentFromExcel, error) {
	students, err := s.stdRepo.CreateStudent(req)
	if err != nil {
		return students, fmt.Errorf("1302")
	}

	return students, nil
}

func (s studentService) ListInfoStudent(req *domain.RequestGetAll) ([]model.StudentCustom, error) {
	students, err := s.stdRepo.GetInfoStudent(req, *middlewares.UserClaims.UserId)
	if err != nil {
		return students, fmt.Errorf("1307")
	}

	return students, nil
}

func (s studentService) ListInfoStudentById(req domain.RequestGetOneStudent) (model.StudentCustom, error) {
	students, err := s.stdRepo.GetInfoStudentById(req)
	if err != nil {
		return students, fmt.Errorf("1308")
	}
	return students, nil
}

func (s studentService) RemoveStudentById(req []string) error {
	if err := s.stdRepo.DeleteStudentById(req); err != nil {
		return fmt.Errorf("1318")
	}
	return nil
}
