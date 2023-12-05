package models

import (
	"github.com/google/uuid"
)

/*
	モデルを定義しているファイルは、ドメインとも呼ばれ、
	そのアプリで解決したい、ビジネスの領域に関するロジックを詰め込むところになる。
	そのため、他の技術には依存したくない。
*/

type Product struct {
	ID   uuid.UUID
	// ID を一意にするため
	Name string
}

// ファクトリメソッド : 何かを生成する際のロジックを閉じ込めて、呼び出したら出来上がるようなメソッド
// 集約を生成するときは、毎回呼び出すようなメソッド
func (product *Product) Create(name string) {
	product.ID = uuid.New()
	// DBに保存する前に、IDにUUIDをセットする
	product.Name = name
}
