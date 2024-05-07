package department

import (
	model "template-golang/internal/domain/department"
	repository "template-golang/internal/repository/department"
)

type Service struct {
	repo repository.Repository
}

func NewService(repo repository.Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetAllDepartments() ([]model.Department, error) {
	return s.repo.GetAllDepartments()
}

func (s *Service) CreateDepartment(e model.Department) {
	err := s.repo.CreateDepartment(e)
	if err != nil {
		return
	}
}
func (s *Service) GetDepartmentByID(id string) (*model.Department, error) {
	return s.repo.GetDepartmentByID(id)
}

func (s *Service) UpdateDepartmentByID(id string, update model.Department) error {
	return s.repo.UpdateDepartmentByID(id, update)
}

func (s *Service) DeleteDepartmentByID(id string) error {
	return s.repo.DeleteDepartmentByID(id)
}
