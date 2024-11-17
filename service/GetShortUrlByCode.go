package service

import (
	"database/sql"

	"github.com/link33d/shorter/config"
	"github.com/link33d/shorter/model"
)

func GetShortUrlByCode(code string) (*model.Link, error) {

	query, err := config.GetDB().Prepare("SELECT * FROM link WHERE code=$1")
	if err != nil {
		return nil, err
	}

	var link model.Link

	err = query.QueryRow(code).Scan(
		&link.Id,
		&link.Url,
		&link.Code,
	)

	if err != nil {

		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	query.Close()

	return &link, nil

}
