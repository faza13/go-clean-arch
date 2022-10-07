package database

import (
	"base/app/common"
	"gorm.io/gorm"
	"math"
)

type Provider struct {
	DB *gorm.DB
}

func (p *Provider) ScopePaginate(page int, perPage int) func(db *gorm.DB) *gorm.DB {
	offset := (page - 1) * perPage
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(offset).Limit(perPage)
	}
}

func (p *Provider) GetPaginate(query *gorm.DB, page int, perPage int) common.Pagination {
	var totalRows int64
	query.Count(&totalRows)

	return common.Pagination{
		TotalRows:   int(totalRows),
		CurrentPage: page,
		TotalPage:   p.calcTotalPage(totalRows, perPage),
		PerPage:     perPage,
	}
}

func (p *Provider) calcTotalPage(totalRow int64, perPage int) int {
	return int(math.Ceil(float64(int(totalRow) / perPage)))
}
