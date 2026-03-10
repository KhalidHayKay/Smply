package service

import (
	"shortener/config"
	"shortener/model"
	"shortener/utils"
	"time"
)

func StoreUrl(url string) (model.Url, error) {
	saved, _ := GetByOriginal(url)

	if (saved != model.Url{}) {
		return saved, nil
	}

	now := time.Now()

	tx, err := config.DB.Begin()
	if err != nil {
		return model.Url{}, err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	res, err := tx.Exec(
		`INSERT INTO urls (original, created) VALUES (?, ?)`,
		url,
		now,
	)
	if err != nil {
		return model.Url{}, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return model.Url{}, err
	}

	short := utils.EncodeWithPadding(id)

	_, err = tx.Exec(`
		UPDATE urls
		SET short = ?
		WHERE id = ?
	`, short, id)
	if err != nil {
		return model.Url{}, err
	}

	if err = tx.Commit(); err != nil {
		return model.Url{}, err
	}

	result := model.Url{
		Id:       id,
		Original: url,
		Short:    short,
		Visited:  0,
		Created:  now,
	}

	return result, nil
}
