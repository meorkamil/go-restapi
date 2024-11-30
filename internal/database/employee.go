package database

import (
	"errors"
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
func (repo *EmployeeRepository) FindByID(empid string, password string) (model.Employee, error) {
	var employee model.Employee
	result := repo.DB.Where("empid = ? AND password = ?", empid, password).First(&employee)
	return employee, result.Error
}

// Create employee
func (repo *EmployeeRepository) CreateEmployee(employee *model.Employee) error {
	return repo.DB.Create(employee).Error
}

// Create app user
func (repo *EmployeeRepository) CreateAppUser(employee *model.Employee) error {
	err := repo.DB.First(&employee, "empid = ?", employee.Empid).Error
	switch {
	case err != nil:
		return repo.DB.Create(employee).Error
	case errors.Is(err, gorm.ErrRecordNotFound):
		return repo.DB.Create(employee).Error
	default:
		return nil
	}
}
