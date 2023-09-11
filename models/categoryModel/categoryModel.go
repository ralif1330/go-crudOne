package categorymodel

import (
	"crudone/config"
	"crudone/entities"
)

// sesuaikan dengan nama yang ada di database

// nantinya function GetAll akan dipanggil di controllers

func GetAll() []entities.Category {
	rows, err := config.DB.Query(`
	SELECT * FROM categories
	`)
		if err != nil {
			panic(err)
		}

		defer rows.Close()

		var categories []entities.Category

		for rows.Next() {
			var category entities.Category
			rows.Scan(
				&category.Id,
				&category.Name,
				&category.CreatedAt,
				&category.UpdatedAt,
			)
				if err != nil {
					panic(err)
				}

		// pindahkan category ke categories 
				categories = append(categories, category)
		}

		return categories


}