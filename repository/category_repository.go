package repository

import (
	"adit/belajar-golang-restful-api-udemy/model/domain"
	"context"
	"database/sql"
)

type CategoryRepository interface {
	Save(ctx context.Context, tx sql.Tx, category domain.Category) domain.Category
	Update(ctx context.Context, tx sql.Tx, category domain.Category) domain.Category
	Delete(ctx context.Context, tx sql.Tx, category domain.Category)
	FindById(ctx context.Context, tx sql.Tx, categoryId int) domain.Category
	FindAll(ctx context.Context, tx sql.Tx) []domain.Category
	// Save() return-nya
}