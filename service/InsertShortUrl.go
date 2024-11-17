package service

import (
	"github.com/link33d/shorter/config"
	"github.com/link33d/shorter/model"
)

func InsertShortUrl(link model.Link) (*model.Link, error) {

	var id int

	query, err := config.GetDB().Prepare("INSERT INTO link (url, code) VALUES ($1, $2) RETURNING id")
	if err != nil {
		return nil, err
	}

	err = query.QueryRow(link.Url, link.Code).Scan(&id)
	if err != nil {
		return nil, err
	}

	query.Close()

	link.Id = uint(id)

	return &link, nil
}
