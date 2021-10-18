package repositories

import (
	"math"

	db "mynamebvh.com/blog/infrastructures/db"
	"mynamebvh.com/blog/internal/entities"
	"mynamebvh.com/blog/src/category/dto"
)

type CategoryRepositoryInterface interface {
	FindAll() []entities.Category
	FindByID(id uint) entities.Category
	FindPagination(id uint, page int, pageSize int, offset int) dto.CategoryResponse
	Save(category entities.Category) entities.Category
	Update(categoryId uint, categoryUpdate dto.CategoryUpdate) (entities.Category, error)
	Delete(id uint) error
}

type categoryRepository struct {
	DB db.SqlServer
}

func NewUserRepostiory(DB db.SqlServer) CategoryRepositoryInterface {
	return &categoryRepository{
		DB: DB,
	}
}

func (u *categoryRepository) FindAll() []entities.Category {

	var category []entities.Category
	u.DB.DB().Find(&category)

	return category
}

func (u *categoryRepository) FindByID(id uint) entities.Category {
	var category entities.Category
	u.DB.DB().First(&category, id)

	return category
}

func (u *categoryRepository) FindPagination(id uint, page int, pageSize int, offset int) dto.CategoryResponse {
	var categoryRaws []dto.CategoryRaw
	var total int64

	u.DB.DB().Table("posts").Select(
		"title",
		"posts.id as post_id",
		"posts.slug as post_slug",
		"fullname",
		"categories.description",
		"posts.created_at as created_at",
	).
		Joins("JOIN users ON posts.user_id = users.id").
		Joins("JOIN categories ON posts.category_id = categories.id").
		Where("categories.id = ?", id).
		Limit(offset).
		Offset(pageSize).
		Order("created_at desc").
		Scan(&categoryRaws)

	u.DB.DB().Model(&entities.Post{}).Where("category_id = ?", id).Count(&total)

	totalRow := int(math.Ceil(float64(total) / float64(pageSize)))

	return dto.CategoryResponse{
		CategoryRaw: categoryRaws,
		Pagination: dto.Pagination{
			Page:    page,
			PerPage: pageSize,
			Total:   totalRow,
		},
	}
}

func (u *categoryRepository) Save(category entities.Category) entities.Category {

	u.DB.DB().Save(&category)

	return category
}

func (u *categoryRepository) Update(categoryId uint, categoryUpdate dto.CategoryUpdate) (entities.Category, error) {
	var categoryEntities entities.Category

	dataUpdate := map[string]interface{}{
		"name":        categoryUpdate.Name,
		"description": categoryUpdate.Description,
		"slug":        categoryUpdate.Slug,
	}

	err := u.DB.DB().Model(&categoryEntities).Where("id = ?", categoryId).Updates(dataUpdate).Error

	if err != nil {
		return entities.Category{}, err
	}
	return categoryEntities, nil
}

func (u *categoryRepository) Delete(id uint) error {
	return u.DB.DB().Delete(entities.Category{}, id).Error
}
