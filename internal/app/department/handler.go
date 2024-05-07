package department

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"template-golang/internal/domain/department"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

// @Summary Create Department
// @Description post department
// @Tags department
// @Accept  json
// @Produce  json
// @Success 200 {array} Department
// @Router /departments [post]
func (h *Handler) CreateDepartment(w http.ResponseWriter, r *http.Request) {
	var emp department.Department
	if err := json.NewDecoder(r.Body).Decode(&emp); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	h.service.CreateDepartment(emp)

	w.WriteHeader(http.StatusCreated)
	err := json.NewEncoder(w).Encode(emp)
	if err != nil {
		return
	}
}

// @Summary Get Departments
// @Description get list of departments
// @Tags departments
// @Accept  json
// @Produce  json
// @Success 200 {array} Department
// @Router /departments [get]
func (h *Handler) GetAllDepartments(w http.ResponseWriter, r *http.Request) {
	departments, err := h.service.GetAllDepartments()
	if err != nil {
		http.Error(w, "Failed to retrieve departments", http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(departments)
	if err != nil {
		return
	}
}

// @Summary Get Department by ID
// @Description get department by ID
// @Tags department
// @Accept  json
// @Produce  json
// @Param id path string true "ID"
// @Success 200 {array} Department
// @Router /departments/{id} [get]
func (h *Handler) GetDepartmentById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		http.Error(w, "Department ID is required", http.StatusBadRequest)
		return
	}

	department, err := h.service.GetDepartmentByID(id)
	if err != nil {
		http.Error(w, "Department not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(department)
	if err != nil {
		return
	}
}

// @Summary Updete Department
// @Description update department
// @Tags department
// @Accept  json
// @Produce  json
// @Param id path string true "ID"
// @Created 201
// @Router /departments/{id] [put]
func (h *Handler) UpdateDepartmentByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		http.Error(w, "Department ID is required", http.StatusBadRequest)
		return
	}

	var emp department.Department
	if err := json.NewDecoder(r.Body).Decode(&emp); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Set the Department ID from the URL parameter to ensure consistency
	emp.ID = id
	if err := h.service.UpdateDepartmentByID(id, emp); err != nil {
		http.Error(w, "Department not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(emp)
	if err != nil {
		return
	}
}

// @Summary Delete Department
// @Description delete department
// @Tags department
// @Accept  json
// @Param id path string true "ID"
// @NoContent 204
// @Router /departments/{id] [delete]
func (h *Handler) DeleteDepartmentByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		http.Error(w, "Department ID is required", http.StatusBadRequest)
		return
	}

	err := h.service.DeleteDepartmentByID(id)
	if err != nil {
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
