package calculationService

import (
	"back/internal/database"
	"gorm.io/gorm"
)

type CalculationRepository interface {
	CreateCalculation(calc Calculations) error
	GetAllCalculation() ([]Calculations, error)
	GetCalculationByID(id string) (Calculations, error)
	UpdateCalculation(calc Calculations) error
	DeleteCalculation(id string) error
}

type calcRepository struct {
	db *gorm.DB
}

func NewCalculationRepository(db *gorm.DB) CalculationRepository {
	return &calcRepository{db: db}
}

func (r *calcRepository) CreateCalculation(calc Calculations) error {
	return r.db.Create(&calc).Error
}

func (r *calcRepository) GetAllCalculation() ([]Calculations, error) {
	var calculations []Calculations
	err := r.db.Find(&calculations).Error
	return calculations, err
}

func (r *calcRepository) GetCalculationByID(id string) (Calculations, error) {
	var calculation Calculations
	err := database.DB.First(&calculation, "id = ?", id).Error
	return calculation, err
}

func (r *calcRepository) UpdateCalculation(calc Calculations) error {
	return r.db.Save(&calc).Error
}

func (r *calcRepository) DeleteCalculation(id string) error {
	return r.db.Delete(&Calculations{}, "id = ?", id).Error
}
