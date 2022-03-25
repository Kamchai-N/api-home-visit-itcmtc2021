package repository

import (
	"github.com/google/uuid"
	"github.com/kamchai-n/api-student-home-visit/domain"
	"github.com/kamchai-n/api-student-home-visit/model"
	"gorm.io/gorm"
)

type studentRepository struct {
	db *gorm.DB
}

func NewStudentRepository(db *gorm.DB) StudentRepository {
	return studentRepository{db: db}
}

func (s studentRepository) CreateStudentFromExcel(stdForm []domain.StudentForm) ([]domain.ResponseCreateStudentFromExcel, error) {
	tx := s.db.Begin()

	var parentStdAll []model.ParentStudent
	var addressStdAll []model.AddressStudent
	var response []domain.ResponseCreateStudentFromExcel
	for _, dataStd := range stdForm {
		studentCustom := model.StudentCustom{
			GroupId:      dataStd.GroupId,
			StudentId:    dataStd.StudentId,
			StudentName:  dataStd.StudentName,
			StudentPhone: dataStd.StudentPhone,
			StudentLevel: dataStd.StudentLevel,
			StudentImage: dataStd.StudentImage,
			CreatedBy:    dataStd.CreatedBy,
		}
		if err := tx.Model(&model.StudentCustom{}).Create(&studentCustom).Error; err != nil {
			tx.Rollback()
			return response, err
		}

		addressStd := model.AddressStudent{
			AddressNo:          dataStd.AddressNo,
			AddressLane:        dataStd.AddressLane,
			AddressRoad:        dataStd.AddressRoad,
			AddressSubDistrict: dataStd.AddressSubDistrict,
			AddressDistrict:    dataStd.AddressDistrict,
			Latitude:           dataStd.Latitude,
			Longitude:          dataStd.Longitude,
			AddressProvince:    dataStd.AddressProvince,
			StudentCustomID:    studentCustom.ID,
		}
		addressStdAll = append(addressStdAll, addressStd)

		parentStd := model.ParentStudent{
			ParentName:      dataStd.ParentName,
			ParentPhone:     *dataStd.StudentPhone,
			StudentCustomID: studentCustom.ID,
		}
		parentStdAll = append(parentStdAll, parentStd)

		response = append(response, domain.ResponseCreateStudentFromExcel{
			StudentCustom:  studentCustom,
			AddressStudent: addressStd,
			ParentStudent:  parentStd,
		})
	}
	if err := tx.Model(&model.ParentStudent{}).Create(&parentStdAll).Error; err != nil {
		tx.Rollback()
		return response, err
	}
	if err := tx.Model(&model.AddressStudent{}).Create(&addressStdAll).Error; err != nil {
		tx.Rollback()
		return response, err
	}

	tx.Commit()
	return response, nil
}

func (s studentRepository) CreateStudent(stdForm domain.StudentForm) (domain.ResponseCreateStudentFromExcel, error) {
	tx := s.db.Begin()

	var parentStdAll []model.ParentStudent
	var addressStdAll []model.AddressStudent
	var response domain.ResponseCreateStudentFromExcel
	studentCustom := model.StudentCustom{
		GroupId:      stdForm.GroupId,
		StudentId:    stdForm.StudentId,
		StudentName:  stdForm.StudentName,
		StudentPhone: stdForm.StudentPhone,
		StudentLevel: stdForm.StudentLevel,
		StudentImage: stdForm.StudentImage,
		CreatedBy:    stdForm.CreatedBy,
	}
	if err := tx.Model(&model.StudentCustom{}).Create(&studentCustom).Error; err != nil {
		tx.Rollback()
		return response, err
	}

	addressStd := model.AddressStudent{
		AddressNo:          stdForm.AddressNo,
		AddressLane:        stdForm.AddressLane,
		AddressRoad:        stdForm.AddressRoad,
		AddressSubDistrict: stdForm.AddressSubDistrict,
		AddressDistrict:    stdForm.AddressDistrict,
		Latitude:           stdForm.Latitude,
		Longitude:          stdForm.Longitude,
		AddressProvince:    stdForm.AddressProvince,
		StudentCustomID:    studentCustom.ID,
	}
	addressStdAll = append(addressStdAll, addressStd)

	parentStd := model.ParentStudent{
		ParentName:      stdForm.ParentName,
		ParentPhone:     *stdForm.StudentPhone,
		StudentCustomID: studentCustom.ID,
	}
	parentStdAll = append(parentStdAll, parentStd)

	response = domain.ResponseCreateStudentFromExcel{
		StudentCustom:  studentCustom,
		AddressStudent: addressStd,
		ParentStudent:  parentStd,
	}

	if err := tx.Model(&model.ParentStudent{}).Create(&parentStdAll).Error; err != nil {
		tx.Rollback()
		return response, err
	}
	if err := tx.Model(&model.AddressStudent{}).Create(&addressStdAll).Error; err != nil {
		tx.Rollback()
		return response, err
	}

	tx.Commit()
	return response, nil
}

func (s studentRepository) GetInfoStudent(req *domain.RequestGetAll, user_id uuid.UUID) ([]model.StudentCustom, error) {
	var studentCustom []model.StudentCustom

	tx := s.db.Preload("AddressStudent").Preload("ParentStudent")
	if req != nil {
		if req.GroupId != nil {
			tx.Where("group_id LIKE '%" + *req.GroupId + "%'")
		}

		if req.StudentId != nil {
			tx.Where("student_id LIKE '%" + *req.StudentId + "%'")
		}

		if len(*req.StudentLevel) > 0 {
			tx.Where("student_level IN (?)", *req.StudentLevel)
		}

		if req.AddressNo != nil {
			tx.Where("id IN (SELECT student_custom_id FROM address_students WHERE address_no LIKE '%" + *req.AddressNo + "%')")
		}

		if req.AddressSubDistrict != nil {
			tx.Where("id IN (SELECT student_custom_id FROM address_students WHERE address_sub_district LIKE '%" + *req.AddressSubDistrict + "%')")
		}

		if req.AddressDistrict != nil {
			tx.Where("id IN (SELECT student_custom_id FROM address_students WHERE address_district LIKE '%" + *req.AddressDistrict + "%')")
		}

		if req.AddressProvince != nil {
			tx.Where("id IN (SELECT student_custom_id FROM address_students WHERE address_province LIKE '%" + *req.AddressProvince + "%')")
		}
	}
	if err := tx.Preload("EducationLevel").Where("created_by = ?", user_id).Find(&studentCustom).Error; err != nil {
		return studentCustom, err
	}

	return studentCustom, nil
}

func (s studentRepository) GetInfoStudentById(req domain.RequestGetOneStudent) (model.StudentCustom, error) {
	var studentCustom model.StudentCustom

	tx := s.db.Preload("AddressStudent").Preload("ParentStudent")

	if req.ID != nil {
		tx.Where("id = ?", req.ID)
	}

	if req.StudentId != nil {
		tx.Where("student_id = ?", req.StudentId)
	}

	if err := tx.First(&studentCustom).Error; err != nil {
		return studentCustom, err
	}

	return studentCustom, nil
}

func (s studentRepository) DeleteStudentById(id []string) error {
	if err := s.db.Where("id IN (?)", id).Delete(&model.StudentCustom{}).Error; err != nil {
		return err
	}
	return nil
}
