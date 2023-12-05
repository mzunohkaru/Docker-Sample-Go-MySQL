package repositories

import (
	"docker_gorm_mysql/models"
	"gorm.io/gorm"
)

/*
	データベースに整合性を保ったまま保存したい単位
	モデルと1対1の関係
	モデルをDBに保存や取得をする
*/

func CreateProduct(db *gorm.DB, product *models.Product) error {
	return db.Create(product).Error
}
