package main

import (
	"gorm.io/driver/mysql"
	// 接続処理用のライブラリ
	"gorm.io/gorm"
	// ORMライブラリ：リレーショナルDB内のデータをGoで扱えるようオブジェクトに変換してくれる
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint // uint : 価格は常に非負であるため、この型が適しています
}

func main() {
	dsn := "root:password@tcp(mysql_db:3306)/dev?charset=utf8mb4&parseTime=True&loc=Local"
	// 対象のDBとの接続設定
	// ユーザー名:パスワード＠プロトコル(アドレス)/データベース名？オプション1=値&オプション2=値
	// Dockerの場合は、アドレスが毎回変わるので、コンテナ名を指定して接続する
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// dsnで指定した内容でDBのインスタンスを初期化

	db.AutoMigrate(&Product{})
	// Product構造体の内容を反映したテーブルが作られる
	db.Create(&Product{
		Code:  "D2Micro",
		Price: 10,
	})
	// DBにデータを登録する

	// Read
	var product Product
	db.First(&product, 1)
	// テーブルの一行目を取得する

	router := gin.Default()
	router.Use(cors.Default())

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"code": product,
			// 取得したデータをJsonに入れる
		})
	})

	router.Run(":8080")
}
