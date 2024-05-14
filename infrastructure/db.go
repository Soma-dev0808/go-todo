package infrastructure

import (
	"github.com/go-sql-driver/mysql"
	gormmysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func NewDB() (*gorm.DB, error){
	// goはDocker上で動作しており、goのコンテナからmysqlのコンテナにアクセスするには、Docker内部ネットワークでアクセスする
	// なのでmysqlコンテナの3306(mysql:3306)を指定。外部からの場合は127.0.0.1:3307
	// dsn := "root:root@tcp(mysql:3306)/sampledb?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := mysql.Config{
		DBName: "sampledb",
		User: "root",
		Passwd: "root",
		Addr: "mysql:3306",
		Net: "tcp",
		ParseTime: true,
		Collation: "utf8mb4_unicode_ci",
		AllowNativePasswords: true,
	}

	db, err := gorm.Open(gormmysql.Open(dsn.FormatDSN()), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			// use singular table name, table for `User` would be `user` with this option enabled
			// https://gorm.io/ja_JP/docs/gorm_config.html
			SingularTable: true,
		},
	})

	if err != nil {
		println(err.Error())
		return nil, err
	}
	return db, nil
}