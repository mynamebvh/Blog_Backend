package repositories

import (
	db "mynamebvh.com/blog/infrastructures/db"
	entities "mynamebvh.com/blog/internal/entities"
	dto "mynamebvh.com/blog/src/user/dto"
)

type UserRepositoryInterface interface {
	FindAll() []entities.User
	FindByID(id uint) entities.User
	FindByEmail(email string) entities.User
	Save(user entities.User) entities.User
	Update(userId uint, user dto.UserUpdate) (entities.User, error)
	Delete(id uint) error
}

type userRepository struct {
	DB db.SqlServer
}

func NewUserRepostiory(DB db.SqlServer) UserRepositoryInterface {
	return &userRepository{
		DB: DB,
	}
}

func (u *userRepository) FindAll() []entities.User {
	var products []entities.User
	u.DB.DB().Find(&products)

	return products
}

func (u *userRepository) FindByID(id uint) entities.User {
	var user entities.User
	u.DB.DB().First(&user, id)

	return user
}

func (u *userRepository) FindByEmail(email string) entities.User {
	var user entities.User
	u.DB.DB().Where("email = ?", email).First(&user)

	return user
}

func (u *userRepository) Save(user entities.User) entities.User {
	u.DB.DB().Save(&user)

	return user
}

func (u *userRepository) Update(userId uint, userUpdate dto.UserUpdate) (entities.User, error) {
	var user entities.User

	dataUpdate := map[string]interface{}{
		"fullname": userUpdate.Fullname,
		"gender":   userUpdate.Gender,
	}

	err := u.DB.DB().Model(&user).Where("id = ?", userId).Updates(dataUpdate).Error

	if err != nil {
		return entities.User{}, err
	}
	return user, nil
}

func (u *userRepository) Delete(id uint) error {
	return u.DB.DB().Delete(entities.User{}, id).Error
}
