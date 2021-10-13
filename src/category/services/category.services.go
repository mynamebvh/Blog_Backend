package services

import (
	"fmt"

	"github.com/gosimple/slug"
	"mynamebvh.com/blog/internal/entities"
	"mynamebvh.com/blog/src/category/dto"
	"mynamebvh.com/blog/src/category/repositories"
)

type CategoryServiceInterface interface {
	FindByAll() []entities.Category
	FindById(id uint) entities.Category
	Save(category dto.Category) (entities.Category, error)
	Delete(id uint) error
	Update(id uint, category dto.Category) (entities.Category, error)
}

type CategoryService struct {
	categoryRepository repositories.CategoryRepositoryInterface
}

func NewUserService(
	categoryRepository repositories.CategoryRepositoryInterface,
) CategoryServiceInterface {
	return &CategoryService{
		categoryRepository: categoryRepository,
	}
}

func (c *CategoryService) FindById(id uint) entities.Category {
	return c.categoryRepository.FindByID(id)
}

func (c *CategoryService) FindByAll() []entities.Category {
	return c.categoryRepository.FindAll()
}

func (c *CategoryService) Save(category dto.Category) (entities.Category, error) {

	slug := slug.Make(category.Name)

	newCategory := entities.Category{
		Name:        category.Name,
		Description: category.Description,
		Slug:        slug,
	}

	return c.categoryRepository.Save(newCategory), nil
}

func (c *CategoryService) Delete(id uint) error {
	isId := c.categoryRepository.FindByID(id)

	if isId.ID == 0 {
		return fmt.Errorf("Id không tồn tại")
	}

	if err := c.categoryRepository.Delete(id); err != nil {
		return err
	}
	return nil
}

func (c *CategoryService) Update(id uint, category dto.Category) (entities.Category, error) {

	slug := slug.Make(category.Name)

	isId := c.categoryRepository.FindByID(id)

	if isId.ID == 0 {
		return entities.Category{}, fmt.Errorf("Id không tồn tại")
	}

	categoryUpdate := dto.CategoryUpdate{
		Name:        category.Name,
		Description: category.Description,
		Slug:        slug,
	}

	if user, err := c.categoryRepository.Update(id, categoryUpdate); err != nil {
		return entities.Category{}, err
	} else {
		return user, nil
	}
}
