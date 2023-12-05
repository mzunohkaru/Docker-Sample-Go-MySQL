package main

// * GORM

import (
	"gorm.io/driver/mysql"
	// 接続処理用のライブラリ
	"gorm.io/gorm"
	// ORMライブラリ : リレーショナルDB内のデータをGoで扱えるようオブジェクトに変換してくれる
	// 	GAORMのメソッドを使用することで、SQL文を書かなくて良くなる
	// GORMでできること : 基本的なDB操作に加え、DB操作の前後に呼び出せれる便利なメソッドがある。ログ出力機能。
	// 	便利なメソッドとして、データを登録する前に独自のルールで生成したIDを使ってデータを登録するものがある
)

type Product struct {
	gorm.Model
	Name  string
	Price uint // uint : 価格は常に非負であるため、この型が適しています
}

func sample_code() {
	dsn := "root:password@tcp(mysql_db:3306)/dev?charset=utf8mb4&parseTime=True&loc=Local"
	// 対象のDBとの接続設定
	// ユーザー名:パスワード＠プロトコル(アドレス)/データベース名？オプション1=値&オプション2=値
	// Dockerの場合は、アドレスが毎回変わるので、コンテナ名を指定して接続する
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// dsnで指定した内容でDBのインスタンスを初期化

	db.AutoMigrate(&Product{})
	// Product構造体の内容を反映したテーブルが作られる

	product_create := Product{Name: "マウス-Base", Price: 3700}
	db.Create(&product_create)
	// // DBにデータを登録する

	product_batches := []Product{
		{
			Name: "マウス-R", Price: 3700,
		},
		{
			Name: "マウス-SR", Price: 55000,
		},
		{
			Name: "人間", Price: 890000,
		},
	}
	db.CreateInBatches(product_batches, 2)
	// CreateInBatches(登録するデータ, まとめて登録する個数)
	// DBにデータをまとめて登録する

	var product_read Product
	db.First(&product_read, 1)
	// IDで並び替えて、一番初めのデータを取得

	db.Where("name = ?", "マウス-SR").First(&product_read)
	// Where("条件( 文字列 ) = ?", "第一引数の?に当たる値")

	product_update := Product{Name: "PC", Price: 100}
	db.Save(product_update)
	// 渡されたデータのIDがデータベースにない場合は、新規登録処理を行う

	db.Delete(&Product{}, 24)
	// この場合は、 user テーブルの ID が 5 のデータを削除
	// Delete(&モデル( どのテーブルに対するものか指定 ){}, IDを指定)

	var product_order []Product
	db.Order("price desc").Find(&product_order)
	// 並び替えの指定。この場合は、price を降順で並び替える
}