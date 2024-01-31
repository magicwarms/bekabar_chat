package utils

import "gorm.io/gorm"

type paginate struct {
	limit int
	page  int64
}

func NewPaginate(limit int, page int64) *paginate {
	return &paginate{limit: limit, page: page}
}

func (p *paginate) PaginatedResult(db *gorm.DB) *gorm.DB {
	offset := (p.page - 1) * int64(p.limit)

	return db.Offset(int(offset)).
		Limit(p.limit)
}
