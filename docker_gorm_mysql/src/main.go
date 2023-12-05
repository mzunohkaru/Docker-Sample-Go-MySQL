package main

import (
	"docker_gorm_mysql/models"
	"docker_gorm_mysql/repositories"
	"docker_gorm_mysql/utils"
	"fmt"
)

func main() {
	db, err := utils.NewDBConnection()
	if err != nil {
		panic("データベースを開けませんでした")
	}

	err = db.AutoMigrate(&models.Product{})
	if err != nil {
		panic("マイグレーションに失敗しました")
	}

	// 新規ユーザーを作成
	var product models.Product

	product.Create("ニャンコ")

	err = repositories.CreateProduct(db, &product)
	// DBへ登録
	if err != nil {
		panic("商品の作成に失敗しました")
	}

	fmt.Println("新しい商品のID: ", product.ID)
}
