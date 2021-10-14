package services

import (
	"fmt"

	"github.com/gosimple/slug"
	"mynamebvh.com/blog/internal/entities"
	"mynamebvh.com/blog/src/post/dto"
	"mynamebvh.com/blog/src/post/repositories"
)

type PostServiceInterface interface {
	FindByAll() []entities.Post
	FindById(id uint) entities.Post
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

func (c *PostService) FindById(id uint) entities.Post {
	return c.postRepository.FindByID(id)
}

func (c *PostService) FindByAll() []entities.Post {
	return c.postRepository.FindAll()
}

func (c *PostService) Save(post dto.Post) (entities.Post, error) {

	slug := slug.Make(post.Title)

	newPost := entities.Post{
		Title:     post.Title,
		Published: post.Published,
		Content:   post.Content,
		UserID:    post.UserID,
		Slug:      slug,
	}

	return c.postRepository.Save(newPost), nil
}

func (c *PostService) Delete(id uint) error {
	isId := c.postRepository.FindByID(id)

	if isId.ID == 0 {
		return fmt.Errorf("Id không tồn tại")
	}

	if err := c.postRepository.Delete(id); err != nil {
		return err
	}
	return nil
}

func (c *PostService) Update(id uint, post dto.PostUpdate) (entities.Post, error) {

	slug := slug.Make(post.Title)

	isId := c.postRepository.FindByID(id)

	if isId.ID == 0 {
		return entities.Post{}, fmt.Errorf("Id không tồn tại")
	}

	postUpdate := dto.PostUpdate{
		Title:     post.Title,
		Content:   post.Content,
		Published: post.Published,
		Slug:      slug,
	}

	if result, err := c.postRepository.Update(id, postUpdate); err != nil {
		return entities.Post{}, err
	} else {
		return result, nil
	}
}
