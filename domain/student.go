package domain

import (
	"github.com/google/uuid"
	"github.com/kamchai-n/api-student-home-visit/model"
)

type StudentForm struct {
	GroupId            *string   `json:"group_id"`
	StudentId          string    `json:"student_id"`
	StudentName        string    `json:"student_name"`
	StudentPhone       *string   `json:"student_phone"`
	StudentLevel       *uint     `json:"student_level"`
	StudentImage       *string   `json:"student_image"`
	AddressNo          *string   `json:"address_no"`
	AddressLane        *string   `json:"address_lane"`
	AddressRoad        *string   `json:"address_road"`
	AddressSubDistrict string    `json:"address_sub_district"`
	AddressDistrict    string    `json:"address_district"`
	AddressProvince    string    `json:"address_province"`
	Latitude           string    `json:"latitude"`
	Longitude          string    `json:"longitude"`
	ParentName         string    `json:"parent_name"`
	ParentPhone        string    `json:"parent_phone"`
	CreatedBy          uuid.UUID `json:"created_by"`
}

type ResponseCreateStudentFromExcel struct {
	StudentCustom  model.StudentCustom  `json:"studentCustom"`
	ParentStudent  model.ParentStudent  `json:"parentStudent"`
	AddressStudent model.AddressStudent `json:"addressStudent"`
}

type RequestGetOneStudent struct {
	ID        *uuid.UUID `json:"id"`
	StudentId *string    `json:"student_id"`
}

type RequestDeleteStudent struct {
	UserID []string `json:"id"`
}

type RequestGetAll struct {
	GroupId            *string `json:"group_id"`
	StudentId          *string `json:"student_id"`
	StudentLevel       *[]uint `json:"student_level"`
	AddressNo          *string `json:"address_no"`
	AddressSubDistrict *string `json:"address_sub_district"`
	AddressDistrict    *string `json:"address_district"`
	AddressProvince    *string `json:"address_province"`
}
