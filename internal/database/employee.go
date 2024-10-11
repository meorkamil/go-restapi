package database

import (
	"go-restapi/internal/model"

	"gorm.io/gorm"
)

type EmployeeRepository struct {
	DB *gorm.DB
}

// Init Employee repository
func InitRepo(c *gorm.DB) *EmployeeRepository {
	return &EmployeeRepository{
		DB: c,
	}
}

// Find all
func (repo *EmployeeRepository) FindAll() ([]model.Employee, error) {
	var employees []model.Employee
	result := repo.DB.Find(&employees)
	return employees, result.Error
}

// Find by id employee
func (repo *EmployeeRepository) FindByID(id uint) (model.Employee, error) {
	var employee model.Employee
	result := repo.DB.First(&employee, id)
	return employee, result.Error
}

// Create employee
func (repo *EmployeeRepository) CreateEmployee(employee *model.Employee) error {
	return repo.DB.Create(employee).Error
}
