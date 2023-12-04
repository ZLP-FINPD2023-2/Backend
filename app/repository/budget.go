package repository

import (
	"errors"
	"gorm.io/gorm"

	"finapp/lib"
	"finapp/models"
)

type BudgetRepository struct {
	logger   lib.Logger
	Database lib.Database
}

func NewBudgetRepository(logger lib.Logger, db lib.Database) BudgetRepository {
	return BudgetRepository{
		logger:   logger,
		Database: db,
	}
}

func (r BudgetRepository) WithTrx(trxHandle *gorm.DB) BudgetRepository {
	if trxHandle == nil {
		r.logger.Error("Transaction Database not found in gin context. ")
		return r
	}
	r.Database.DB = trxHandle
	return r
}

func (r BudgetRepository) Get(id, userID uint) (models.Budget, error) {
	var budget models.Budget
	err := r.Database.Where("user_id = ?", userID).Where("id = ?", id).First(&budget).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("can't find budget")
	}
	if errors.Is(err, gorm.ErrNotImplemented) {
		err = errors.New("unknown error")
	}
	return budget, err
}
