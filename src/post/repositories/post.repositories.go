package repositories

import (
	"math"
	"strings"
	"sync"

	"github.com/gosimple/slug"
	db "mynamebvh.com/blog/infrastructures/db"
	"mynamebvh.com/blog/internal/entities"
	"mynamebvh.com/blog/src/post/dto"
)

type PostRepositoryInterface interface {
	FindAll(page int, pageSize int, offset int) dto.PostPagination
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

func (u *PostRepository) FindAll(page int, pageSize int, offset int) dto.PostPagination {
	var total int64
	u.DB.DB().Model(&entities.Post{}).Count(&total)

	var postRaw []dto.PostEntitiesRaw
	var postPagination []dto.PostResponse
	postListMap := map[uint]dto.PostResponse{}

	subQuery := `(SELECT id, created_at, user_id, title, slug, published
								 FROM posts
                 ORDER BY created_at desc 
                 OFFSET ?
                 ROWS FETCH NEXT ? ROWS ONLY)`

	u.DB.DB().Debug().Table("post_tags").Select(
		"posts.id as post_id",
		"posts.title",
		"posts.slug",
		"users.fullname",
		"posts.slug",
		"user_id",
		"STRING_AGG(tags.slug,',') as tag_slug",
		"posts.created_at",
	).
		Joins("JOIN"+subQuery+"as posts ON post_tags.post_id = posts.id", offset, pageSize).
		Joins("JOIN tags ON post_tags.tag_id = tags.id").
		Joins("JOIN users ON posts.user_id = users.id").
		Group("posts.id,posts.title,posts.slug,users.fullname,posts.slug,user_id,posts.created_at").
		Order("created_at desc").
		Where("published = 'true'").
		Scan(&postRaw)

	for _, v := range postRaw {
		if postListMap[v.PostID].Fullname == "" {

			var arrSlug []string = strings.Split(v.TagSlug, ",")

			var postP = dto.PostResponse{
				PostID:   v.PostID,
				Slug:     v.Slug,
				Title:    v.Title,
				Content:  v.Content,
				Fullname: v.Fullname,
				UserID:   v.UserID,
				TagSlug:  arrSlug,
			}
			postPagination = append(postPagination, postP)
		}
	}

	totalRow := int(math.Ceil(float64(total) / float64(pageSize)))

	return dto.PostPagination{
		Posts: postPagination,
		Pagination: dto.Pagination{
			Page:    page,
			PerPage: pageSize,
			Total:   totalRow,
		},
	}
}

func (u *PostRepository) FindByID(id uint) dto.PostResponse {
	var result []dto.PostRaw
	var listSlug []string

	u.DB.DB().Raw(`
	SELECT title, content, published, user_id, tags.slug as tag_slug
	,posts.slug, users.fullname, posts.updated_at
	FROM post_tags 
	INNER JOIN posts ON post_tags.post_id = posts.id
	INNER JOIN tags ON post_tags.tag_id = tags.id
	INNER JOIN users ON posts.user_id = users.id
	WHERE posts.id = ?
`, id).Scan(&result)

	for _, v := range result {
		listSlug = append(listSlug, v.TagSlug)
	}

	return dto.PostResponse{
		Title:     result[0].Title,
		Content:   result[0].Content,
		UserID:    result[0].UserID,
		Published: result[0].Published,
		TagSlug:   listSlug,
		Fullname:  result[0].Fullname,
		UpdateAt:  result[0].UpdateAt,
		Slug:      result[0].Slug,
	}
}

func (u *PostRepository) Save(post entities.Post, tags []entities.Tag) entities.Post {
	u.DB.DB().Save(&post)
	u.DB.DB().Save(&tags)
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
