package repository

import (
	"fmt"

	"gorm.io/gorm"
)

func BuildQuery(db *gorm.DB, query map[string]interface{}) *gorm.DB {
	for key, value := range query {
		if key == "$or" {

			conds, ok := value.([]interface{})
			if !ok {
				continue
			}
			for _, cond := range conds {
				subCond, ok := cond.(map[string]interface{})
				if !ok {
					continue
				}

				db = db.Or(func(db *gorm.DB) *gorm.DB {
					return BuildQuery(db, subCond)
				})
			}
			continue
		}

		if key == "$and" {
			conds, ok := value.([]interface{})
			if !ok {
				continue
			}
			for _, cond := range conds {
				subCond, ok := cond.(map[string]interface{})
				if !ok {
					continue
				}
				db = db.Where(func(db *gorm.DB) *gorm.DB {
					return BuildQuery(db, subCond)
				})
			}
			continue
		}

		if opMap, ok := value.(map[string]interface{}); ok {
			for op, v := range opMap {
				switch op {
				case "$eq":
					db = db.Where(fmt.Sprintf("%s = ?", key), v)
				case "$ne":
					db = db.Where(fmt.Sprintf("%s <> ?", key), v)
				case "$gt":
					db = db.Where(fmt.Sprintf("%s > ?", key), v)
				case "$lt":
					db = db.Where(fmt.Sprintf("%s < ?", key), v)
				case "$gte":
					db = db.Where(fmt.Sprintf("%s >= ?", key), v)
				case "$lte":
					db = db.Where(fmt.Sprintf("%s <= ?", key), v)
				case "$in":
					db = db.Where(fmt.Sprintf("%s IN ?", key), v)
				case "$nin":
					db = db.Where(fmt.Sprintf("%s NOT IN ?", key), v)
				case "$like":
					db = db.Where(fmt.Sprintf("%s LIKE ?", key), v)
				case "$not":

					if subMap, ok := v.(map[string]interface{}); ok {
						for subOp, subVal := range subMap {
							switch subOp {
							case "$eq":
								db = db.Not(fmt.Sprintf("%s = ?", key), subVal)
							case "$ne":
								db = db.Not(fmt.Sprintf("%s <> ?", key), subVal)
							case "$gt":
								db = db.Not(fmt.Sprintf("%s > ?", key), subVal)
							case "$lt":
								db = db.Not(fmt.Sprintf("%s < ?", key), subVal)
							case "$gte":
								db = db.Not(fmt.Sprintf("%s >= ?", key), subVal)
							case "$lte":
								db = db.Not(fmt.Sprintf("%s <= ?", key), subVal)
							case "$in":
								db = db.Not(fmt.Sprintf("%s IN ?", key), subVal)
							case "$nin":
								db = db.Not(fmt.Sprintf("%s NOT IN ?", key), subVal)
							case "$like":
								db = db.Not(fmt.Sprintf("%s LIKE ?", key), subVal)
							}
						}
					} else {
						db = db.Not(fmt.Sprintf("%s = ?", key), v)
					}
				default:
				}
			}
		} else {

			db = db.Where(fmt.Sprintf("%s = ?", key), value)
		}
	}
	return db
}
