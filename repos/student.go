package repos

import (
	"github.com/jinzhu/gorm"
)

type Students struct {
	Id        string
	FirstName string
	LastName  string
	Age       int
}

type StudentRepo struct {
	DB *gorm.DB
}

func NewStudentRepo(db *gorm.DB) StudentRepo {
	return StudentRepo{
		DB: db,
	}
}

type StudentRepoInterface interface {
	GetStudent(id string) (Students, error)
	CreateStudent(student Students) error
}

func (s *StudentRepo) GetStudent(id string) (Students, error) {
	var student Students
	err := s.DB.Table("students").Where("id = ?", id).Find(&student).Error
	return student, err
}

func (s *StudentRepo) CreateStudent(student Students) error {
	err := s.DB.Table("students").Create(&student).Error
	return err
}
