package services

import (
	"fmt"

	slugUtils "github.com/gosimple/slug"
	"mynamebvh.com/blog/internal/entities"
	"mynamebvh.com/blog/internal/utils"
	"mynamebvh.com/blog/src/post/dto"
	"mynamebvh.com/blog/src/post/repositories"
)

type PostServiceInterface interface {
	FindByAll() []entities.Post
	FindById(id uint) dto.PostResponse
	Save(post dto.Post) (entities.Post, error)
	Delete(id uint) error
	Update(id uint, post dto.PostUpdate) (entities.Post, error)
}

type PostService struct {
	postRepository repositories.PostRepositoryInterface
}

func NewUserService(
	postRepository repositories.PostRepositoryInterface,
) PostServiceInterface {
	return &PostService{
		postRepository: postRepository,
	}
}

func (c *PostService) FindById(id uint) dto.PostResponse {
	return c.postRepository.FindByID(id)
}

func (c *PostService) FindByAll() []entities.Post {
	return c.postRepository.FindAll()
}

func (c *PostService) Save(post dto.Post) (entities.Post, error) {

	arrId := []string{post.UserID, post.CategoryID}

	result, err := utils.ParseUint(arrId)

	if err != nil {
		return entities.Post{}, err
	}

	slug := slugUtils.Make(post.Title)

	newPost := entities.Post{
		Title:      post.Title,
		Published:  post.Published,
		Content:    post.Content,
		UserID:     result[0],
		CategoryID: result[1],
		Slug:       slug,
	}

	newSlugs := []entities.Tag{}
	listTag := post.Tags

	for _, t := range listTag {
		slug := slugUtils.Make(t.Name)
		newSlugs = append(newSlugs, entities.Tag{
			Name: t.Name,
			Slug: slug,
		})
	}

	return c.postRepository.Save(newPost, newSlugs), nil
}

func (c *PostService) Delete(id uint) error {
	isId := c.postRepository.FindByID(id)

	if isId.UserID == 0 {
		return fmt.Errorf("Id không tồn tại")
	}

	if err := c.postRepository.Delete(id); err != nil {
		return err
	}
	return nil
}

func (c *PostService) Update(id uint, post dto.PostUpdate) (entities.Post, error) {

	isId := c.postRepository.FindByID(id)

	if isId.UserID == 0 {
		return entities.Post{}, fmt.Errorf("Id không tồn tại")
	}

	postUpdate := dto.PostUpdate{
		Title:     post.Title,
		Content:   post.Content,
		Published: post.Published,
	}

	if result, err := c.postRepository.Update(id, postUpdate); err != nil {
		return entities.Post{}, err
	} else {
		return result, nil
	}
}
