package repositories

import (
	db "mynamebvh.com/blog/infrastructures/db"
	"mynamebvh.com/blog/internal/entities"
	"mynamebvh.com/blog/src/post/dto"
)

type PostRepositoryInterface interface {
	FindAll() []entities.Post
	FindByID(id uint) entities.Post
	Save(post entities.Post) entities.Post
	Update(id uint, postUpdate dto.PostUpdate) (entities.Post, error)
	Delete(id uint) error
}

type PostRepository struct {
	DB db.SqlServer
}

func NewUserRepostiory(DB db.SqlServer) PostRepositoryInterface {
	return &PostRepository{
		DB: DB,
	}
}

func (u *PostRepository) FindAll() []entities.Post {
	var post []entities.Post
	u.DB.DB().Find(&post)

	return post
}

func (u *PostRepository) FindByID(id uint) entities.Post {
	var post entities.Post
	u.DB.DB().First(&post, id)

	return post
}

func (u *PostRepository) Save(post entities.Post) entities.Post {

	u.DB.DB().Save(&post)

	return post
}

func (u *PostRepository) Update(id uint, tagUpdate dto.PostUpdate) (entities.Post, error) {
	var postEntities entities.Post

	dataUpdate := map[string]interface{}{
		"name":        tagUpdate.Content,
		"description": tagUpdate.Content,
		"slug":        tagUpdate.Content,
	}

	err := u.DB.DB().Model(&postEntities).Where("id = ?", id).Updates(dataUpdate).Error

	if err != nil {
		return entities.Post{}, err
	}
	return postEntities, nil
}

func (u *PostRepository) Delete(id uint) error {
	return u.DB.DB().Delete(entities.Post{}, id).Error
}
