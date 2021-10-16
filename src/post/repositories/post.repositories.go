package repositories

import (
	"sync"

	"github.com/gosimple/slug"
	db "mynamebvh.com/blog/infrastructures/db"
	"mynamebvh.com/blog/internal/entities"
	"mynamebvh.com/blog/src/post/dto"
)

type PostRepositoryInterface interface {
	FindAll() []entities.Post
	FindByID(id uint) dto.PostResponse
	Save(post entities.Post, tags []entities.Tag) entities.Post
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

func (u *PostRepository) FindByID(id uint) dto.PostResponse {
	var result []dto.PostRaw
	var listSlug []string

	u.DB.DB().Raw("SELECT title, content, published, user_id, tags.slug, users.fullname, posts.updated_at FROM post_tags INNER JOIN posts ON post_tags.post_id = posts.id INNER JOIN tags ON post_tags.tag_id = tags.id INNER JOIN users ON posts.user_id = users.id WHERE posts.id = ?", id).Scan(&result)

	for _, v := range result {
		listSlug = append(listSlug, v.Slug)
	}

	return dto.PostResponse{
		Title:     result[0].Title,
		Content:   result[0].Content,
		UserID:    result[0].UserID,
		Published: result[0].Published,
		Slug:      listSlug,
		Fullname:  result[0].Fullname,
		UpdateAt:  result[0].UpdateAt,
	}
}

func (u *PostRepository) Save(post entities.Post, tags []entities.Tag) entities.Post {
	u.DB.DB().Save(&post)
	u.DB.DB().Debug().Save(&tags)
	var wg sync.WaitGroup

	for _, v := range tags {
		wg.Add(1)
		go func(v entities.Tag) {

			if v.ID == 0 {
				u.DB.DB().Where("slug = ?", v.Slug).First(&v)
			}

			temp := entities.PostTag{
				PostID: post.ID,
				TagID:  v.ID,
			}
			u.DB.DB().Debug().Save(&temp)

			defer wg.Done()
		}(v)
	}
	wg.Wait()
	return post
}

func (u *PostRepository) Update(id uint, postUpdate dto.PostUpdate) (entities.Post, error) {
	var postEntities entities.Post

	slug := slug.Make(postUpdate.Title)

	dataUpdate := map[string]interface{}{
		"title":     postUpdate.Title,
		"content":   postUpdate.Content,
		"published": postUpdate.Published,
		"slug":      slug,
	}

	err := u.DB.DB().Model(&postEntities).Where("id = ?", id).Updates(&dataUpdate).Error

	if err != nil {
		return entities.Post{}, err
	}
	return postEntities, nil
}

func (u *PostRepository) Delete(id uint) error {
	return u.DB.DB().Delete(entities.Post{}, id).Error
}
