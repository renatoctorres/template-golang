package employee

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"template-golang/internal/domain/employee"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

// @Summary Create Employee
// @Description post employee
// @Tags employee
// @Accept  json
// @Produce  json
// @Success 200 {array} Employee
// @Router /employees [post]
func (h *Handler) CreateEmployee(w http.ResponseWriter, r *http.Request) {
	var emp employee.Employee
	if err := json.NewDecoder(r.Body).Decode(&emp); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	h.service.CreateEmployee(emp)

	w.WriteHeader(http.StatusCreated)
	err := json.NewEncoder(w).Encode(emp)
	if err != nil {
		return
	}
}

// @Summary Get Employees
// @Description get list of employees
// @Tags employees
// @Accept  json
// @Produce  json
// @Success 200 {array} Employee
// @Router /employees [get]
func (h *Handler) GetAllEmployees(w http.ResponseWriter, r *http.Request) {
	employees, err := h.service.GetAllEmployees()
	if err != nil {
		http.Error(w, "Failed to retrieve employees", http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(employees)
	if err != nil {
		return
	}
}

// @Summary Get Employee by ID
// @Description get employee by ID
// @Tags employee
// @Accept  json
// @Produce  json
// @Param id path string true "ID"
// @Success 200 {array} Employee
// @Router /employees/{id} [get]
func (h *Handler) GetEmployeeById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		http.Error(w, "Employee ID is required", http.StatusBadRequest)
		return
	}

	employee, err := h.service.GetEmployeeByID(id)
	if err != nil {
		http.Error(w, "Employee not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(employee)
	if err != nil {
		return
	}
}

// @Summary Updete Employee
// @Description update employee
// @Tags employee
// @Accept  json
// @Produce  json
// @Param id path string true "ID"
// @Created 201
// @Router /employees/{id] [put]
func (h *Handler) UpdateEmployeeByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		http.Error(w, "Employee ID is required", http.StatusBadRequest)
		return
	}

	var emp employee.Employee
	if err := json.NewDecoder(r.Body).Decode(&emp); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Set the Employee ID from the URL parameter to ensure consistency
	emp.ID = id
	if err := h.service.UpdateEmployeeByID(id, emp); err != nil {
		http.Error(w, "Employee not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(emp)
	if err != nil {
		return
	}
}

// @Summary Delete Employee
// @Description delete employee
// @Tags employee
// @Accept  json
// @Param id path string true "ID"
// @NoContent 204
// @Router /employees/{id] [delete]
func (h *Handler) DeleteEmployeeByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		http.Error(w, "Employee ID is required", http.StatusBadRequest)
		return
	}

	err := h.service.DeleteEmployeeByID(id)
	if err != nil {
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// @Summary Get Employees by Department ID
// @Description get list of employees
// @Tags employees
// @Accept  json
// @Produce  json
// @Param dept_id path string true "Department ID"
// @Success 200 {array} Employee
// @Router /departments/{dept_id}/employees [get]
func (h *Handler) GetAllEmployeesByDepartmentID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	deptID, ok := vars["deptId"]
	if !ok {
		http.Error(w, "Department ID is required", http.StatusBadRequest)
		return
	}

	employees, err := h.service.GetAllEmployeesByDepartmentID(deptID)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	if len(employees) == 0 {
		w.WriteHeader(http.StatusNotFound)
		_, err := w.Write([]byte("No employees found in this department"))
		if err != nil {
			return
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(employees)
	if err != nil {
		return
	}
}
