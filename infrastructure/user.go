package infrastructure

import (
	"go_todo/domain/model"
	"go_todo/domain/repository"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repository.User {
	return &UserRepository{db: db}
}

func (ud *UserRepository) Create(u *model.User, g *model.Grade) error {
	// NOTE: 明示的にCommit,Rollbackを呼び出さない場合は、
	// 無名関数の返り値がerrorではない場合 => 自動でコミット
	// エラーが発生した場合 => トランザクション全体がロールバックされる
	return ud.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(u).Error; err != nil {
			return err
		}

		g.UserID = u.ID

		if err := tx.Create(g).Error; err != nil {
			return err
		}
		return nil
	})
}

// func (ud *UserRepository) Create(u *model.User) error {
// 	return ud.db.Transaction(func(tx *gorm.DB) error {
// 		if err := tx.Create(u).Error; err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }

func (ud *UserRepository) Update(u *model.User) error {
	return ud.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Debug().Save(u).Error; err != nil {
			return err
		}

		if u.Grade != nil {
			if err := tx.Debug().Save(u.Grade).Error; err != nil {
				return err
			}
		}

		return nil
	})
}

func (ud *UserRepository) Delete(id int) error {
	if err := ud.db.Where("id = ?", id).Delete(&model.User{}).Error; err != nil {
		return err
	}
	return nil
}

func (ud *UserRepository) Find(id int) (*model.User, error) {
	var user *model.User
	err := ud.db.Debug().Preload("Grade").First(&user, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

// TODO: 完成させる
func (ud *UserRepository) FindBy(name * string, grade *int) (*model.User, error) {
	var user *model.User;
	q := ud.db.Debug().Preload("Grade").Find(&user);
	if name != nil {
		q.Where("name = ?", name)
	}

	if grade != nil {
		q.Where("grade.score >= ?", grade)
	}

	err := q.Error;

	if err != nil {
		return nil, err
	}
	return user, nil
}

func (ud *UserRepository) FindAll() ([]*model.User, error) {
	var users []*model.User
	// err := ud.db.Debug().Preload("Grade").Joins("JOIN grade on grade.user_id = user.id AND grade.score >= ?", 5).Find(&users).Error;
	err := ud.db.Preload("Grade").Find(&users).Error;

	if err != nil {
		return nil, err
	}
	return users, nil
}