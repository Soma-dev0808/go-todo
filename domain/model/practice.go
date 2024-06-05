package model

import (
	"time"

	"gorm.io/gorm"
)

// リレーション勉強
type Sample struct {
	// gorm.Model : ID, CreatedAt, UpdatedAt, DeletedAtをフィールドに持つ構造体
	gorm.Model
	Name string
	Birthday *time.Time
}

/*
*
* Belongs To
* Belongs Toは自分のテーブルが対象テーブルのレコードに所属する関係です。
* 下記は、ChildがParentに対してBelongs To(所属している)例です。
* Childにとって、Parentはなくてはならない存在です。
* ParentIDを外部キーにすることでモデル間の紐付けをしています。
*/
type Parent struct {
	gorm.Model
	Name string
  }
  
// `Child` belongs to `Parent`, `ParentID` is the foreign key
type Child struct {
	gorm.Model
	ParentID int
	Parent   Parent
	Name   string
}

/*
*
* Has One
* Has Oneは自分のテーブルが対象テーブルを1つ持っている関係です。
* 下記は、PersonがCreditCardを1つ持っている例です。
* PersonにとってCreditCardはなくてはならないという存在ではありません。（CreditCardを持っていないPersonもありえる場合）
* CreditCardのIDが外部キーになっています。
*
*/
type Person struct {
	gorm.Model
	CreditCard   CreditCard
}

// Person has one CreditCard, CreditCardID is the foreign key
type CreditCard struct {
	gorm.Model
	Number   string
	PersonID   uint
}
  
/*
*
* Has Many
* Has Manyは自分のテーブルが対象テーブルを0以上（複数）持っている関係です。
* 下記は、Person2がCreditCard2を複数持っている例です。スライスになっています。
*
*/
// Person2 has many CreditCards, Person2ID is the foreign key
type Person2 struct {
	gorm.Model
	CreditCard2 []CreditCard2
  }
  
type CreditCard2 struct {
	gorm.Model
	Number   string
	Person2ID  uint
}

  /*
*
* Many To Many
* Many To Manyは自分のテーブルと対象のテーブルが複対複の関係です。
* 下記は、あるGeniusが複数のLanguageを話せ、複数のGeniusがあるLanguageを話す例です。
*
*/
// Genius has and belongs to many languages, use `genius_languages` as join table
type Genius struct {
	gorm.Model
	Languages []Language `gorm:"many2many:genius_languages;"`
  }
  
type Language struct {
	gorm.Model
	Name string
}

// Modelからマイグレーションする場合は自動的にテーブル名が複数形になる
// Model: User => Table: users