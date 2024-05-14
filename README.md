# コマンド

## Docker 起動

`docker compose up`

## Mysql アクセス

`mysql -h 127.0.0.1 -P 3307 -u root -p`

## コンテナから Mysql にアクセス

`docker exec -it db-for-go bash
`
`mysql -u root -p`

USE {DB_NAME}

## Go コンテナでマイグレーションを実行

`docker exec -it go_todo-go-1 bash`

`migrate -path=/app/migrations -database="mysql://root:root@tcp(mysql:3306)/sampledb?charset=utf8mb4&parseTime=True&loc=Local" up`

## Test 方法

`docker exec -it go_todo-go-1 bash`

xxx_test.go のあるディレクトリで

`go test`を実行
