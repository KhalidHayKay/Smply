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

	short := utils.Encode(time.Now().UnixMicro())
	now := time.Now()

	res, err := config.DB.Exec(
		`INSERT INTO urls (original, short, created) VALUES (?, ?, ?)`,
		url,
		short,
		now,
	)
	if err != nil {
		return model.Url{}, err
	}

	id, err := res.LastInsertId()
	if err != nil {
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
