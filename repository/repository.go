package repository

import (
	"context"

	"cloud.google.com/go/datastore"
)

// Article represents an article.
type Article struct {
	Author string
	Title  string
}

// Repository is a repository which interacts with datastore.
type Repository interface {
	CreateArticle(ctx context.Context, article *Article) error
	GetArticle(ctx context.Context, identifier string) (Article, error)
}

type repository struct {
	client *datastore.Client
}

// NewRepository creates a new instance of Repository.
func NewRepository(client *datastore.Client) Repository {
	return &repository{
		client: client,
	}
}

// CreateArticle creates a new article in datastore.
func (r repository) CreateArticle(ctx context.Context, article *Article) error {
	key := datastore.NameKey("Article", article.Title, nil)
	_, err := r.client.Put(ctx, key, article)
	return err
}

// GetArticle gets an article from datastore.
func (r repository) GetArticle(ctx context.Context, identifier string) (Article, error) {
	var article Article
	key := datastore.NameKey("Article", identifier, nil)
	err := r.client.Get(ctx, key, &article)
	if err != nil {
		return Article{}, err
	}

	return article, nil
}
