package employee

import (
	model "template-golang/internal/domain/employee"
	repository "template-golang/internal/repository/employee"
)

type Service struct {
	repo repository.Repository
}

func NewService(repo repository.Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetAllEmployees() ([]model.Employee, error) {
	return s.repo.GetAllEmployees()
}

func (s *Service) CreateEmployee(e model.Employee) {
	err := s.repo.CreateEmployee(e)
	if err != nil {
		return
	}
}
func (s *Service) GetEmployeeByID(id string) (*model.Employee, error) {
	return s.repo.GetEmployeeByID(id)
}

func (s *Service) UpdateEmployeeByID(id string, update model.Employee) error {
	return s.repo.UpdateEmployeeByID(id, update)
}

func (s *Service) DeleteEmployeeByID(id string) error {
	return s.repo.DeleteEmployeeByID(id)
}

func (s *Service) GetAllEmployeesByDepartmentID(deptID string) ([]model.Employee, error) {
	return s.repo.GetAllEmployeesByDepartmentID(deptID)
}
