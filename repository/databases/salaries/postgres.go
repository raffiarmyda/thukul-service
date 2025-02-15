package salaries

import (
	"aprian1337/thukul-service/business/salaries"
	"context"
	"gorm.io/gorm"
)

type PostgresSalariesRepository struct {
	ConnPostgres *gorm.DB
}

func NewPostgresSalariesRepository(conn *gorm.DB) *PostgresSalariesRepository {
	return &PostgresSalariesRepository{
		ConnPostgres: conn,
	}
}

func (repo *PostgresSalariesRepository) GetList(_ context.Context, search string) ([]salaries.Domain, error) {
	var data []Salaries
	err := repo.ConnPostgres.Find(&data)
	if err.Error != nil {
		return []salaries.Domain{}, err.Error
	}
	return ToListDomain(data), nil
}

func (repo *PostgresSalariesRepository) GetById(_ context.Context, id uint) (salaries.Domain, error) {
	var data Salaries
	err := repo.ConnPostgres.First(&data, "id=?", id)
	if err.Error != nil {
		return salaries.Domain{}, err.Error
	}
	return data.ToDomain(), nil
}

func (repo *PostgresSalariesRepository) Create(_ context.Context, domain salaries.Domain) (salaries.Domain, error) {
	salary := FromDomain(domain)
	err := repo.ConnPostgres.Create(&salary)
	if err.Error != nil {
		return salaries.Domain{}, err.Error
	}
	return salary.ToDomain(), nil
}
func (repo *PostgresSalariesRepository) Update(_ context.Context, domain salaries.Domain) (salaries.Domain, error) {
	salary := FromDomain(domain)
	err := repo.ConnPostgres.First(&salary)
	if err.Error != nil {
		return salaries.Domain{}, err.Error
	}
	salary.Minimal = domain.Minimal
	salary.Maximal = domain.Maximal
	repo.ConnPostgres.Save(&salary)
	return salary.ToDomain(), nil
}

func (repo *PostgresSalariesRepository) Delete(_ context.Context, id uint) (int64, error) {
	salary := Salaries{}
	err := repo.ConnPostgres.Delete(&salary, id)
	if err.Error != nil {
		return 0, err.Error
	}
	return err.RowsAffected, nil
}
