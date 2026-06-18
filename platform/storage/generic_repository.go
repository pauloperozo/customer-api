package storage

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type BaseRepository[T any] struct {
	db *gorm.DB
}

func NewBaseRepository[T any](db *gorm.DB) *BaseRepository[T] {
	return &BaseRepository[T]{db: db}
}

func (repository *BaseRepository[T]) GetAll() ([]T, error) {
	var results []T
	err := repository.db.Find(&results).Error
	return results, err
}

func (repository *BaseRepository[T]) GetByID(id any) (*T, error) {
	var result T
	err := repository.db.First(&result, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &result, nil
}

func (repository *BaseRepository[T]) Create(model *T) (*T, error) {
	err := repository.db.Create(model).Error
	if err != nil {
		return nil, err
	}
	return model, nil
}

func (repository *BaseRepository[T]) Update(id any, updateData map[string]any) (*T, error) {
	var result T

	err := repository.db.First(&result, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	err = repository.db.Model(&result).Updates(updateData).Error
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (repository *BaseRepository[T]) Delete(id any) error {
	var result T
	return repository.db.Delete(&result, id).Error
}

func (repository *BaseRepository[T]) GetByField(fieldName string, value any) (*T, error) {
	var result T

	query := fmt.Sprintf("%s = ?", fieldName)

	err := repository.db.Where(query, value).First(&result).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &result, nil
}
