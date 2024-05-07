package department

import (
	"encoding/json"
	"errors"
	"go.etcd.io/bbolt"
	"template-golang/internal/domain/department"
)

const (
	departmentBucket = "Departments"
)

type Repository interface {
	CreateDepartment(e department.Department) error
	GetAllDepartments() ([]department.Department, error)
	GetDepartmentByID(id string) (*department.Department, error)
	UpdateDepartmentByID(id string, update department.Department) error
	DeleteDepartmentByID(id string) error
}

type BoltRepository struct {
	db *bbolt.DB
}

func NewBoltRepository(db *bbolt.DB) *BoltRepository {
	return &BoltRepository{db: db}
}

func (r *BoltRepository) GetAllDepartments() ([]department.Department, error) {
	var departments []department.Department
	err := r.db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(departmentBucket))
		if b == nil {
			return nil // Nenhum bucket, sem dados
		}
		return b.ForEach(func(k, v []byte) error {
			var emp department.Department
			if err := json.Unmarshal(v, &emp); err != nil {
				return err
			}
			departments = append(departments, emp)
			return nil
		})
	})
	return departments, err
}

func (r *BoltRepository) CreateDepartment(e department.Department) error {
	return r.db.Update(func(tx *bbolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(departmentBucket))
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

func (r *BoltRepository) GetDepartmentByID(id string) (*department.Department, error) {
	var request *department.Department
	err := r.db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(departmentBucket))
		if b == nil {
			return errors.New("Department bucket does not exist")
		}
		v := b.Get([]byte(id)) // Convertendo ID int para string para busca
		if v == nil {
			return errors.New("Department not found")
		}
		return json.Unmarshal(v, &request)
	})
	if err != nil {
		return nil, err
	}
	return request, nil
} // Implementar os outros métodos CRUD semelhantemente

func (r *BoltRepository) UpdateDepartmentByID(id string, update department.Department) error {
	return r.db.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(departmentBucket))
		if b == nil {
			return errors.New("Department bucket does not exist")
		}

		current := b.Get([]byte(id))
		if current == nil {
			return errors.New("Department not found")
		}

		// Optionally, you might want to unmarshal the current department data
		// and apply only specific changes or validate the changes
		var emp department.Department
		if err := json.Unmarshal(current, &emp); err != nil {
			return err
		}

		// Updating the department with new data
		emp.Name = update.Name

		// Marshal the updated department back to JSON
		updated, err := json.Marshal(emp)
		if err != nil {
			return err
		}

		// Save the updated department back to the database
		return b.Put([]byte(id), updated)
	})
}

func (r *BoltRepository) DeleteDepartmentByID(id string) error {
	return r.db.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(departmentBucket))
		if b == nil {
			return errors.New("Department bucket does not exist")
		}

		if exists := b.Get([]byte(id)); exists == nil {
			return errors.New("Department not found")
		}

		// Delete the department
		if err := b.Delete([]byte(id)); err != nil {
			return err
		}

		return nil
	})
}
