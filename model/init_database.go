package model

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/spf13/viper"
	"github.com/tidwall/gjson"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDatabase() *gorm.DB {

	host := viper.GetString("database.host")
	if len(os.Getenv("DATABASE_HOST")) > 0 {
		host = os.Getenv("DATABASE_HOST")
	}

	username := viper.GetString("database.username")
	if len(os.Getenv("DATABASE_USERNAME")) > 0 {
		username = os.Getenv("DATABASE_USERNAME")
	}

	password := viper.GetString("database.password")
	if len(os.Getenv("DATABASE_PASSWORD")) > 0 {
		password = os.Getenv("DATABASE_PASSWORD")
	}

	databaseName := viper.GetString("database.databaseName")
	if len(os.Getenv("DATABASE_NAME")) > 0 {
		databaseName = os.Getenv("DATABASE_NAME")
	}

	post := viper.GetString("database.post")
	if len(os.Getenv("DATABASE_POST")) > 0 {
		post = os.Getenv("DATABASE_POST")
	}

	dsn := fmt.Sprintf(
		"host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=Asia/Bangkok",
		host,
		username,
		password,
		databaseName,
		post,
	)

	// dsn := "postgres://kamchai@project2021:123456789@host.docker.internal:5432/home-visit?sslmode=disable"

	// fmt.Println(dsn)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	autoMigrate(db)

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	return db
}

func autoMigrate(db *gorm.DB) {
	db.AutoMigrate(&EducationLevel{})
	initialEducationLevel(db)
	db.AutoMigrate(&User{})
	db.AutoMigrate(&StudentCustom{})
	db.AutoMigrate(&ParentStudent{})
	db.AutoMigrate(&AddressStudent{})
	db.AutoMigrate(&Setting{})
	db.AutoMigrate(&Visit{})
	db.AutoMigrate(&Degree{})
	initialDegree(db)
	db.AutoMigrate(&Faculty{})
	initialFaculty(db)
	db.AutoMigrate(&Major{})
	initialMajor(db)
	db.AutoMigrate(&VisitForm{})
}

func initialEducationLevel(db *gorm.DB) {
	var countEducationLevel EducationLevel
	countNum := db.Model(&EducationLevel{}).Find(&countEducationLevel).RowsAffected
	if countNum == 0 {
		fmt.Println(countNum)
		body, _ := ioutil.ReadFile("./configs/initialData.json")
		educationLevelJson := gjson.Get(string(body), "educationLevel")
		educationLevelData := []EducationLevel{}
		json.Unmarshal([]byte(educationLevelJson.String()), &educationLevelData)
		db.Model(&EducationLevel{}).Create(&educationLevelData)
		return
	} else {
		return
	}
}

func initialDegree(db *gorm.DB) {
	var countDegree Degree
	countNum := db.Model(&Degree{}).Find(&countDegree).RowsAffected
	if countNum == 0 {
		body, _ := ioutil.ReadFile("./configs/initialData.json")
		degreeJson := gjson.Get(string(body), "degree")
		degreeData := []Degree{}
		json.Unmarshal([]byte(degreeJson.String()), &degreeData)
		db.Model(&Degree{}).Create(&degreeData)
		return
	} else {
		return
	}
}

func initialFaculty(db *gorm.DB) {
	var countFaculty Faculty
	countNum := db.Model(&Faculty{}).Find(&countFaculty).RowsAffected
	if countNum == 0 {
		body, _ := ioutil.ReadFile("./configs/initialData.json")
		facultyJson := gjson.Get(string(body), "faculty")
		facultyData := []Faculty{}
		json.Unmarshal([]byte(facultyJson.String()), &facultyData)
		db.Model(&Faculty{}).Create(&facultyData)
		return
	} else {
		return
	}
}

func initialMajor(db *gorm.DB) {
	var countMajor Major
	countNum := db.Model(&Major{}).Find(&countMajor).RowsAffected
	if countNum == 0 {
		body, _ := ioutil.ReadFile("./configs/initialData.json")
		majorJson := gjson.Get(string(body), "major")
		majorData := []Major{}
		json.Unmarshal([]byte(majorJson.String()), &majorData)
		db.Model(&Major{}).Create(&majorData)
		return
	} else {
		return
	}
}
