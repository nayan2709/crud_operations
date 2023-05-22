package service

import (
	"errors"
	"fmt"
	"github.com/dunzoit/projects/crud_operation_project/dtos"
	"github.com/dunzoit/projects/crud_operation_project/repos"
	"github.com/jinzhu/gorm"
)

type StudentService struct {
	StudentRepo repos.StudentRepo
}

func NewStudentService(db *gorm.DB) StudentService {
	return StudentService{
		StudentRepo: repos.NewStudentRepo(db),
	}
}

type StudentServiceInterface interface {
	GetStudent(id string) (*dtos.Student, error)
	CreateStudent(student dtos.AddStudentRequest) error
}

func (s *StudentService) GetStudent(id string) (*dtos.Student, error) {
	var student repos.Students
	student, err := s.StudentRepo.GetStudent(id)
	if err != nil {
		fmt.Println("error while fetching student", err)
		return nil, err
	}
	return &dtos.Student{
		Id:        student.Id,
		FirstName: student.FirstName,
		LastName:  student.LastName,
		Age:       student.Age,
	}, nil
}

func (s *StudentService) CreateStudent(student dtos.AddStudentRequest) error {
	// handle validation
	if student.FirstName == "" || student.LastName == "" || student.Age == 0 {
		return errors.New("invalid student details")
	}
	err := s.StudentRepo.CreateStudent(repos.Students{
		FirstName: student.FirstName,
		LastName:  student.LastName,
		Age:       student.Age,
	})
	if err != nil {
		fmt.Println("error while creating student", err)
		return err
	}
	return nil
}
