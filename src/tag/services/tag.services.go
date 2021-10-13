package services

import (
	"fmt"

	"github.com/gosimple/slug"
	"mynamebvh.com/blog/internal/entities"
	"mynamebvh.com/blog/src/tag/dto"
	"mynamebvh.com/blog/src/tag/repositories"
)

type TagServiceInterface interface {
	FindByAll() []entities.Tag
	FindById(id uint) entities.Tag
	Save(tag dto.Tag) (entities.Tag, error)
	Delete(id uint) error
	Update(id uint, tag dto.Tag) (entities.Tag, error)
}

type TagService struct {
	tagRepository repositories.TagRepositoryInterface
}

func NewUserService(
	tagRepository repositories.TagRepositoryInterface,
) TagServiceInterface {
	return &TagService{
		tagRepository: tagRepository,
	}
}

func (c *TagService) FindById(id uint) entities.Tag {
	return c.tagRepository.FindByID(id)
}

func (c *TagService) FindByAll() []entities.Tag {
	return c.tagRepository.FindAll()
}

func (c *TagService) Save(tag dto.Tag) (entities.Tag, error) {

	slug := slug.Make(tag.Name)

	newTag := entities.Tag{
		Name:        tag.Name,
		Description: tag.Description,
		Slug:        slug,
	}

	return c.tagRepository.Save(newTag), nil
}

func (c *TagService) Delete(id uint) error {
	isId := c.tagRepository.FindByID(id)

	if isId.ID == 0 {
		return fmt.Errorf("Id không tồn tại")
	}

	if err := c.tagRepository.Delete(id); err != nil {
		return err
	}
	return nil
}

func (c *TagService) Update(id uint, tag dto.Tag) (entities.Tag, error) {

	slug := slug.Make(tag.Name)

	isId := c.tagRepository.FindByID(id)

	if isId.ID == 0 {
		return entities.Tag{}, fmt.Errorf("Id không tồn tại")
	}

	tagUpdate := dto.TagUpdate{
		Name:        tag.Name,
		Description: tag.Description,
		Slug:        slug,
	}

	if result, err := c.tagRepository.Update(id, tagUpdate); err != nil {
		return entities.Tag{}, err
	} else {
		return result, nil
	}
}
