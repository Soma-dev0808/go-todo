package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code string
	Price uint
}

func main() {
	// goはDocker上で動作しており、goのコンテナからmysqlのコンテナにアクセスするには、Docker内部ネットワークでアクセスする
	// なのでmysqlコンテナの3306(mysql:3306)を指定。外部からの場合は127.0.0.1:3307
	dsn := "root:root@tcp(mysql:3306)/sampledb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Product{})

	db.Create(&Product{Code: "D42", Price: 100})

	var product Product
	db.First(&product, "code = ?", "D42")
	fmt.Println(product)


	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "Good \n")
	// })

	// log.Fatal(http.ListenAndServe(":8080", nil))
}