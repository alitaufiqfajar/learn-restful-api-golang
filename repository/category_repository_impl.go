package repository

import (
	"at_fajar/belajar-golang-restful-api/helper"
	"at_fajar/belajar-golang-restful-api/model/domain"
	"context"
	"database/sql"
	"errors"
)

type CategoryRepositoryimpl struct {
}

func (repository *CategoryRepositoryimpl) Save(ctx context.Context, tx sql.Tx, category domain.Category) domain.Category {
	SQL := "INSERT INTO categories(name) value (?)"
	result, err := tx.ExecContext(ctx, SQL, category.Name)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	category.Id = int(id)
	return category
}
func (repository *CategoryRepositoryimpl) Update(ctx context.Context, tx sql.Tx, category domain.Category) domain.Category {
	SQL := "UPDATE categories SET name = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, category.Name, category.Id)
	helper.PanicIfError(err)

	return category

}
func (repository *CategoryRepositoryimpl) Delete(ctx context.Context, tx sql.Tx, category domain.Category) {
	SQL := "DELETE FROM categories WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, category.Id)
	helper.PanicIfError(err)
}
func (repository *CategoryRepositoryimpl) FindById(ctx context.Context, tx sql.Tx, categoryId int) (domain.Category, error) {
	SQL := "SELECT id, name FROM categories WHERE id = ?"
	rows, err := tx.QueryContext(ctx, SQL, categoryId)
	helper.PanicIfError(err)

	category := domain.Category{}
	if rows.Next() {
		err := rows.Scan(&category.Id, &category.Name)
		helper.PanicIfError(err)

		return category, nil
	} else {
		return category, errors.New("category not found.")
	}
}
func (repository *CategoryRepositoryimpl) FindAll(ctx context.Context, tx sql.Tx, category domain.Category) []domain.Category {
	SQL := "SELECT id, name FROM categories"

	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)

	var categories []domain.Category

	for rows.Next() {
		category := domain.Category{}
		err := rows.Scan(&category.Id, &category.Name)
		helper.PanicIfError(err)

		categories = append(categories, category)
	}

	return categories
}
