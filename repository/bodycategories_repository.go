package repository

import (
	models "learn/model"
)

type BodycategoriesRepository interface {
	FindBodycategories() ([]models.Bodycategories, error)
	InsertBodycategory(u models.Bodycategories) error
}
