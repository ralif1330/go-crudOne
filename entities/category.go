package entities

import "time"

// sesuaikan data di struct dengan di database
// dan awali data di struct dengan huruf Kapital kecuali type data

type Category struct {
	Id        uint
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}