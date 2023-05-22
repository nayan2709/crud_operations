package apis

import (
	"encoding/json"
	"fmt"
	"github.com/dunzoit/projects/crud_operation_project/dtos"
	"github.com/dunzoit/projects/crud_operation_project/service"
	"io/ioutil"
	"net/http"
)

type StudentHandler struct {
	StudentService service.StudentService
}

func NewStudentHandler(studentService service.StudentService) *StudentHandler {
	return &StudentHandler{
		StudentService: studentService,
	}
}

type StudentHandlerInterface interface {
	GetStudent(w http.ResponseWriter, r *http.Request)
	CreateStudent(w http.ResponseWriter, r *http.Request)
}

func (h *StudentHandler) GetStudent(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	student, err := h.StudentService.GetStudent(id)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		fmt.Println("error while fetching student", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("{\"message\": \"error while fetching student record\"}"))
		return
	}
	data, err := json.Marshal(student)
	if err != nil {
		fmt.Println("error while marshalling student", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("{\"message\": \"error while marshaling student record\"}"))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("{\"message\": \"student data fetched successfully\", \"data\": %s}", data)))
	return
}

func (h *StudentHandler) CreateStudent(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var student dtos.AddStudentRequest
	err = json.Unmarshal(body, &student)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = h.StudentService.CreateStudent(student)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		fmt.Println("error while creating student", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("{\"message\": \"error while creating student\"}"))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("{\"message\": \"student created successfully\"}"))
	return
}
