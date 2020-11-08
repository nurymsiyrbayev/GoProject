package model

type News struct {
	Id        uint
	Title     string
	Author    string
	Text      string
	Category  string
}

type NewsBuilder interface {
	SetId        (uint)    newsBuilder
	SetTitle     (string)  newsBuilder
	SetAuthor    (string)  newsBuilder
	SetText      (string)  newsBuilder
	SetCategory  (string)  newsBuilder
  Build()      *News
}

type newsBuilder struct {
	news News
}

func GetNewsBuilder() *newsBuilder {
	return &newsBuilder{
		news: News{},
	}
}

func (n *newsBuilder) Id(id uint) *newsBuilder {
	n.news.Id = id
	return n
}

func (n *newsBuilder) SetTitle(title string) *newsBuilder {
	n.news.Title = title
	return n
}

func (n *newsBuilder) SetAuthor(author string) *newsBuilder {
	n.news.Author = author
	return n
}

func (n *newsBuilder) SetText(text string) *newsBuilder {
	n.news.Text = text
	return n
}

func (n *newsBuilder) SetCategory(category string) *newsBuilder {
	n.news.Category = category
	return n
}

func (n *newsBuilder) Build() News {
	return n.news
}
