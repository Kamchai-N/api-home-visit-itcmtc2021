package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type StudentCustom struct {
	ID           uuid.UUID `gorm:"primaryKey; type:uuid; default:uuid_generate_v4();" json:"id"`
	GroupId      *string   `json:"group_id"`
	StudentId    string    `gorm:"not null" json:"student_id"`
	StudentName  string    `gorm:"not null" json:"student_name"`
	StudentPhone *string   `gorm:"null" json:"student_phone"`
	StudentLevel *uint     `gorm:"not null" json:"student_level"`
	StudentImage *string   `gorm:"null" json:"student_image"`
	CreatedBy    uuid.UUID `gorm:"not null" json:"created_by"`

	User           *User           `gorm:"foreignKey:UserID;references:CreatedBy"  json:"user,omitempty"`
	AddressStudent *AddressStudent `gorm:"foreignKey:StudentCustomID" json:"address,omitempty"`
	ParentStudent  *ParentStudent  `gorm:"foreignKey:StudentCustomID" json:"parent,omitempty"`
	EducationLevel *EducationLevel `gorm:"foreignKey:EnducationLevelId;references:StudentLevel" json:"education_level,omitempty"`

	CreatedAt *time.Time     `gorm:"default:now()" json:"-"`
	UpdatedAt *time.Time     `gorm:"default:now()" json:"-"`
	DeleteAt  gorm.DeletedAt `json:"-"`
}
