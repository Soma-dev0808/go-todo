package usecase

import (
	"fmt"
	"go_todo/domain/model"
	"go_todo/domain/repository"
)

type User interface {
	Create(name string, gradeScore int) error
	Update(id int, name string, gradeScore int) error
	Delete(id int) error
	Find(id int) (*model.User, error)
	FindAll() ([]*model.User, error)
}

type UserUseCase struct {
	userRepository repository.User
}

func NewUserUseCase(u repository.User) User {
	return &UserUseCase{userRepository: u}
}

func (u *UserUseCase) Create(name string, gradeScore int) error {
	if err := u.userRepository.Create(&model.User{Name: name}, &model.Grade{ Score: gradeScore }); err != nil {
		return err
	}
	return nil
}

// func (u *UserUseCase) Create(name string, gradeScore int) error {
// 	newUser := model.NewUser(name, gradeScore)
// 	if err := u.userRepository.Create(newUser); err != nil {
// 		return err
// 	}
// 	return nil
// }

func (u *UserUseCase) Update(id int, name string, gradeScore int) error {
	user, err := u.userRepository.Find(id)
	fmt.Println(user, err)
	if err != nil {
		return err
	}

	// TODO: name, gradeのどちらかがない場合でも問題なく更新できるようにする
	user.Name = name

	if user.Grade == nil {
		user.Grade = model.NewGrade(user.ID, gradeScore)
	} else {
		user.Grade.Score = gradeScore
	}

	if err := u.userRepository.Update(user); err != nil {
		return err
	}

	return nil
}

func (u *UserUseCase) Delete(id int) error {
	if err := u.userRepository.Delete(id); err != nil {
		return err
	} 
	return nil
}

func (u *UserUseCase) Find(id int) (*model.User, error) {
	user, err := u.userRepository.Find(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserUseCase) FindAll() ([]*model.User, error) {
	users, err := u.userRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return users, nil
}