package repository

import (
	"errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"reflect"
)

type Table interface {
	Table() string
}

func FindByID[T Table](tx *gorm.DB, id uint) (*T, error) {
	var (
		result = new(T)
		table  T
		err    error
	)

	if err = tx.Table(table.Table()).Where("id = ?", id).First(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}

func FindOne[T Table](tx *gorm.DB, where map[string]any) (*T, error) {
	var (
		result = new(T)
		table  T
		err    error
	)

	if err = tx.Table(table.Table()).Where(where).First(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}

func FindOneWithAssociation[T Table](tx *gorm.DB, where map[string]any) (*T, error) {
	var (
		result = new(T)
		table  T
		err    error
	)

	if err = tx.Table(table.Table()).Preload(clause.Associations).Where(where).First(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}

func FindMany[T Table](tx *gorm.DB, where map[string]any) ([]*T, error) {
	var (
		result []*T
		table  T
		err    error
	)

	if err = tx.Table(table.Table()).Where(where).Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}

func FindManyWithInlineCondition[T Table](tx *gorm.DB, where string, args ...any) ([]*T, error) {
	var (
		result []*T
		table  T
		err    error
	)

	if err = tx.Table(table.Table()).Where(where, args...).Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}

func FindManyWithLimitOffset[T Table](tx *gorm.DB, where string, limit, offset int, args ...any) ([]*T, error) {
	var (
		result []*T
		table  T
		err    error
	)

	if err = tx.Table(table.Table()).Where(where, args...).Limit(limit).Offset(offset).Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}

func FindOneWithOptions[T Table](tx *gorm.DB, opt *FindOption) (*T, error) {
	var result *T
	var table T

	var query *gorm.DB

	if opt == nil {
		return nil, errors.New("option cannot be nil")
	}

	if opt.Alias() != "" {
		query = tx.Table(opt.Alias())
	} else {
		query = tx.Table(table.Table())
	}

	if opt.Preload() != nil {
		for _, preload := range opt.Preload() {
			query = query.Preload(preload)
		}
	}

	if opt.Distinct() != nil {
		query = query.Distinct(opt.Distinct()...)
	}

	if opt.Field() != "" {
		query = query.Select(opt.Field())
	}

	if len(opt.Join()) > 0 {
		for _, join := range opt.Join() {
			query = query.Joins(join)
		}
	}

	if opt.Where() != "" {
		query = query.Where(opt.Where(), opt.WhereArgs()...)
	}

	if opt.Group() != "" {
		query = query.Group(opt.Group())
	}

	if opt.Count() != nil {
		query.Count(opt.Count())
	}

	if opt.Limit() > 0 {
		query = query.Limit(opt.Limit())
	}

	if opt.Offset() > 0 {
		query = query.Offset(opt.Offset())
	}

	if err := query.First(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}

func FindWithOptions[T Table](tx *gorm.DB, opt *FindOption) ([]*T, error) {
	var result []*T
	var table T

	var query *gorm.DB

	if opt == nil {
		return nil, errors.New("option cannot be nil")
	}

	if opt.Alias() != "" {
		query = tx.Table(opt.Alias())
	} else {
		query = tx.Table(table.Table())
	}

	if opt.Preload() != nil {
		for _, preload := range opt.Preload() {
			query = query.Preload(preload)
		}
	}

	if opt.Distinct() != nil {
		query = query.Distinct(opt.Distinct()...)
	}

	if opt.Field() != "" {
		query = query.Select(opt.Field())
	}

	if len(opt.Join()) > 0 {
		for _, join := range opt.Join() {
			query = query.Joins(" " + join + " ")
		}
	}

	if opt.Where() != "" {
		query = query.Where(opt.Where(), opt.WhereArgs()...)
	}

	if opt.Group() != "" {
		query = query.Group(opt.Group())
	}

	if opt.Limit() > 0 {
		query = query.Limit(opt.Limit())
	}

	if opt.Offset() > 0 {
		query = query.Offset(opt.Offset())
	}

	if opt.Count() != nil {
		if err := query.Find(&result).Count(opt.Count()).Error; err != nil {
			return nil, err
		}
	} else {
		if err := query.Find(&result).Error; err != nil {
			return nil, err
		}
	}

	return result, nil
}

func FindAll[T Table](tx *gorm.DB) ([]*T, error) {
	var (
		result []*T
		table  T
		err    error
	)

	if err = tx.Table(table.Table()).Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}

func Insert[T Table](tx *gorm.DB, model *T) error {
	var (
		table T
		err   error
	)

	if err = tx.Table(table.Table()).Create(&model).Error; err != nil {
		return err
	}

	return nil
}

func InsertMany[T Table](tx *gorm.DB, models []*T, batchSize int) error {
	var (
		table T
		err   error
	)

	if err = tx.Table(table.Table()).CreateInBatches(&models, batchSize).Error; err != nil {
		return err
	}

	return nil
}

func UpsertMany[T Table](tx *gorm.DB, models []*T, batchSize int) error {
	var (
		table T
		err   error
	)

	if err = tx.Table(table.Table()).Clauses(clause.OnConflict{UpdateAll: true}).CreateInBatches(&models, batchSize).Error; err != nil {
		return err
	}

	return nil
}

func Update[T Table](tx *gorm.DB, where, value map[string]any) error {
	var (
		table T
		err   error
	)

	err = tx.Table(table.Table()).Where(where).Updates(value).Error
	return err
}

func UpdateWithInlineCondition[T Table](tx *gorm.DB, where string, value map[string]any, args ...any) error {
	var table T

	return tx.Table(table.Table()).Where(where, args...).Updates(value).Error
}

func UpdateColumn[T Table](tx *gorm.DB, where map[string]any, column string, value any) error {
	var (
		table T
		err   error
	)

	err = tx.Table(table.Table()).Where(where).UpdateColumn(column, value).Error
	return err
}

func Delete[T Table](tx *gorm.DB, where map[string]any) error {
	var (
		table = new(T)
		err   error
	)

	err = tx.Where(where).Delete(&table).Error
	return err
}

func Count[T Table](tx *gorm.DB, where map[string]any) (int64, error) {
	var (
		table T
		count int64
		err   error
	)

	if err = tx.Table(table.Table()).Where(where).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

func FindWithJoin[T Table](tx *gorm.DB, join string, where map[string]any) ([]*T, error) {
	var (
		result []*T
		table  T
		err    error
	)

	if err = tx.Table(table.Table()).Joins(join).Where(where).Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}

func FindWithSubQuery[T Table](tx *gorm.DB, subQuery, where map[string]any) ([]*T, error) {
	var (
		result []*T
		table  T
		err    error
	)

	if err = tx.Table(table.Table()).Where(where).Where(subQuery).Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}

func Query(tx *gorm.DB, query string, dest any, args ...any) error {
	if reflect.TypeOf(dest).Kind() != reflect.Ptr {
		return errors.New("dest must be a pointer")
	}

	return tx.Raw(query, args...).Scan(dest).Error
}

func Exec(tx *gorm.DB, query string, args ...any) error {
	return tx.Exec(query, args...).Error
}
