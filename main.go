package main

import (
	"fmt"
	"github.com/dunzoit/projects/crud_operation_project/apis"
	"github.com/dunzoit/projects/crud_operation_project/database_op"
	"github.com/dunzoit/projects/crud_operation_project/service"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	postges := database_op.NewDatabase("localhost", 5432, "postgres", "postgres", "student_service_db")
	db, err := postges.InitializePostgres()
	if err != nil {
		panic(err)
	}
	studentService := service.NewStudentService(db)
	studentHandler := apis.NewStudentHandler(studentService)

	//curl -X GET '0.0.0.0:8080/student?id=33'
	router.HandleFunc("/student", studentHandler.GetStudent).Methods("GET")
	//curl -X POST '0.0.0.0:8080/student' --header 'Content-Type: application/json' -d '{"first_name":"test","last_name":"test","age":33}'
	router.HandleFunc("/student", studentHandler.CreateStudent).Methods("POST")
	svr := http.Server{
		Addr: "0.0.0.0:8080",
	}
	err = http.ListenAndServe(svr.Addr, router)
	if err != nil {
		panic(err)
	}
	// To avoid panic when server is closed
	defer func() {
		if recover() != nil {
			fmt.Println("panic server closed")
			return
		}
	}()
}
