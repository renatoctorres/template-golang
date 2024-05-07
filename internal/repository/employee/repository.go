package employee

import (
	"encoding/json"
	"errors"
	"fmt"
	"go.etcd.io/bbolt"
	"template-golang/internal/domain/employee"
)

const (
	employeeBucket = "Employees"
)

type Repository interface {
	CreateEmployee(e employee.Employee) error
	GetAllEmployees() ([]employee.Employee, error)
	GetEmployeeByID(id string) (*employee.Employee, error)
	UpdateEmployeeByID(id string, update employee.Employee) error
	DeleteEmployeeByID(id string) error
	GetAllEmployeesByDepartmentID(deptID string) ([]employee.Employee, error)
}

type BoltRepository struct {
	db *bbolt.DB
}

func NewBoltRepository(db *bbolt.DB) *BoltRepository {
	return &BoltRepository{db: db}
}

func (r *BoltRepository) GetAllEmployees() ([]employee.Employee, error) {
	var employees []employee.Employee
	err := r.db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(employeeBucket))
		if b == nil {
			return nil // Nenhum bucket, sem dados
		}
		return b.ForEach(func(k, v []byte) error {
			var emp employee.Employee
			if err := json.Unmarshal(v, &emp); err != nil {
				return err
			}
			employees = append(employees, emp)
			return nil
		})
	})
	return employees, err
}

func (r *BoltRepository) CreateEmployee(e employee.Employee) error {
	return r.db.Update(func(tx *bbolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(employeeBucket))
		if err != nil {
			return err
		}
		encoded, err := json.Marshal(e)
		if err != nil {
			return err
		}
		return b.Put([]byte(e.ID), encoded)
	})
}

func (r *BoltRepository) GetEmployeeByID(id string) (*employee.Employee, error) {
	var employee *employee.Employee
	err := r.db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(employeeBucket))
		if b == nil {
			return errors.New("Employee bucket does not exist")
		}
		v := b.Get([]byte(id)) // Convertendo ID int para string para busca
		if v == nil {
			return errors.New("Employee not found")
		}
		return json.Unmarshal(v, &employee)
	})
	if err != nil {
		return nil, err
	}
	return employee, nil
} // Implementar os outros métodos CRUD semelhantemente

func (r *BoltRepository) UpdateEmployeeByID(id string, update employee.Employee) error {
	return r.db.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(employeeBucket))
		if b == nil {
			return errors.New("Employee bucket does not exist")
		}

		current := b.Get([]byte(id))
		if current == nil {
			return errors.New("Employee not found")
		}

		// Optionally, you might want to unmarshal the current employee data
		// and apply only specific changes or validate the changes
		var emp employee.Employee
		if err := json.Unmarshal(current, &emp); err != nil {
			return err
		}

		// Updating the employee with new data
		emp.Name = update.Name
		emp.Position = update.Position
		emp.DepartmentId = update.DepartmentId

		// Marshal the updated employee back to JSON
		updated, err := json.Marshal(emp)
		if err != nil {
			return err
		}

		// Save the updated employee back to the database
		return b.Put([]byte(id), updated)
	})
}

func (r *BoltRepository) DeleteEmployeeByID(id string) error {
	return r.db.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(employeeBucket))
		if b == nil {
			return errors.New("Employee bucket does not exist")
		}

		if exists := b.Get([]byte(id)); exists == nil {
			return errors.New("Employee not found")
		}

		// Delete the employee
		if err := b.Delete([]byte(id)); err != nil {
			return err
		}

		return nil
	})
}

func (r *BoltRepository) GetAllEmployeesByDepartmentID(deptID string) ([]employee.Employee, error) {
	var employees []employee.Employee

	err := r.db.View(func(tx *bbolt.Tx) error {
		// Acessa o bucket específico do departamento
		b := tx.Bucket(getDepartmentBucketName(deptID))
		if b == nil {
			return nil // Nenhum funcionário neste departamento
		}

		// Itera apenas sobre os empregados neste bucket de departamento
		return b.ForEach(func(k, v []byte) error {
			var emp employee.Employee
			if err := json.Unmarshal(v, &emp); err != nil {
				return err
			}
			employees = append(employees, emp)
			return nil
		})
	})

	if err != nil {
		return nil, err
	}

	return employees, nil
}

func getDepartmentBucketName(deptID string) []byte {
	return []byte(fmt.Sprintf("Department_%d", deptID))
}
