package models

import (
	"olx-women-workshop-2022-backend/database"
)

type Advert struct {
	Id          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Image       string  `json:"ad_image"`
}

func List() ([]Advert, error) {
	var adverts []Advert

	statement, err := database.GetConn().Query(`SELECT * FROM advert ORDER BY id DESC`)
	if err != nil {
		return adverts, err
	}

	for statement.Next() {
		advert := Advert{}

		err = statement.Scan(&advert.Id, &advert.Title, &advert.Description, &advert.Price, &advert.Image)
		if err != nil {
			return adverts, err
		}

		adverts = append(adverts, advert)
	}

	return adverts, nil
}

func Create(advert Advert) (Advert, error) {
	_, err := database.GetConn().Exec(
		`INSERT INTO advert (title,description,price,image_path) VALUE (?,?,?,?)`,
		advert.Title, advert.Description, advert.Price, advert.Image,
	)

	return advert, err
}

func Delete(advertID int) error {
	_, err := database.GetConn().Exec(`DELETE FROM advert WHERE id = ?`, advertID)

	return err
}
