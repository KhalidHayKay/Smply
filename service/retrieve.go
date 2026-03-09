package service

import (
	"shortener/config"
	"shortener/model"
)

func Retrieve(short string) (model.Url, error) {
	var url model.Url

	err := config.DB.QueryRow(
		`SELECT * FROM urls WHERE short = ?`,
		short).Scan(
		&url.Id,
		&url.Original,
		&url.Short,
		&url.Visited,
		&url.Created,
	)

	if err != nil {
		return model.Url{}, err
	}

	return url, nil
}

func GetByOriginal(originalUrl string) (model.Url, error) {
	var url model.Url

	err := config.DB.QueryRow(
		`SELECT * FROM urls WHERE original = ?`,
		originalUrl).Scan(
		&url.Id,
		&url.Original,
		&url.Short,
		&url.Visited,
		&url.Created,
	)

	if err != nil {
		return model.Url{}, err
	}

	return url, nil
}

func IncrementVisited(id int64) error {
	_, err := config.DB.Exec(`
		UPDATE urls
		SET visited = visited + 1
		WHERE id = ?
	`, id)

	return err
}
