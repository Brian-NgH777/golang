package repository

import (
	models "learn/model"
)

type BodycategoriesRepository interface {
	FindBodycategories() error
	InsertBodycategory(u models.Bodycategories) error
}
