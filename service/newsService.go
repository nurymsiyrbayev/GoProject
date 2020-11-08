package service

import (
	"errors"
	"final/database"
	"final/model"
)

type NewsService struct { }

func (NewsService) AddNews(news *model.News) error {
	_, err := database.GetInstance().NamedExec("INSERT INTO news(title, author, text, category_id) " +
																	 "SELECT :title, :author, :text, id FROM categories WHERE name ~* :category", &news);
	if err != nil {
		return errors.New("Something went wrong")
	}

	return nil
}
