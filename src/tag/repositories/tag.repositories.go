package repositories

import (
	db "mynamebvh.com/blog/infrastructures/db"
	"mynamebvh.com/blog/internal/entities"
	"mynamebvh.com/blog/src/tag/dto"
)

type TagRepositoryInterface interface {
	FindAll() []entities.Tag
	FindByID(id uint) entities.Tag
	Save(tag entities.Tag) entities.Tag
	Update(id uint, tagUpdate dto.TagUpdate) (entities.Tag, error)
	Delete(id uint) error
}

type TagRepository struct {
	DB db.SqlServer
}

func NewUserRepostiory(DB db.SqlServer) TagRepositoryInterface {
	return &TagRepository{
		DB: DB,
	}
}

func (u *TagRepository) FindAll() []entities.Tag {
	var tag []entities.Tag
	u.DB.DB().Find(&tag)

	return tag
}

func (u *TagRepository) FindByID(id uint) entities.Tag {
	var tag entities.Tag
	u.DB.DB().First(&tag, id)

	return tag
}

func (u *TagRepository) Save(tag entities.Tag) entities.Tag {

	u.DB.DB().Save(&tag)

	return tag
}

func (u *TagRepository) Update(id uint, tagUpdate dto.TagUpdate) (entities.Tag, error) {
	var tagEntities entities.Tag

	dataUpdate := map[string]interface{}{
		"name":        tagUpdate.Name,
		"description": tagUpdate.Description,
		"slug":        tagUpdate.Slug,
	}

	err := u.DB.DB().Model(&tagEntities).Where("id = ?", id).Updates(dataUpdate).Error

	if err != nil {
		return entities.Tag{}, err
	}
	return tagEntities, nil
}

func (u *TagRepository) Delete(id uint) error {
	return u.DB.DB().Delete(entities.Tag{}, id).Error
}
