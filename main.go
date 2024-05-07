package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/swaggo/http-swagger"
	"go.etcd.io/bbolt"
	"log"
	"net/http"
	"template-golang/internal/app/department"
	"template-golang/internal/app/employee"
	repositoryDept "template-golang/internal/repository/department"
	repositoryEmployee "template-golang/internal/repository/employee"
)

// @title Clean GO API Docs
// @version 1.0.0
// @contact.name Renato Cerqueira Torres
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:port
// @BasePath /
func main() {
	db, err := bbolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *bbolt.DB) {
		err := db.Close()
		if err != nil {
			fmt.Println("Error:", err)
		}
	}(db)

	repo := repositoryEmployee.NewBoltRepository(db)
	service := employee.NewService(repo)
	handler := employee.NewHandler(service)

	r := mux.NewRouter()
	r.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)

	//Employees
	r.HandleFunc(employee.Employees.ByID, handler.GetAllEmployees).Methods("GET")
	r.HandleFunc(employee.Employees.ByID, handler.GetEmployeeById).Methods("GET")
	r.HandleFunc(employee.Employees.ByDepartment, handler.GetAllEmployeesByDepartmentID).Methods("GET")
	r.HandleFunc(employee.Employees.Base, handler.CreateEmployee).Methods("POST")
	r.HandleFunc(employee.Employees.ByID, handler.UpdateEmployeeByID).Methods("PUT")
	r.HandleFunc(employee.Employees.ByID, handler.DeleteEmployeeByID).Methods("DELETE")

	deptRepo := repositoryDept.NewBoltRepository(db)
	deptService := department.NewService(deptRepo)
	deptHandler := department.NewHandler(deptService)

	r.HandleFunc(department.Departments.ByID, deptHandler.GetAllDepartments).Methods("GET")
	r.HandleFunc(department.Departments.ByID, deptHandler.GetDepartmentById).Methods("GET")
	r.HandleFunc(department.Departments.Base, deptHandler.CreateDepartment).Methods("POST")
	r.HandleFunc(department.Departments.ByID, deptHandler.UpdateDepartmentByID).Methods("PUT")
	r.HandleFunc(department.Departments.ByID, deptHandler.DeleteDepartmentByID).Methods("DELETE")

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
