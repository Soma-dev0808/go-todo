package infrastructure_test

import (
	"go_todo/domain/model"
	"go_todo/infrastructure"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func TestCreate(t *testing.T) {
	t.Parallel()
	t.Run("タスクの登録が行えること", func(t *testing.T) {
		todo := &model.Todo{Task: "task", Status: model.Created}
		db, mock, err := newDbMock()
		if err != nil {
			t.Errorf("Failed to initialize mock DB: %v", err)
			return
		}
		// repositoryはsqlmock(db)を用いて作成される
		repository := infrastructure.NewTodo(db)
		// repository.Createを実行した際にDBへ接続されるが、sqlmockにアクセスして実行されるので、
		// 以下のmockで想定する挙動を書き、それと一致しているかを確かめる。
		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `todo` (`task`,`status`) VALUES (?,?)")).
			WithArgs(todo.Task, todo.Status).
			WillReturnResult(sqlmock.NewResult(0, 1))
		mock.ExpectCommit()
		err = repository.Create(todo)
		if err != nil {
			t.Errorf("want = %v, got = %v", nil, err)
		}
	})
}

func TestUpdate(t *testing.T) {
	t.Parallel()
	t.Run("タスクの更新が行えること", func(t *testing.T) {
		todo := &model.Todo{ID: 1, Task: "task", Status: model.Created}
		db, mock, err := newDbMock()
		if err != nil {
			t.Errorf("Failed to initialize mock DB: %v", err)
			return
		}
		repository := infrastructure.NewTodo(db)
		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta("UPDATE `todo` SET `task`=?,`status`=? WHERE `id` = ?")).
			WithArgs(todo.Task, todo.Status, todo.ID).
			WillReturnResult(sqlmock.NewResult(0, 1))
		mock.ExpectCommit()
		if err := repository.Update(todo); err != nil {
			t.Errorf("want = %v, got = %v", nil, err)
		}
	})
}

func TestDelete(t *testing.T) {
	t.Parallel()
	t.Run("タスクの削除が行えること", func(t *testing.T) {
		todo := &model.Todo{ID: 1, Task: "task", Status: model.Created}
		db, mock, err := newDbMock()
		if err != nil {
			t.Errorf("Failed to initialize mock DB: %v", err)
			return
		}
		repository := infrastructure.NewTodo(db)
		// 実際にデータベースにデータが存在するかどうかよりも、sqlmockを使って設定した期待（削除クエリの実行とその結果）が満たされるかどうかを確認
		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta("DELETE FROM `todo` WHERE id = ?")).
			WithArgs(todo.ID).
			WillReturnResult(sqlmock.NewResult(0, 1))
		mock.ExpectCommit()
		if err := repository.Delete(todo.ID); err != nil {
			t.Errorf("want = %v, got = %v", nil, err)
		}
	})
}

func TestFind(t *testing.T) {
	t.Parallel()
	t.Run("タスクの検索が行えること", func(t *testing.T) {
		todo := &model.Todo{ID: 1, Task: "task", Status: model.Created}
		db, mock, err := newDbMock()
		if err != nil {
			t.Errorf("Failed to initialize mock DB: %v", err)
			return
		}
		repository := infrastructure.NewTodo(db)
				// 実際にデータベースにデータが存在するかどうかよりも、sqlmockを使って設定した期待（削除クエリの実行とその結果）が満たされるかどうかを確認
		mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `todo` WHERE id = ? LIMIT ?")).
			WithArgs(todo.ID, 1).
			WillReturnRows(sqlmock.NewRows([]string{"id", "task", "status"}).AddRow(todo.ID, todo.Task, todo.Status))
		_, err = repository.Find(todo.ID)
		if err != nil {
			t.Errorf("want = %v, got = %v", nil, err)
		}
	})
}

func TestFindAll(t *testing.T) {
	t.Parallel()
	t.Run("タスクの検索が行えること", func(t *testing.T) {
		db, mock, err := newDbMock()
		if err != nil {
			t.Errorf("Failed to initialize mock DB: %v", err)
			return
		}
		repository := infrastructure.NewTodo(db)
		mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `todo`")).
			WillReturnRows(&sqlmock.Rows{})
		_, err = repository.FindAll()
		if err != nil {
			t.Errorf("want = %v, got = %v", nil, err)
		}
	})
}

func newDbMock() (*gorm.DB, sqlmock.Sqlmock, error) {
	// 実際にはデータベースへの接続は行われず、sqlmockによる模擬的な操作が行われる
	db, mock, err := sqlmock.New()
	if err != nil {
		return nil, mock, err
	}

	gormDB, err := gorm.Open(mysql.Dialector{
		Config: &mysql.Config{DriverName: "mysql", Conn: db, SkipInitializeWithVersion: true},
	},
		&gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
			},
		})
	if err != nil {
		return gormDB, mock, err
	}
	gormDB.Logger = gormDB.Logger.LogMode(logger.Info)

	return gormDB, mock, err
}