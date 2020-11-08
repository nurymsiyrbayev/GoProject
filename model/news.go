package model

type News struct {
	NewsId uint
	Title string
	Author string
	Text string
}

type NewsBuilder interface {
	SetNewsId(uint) newsBuilder
	SetTitle(string) newsBuilder
	SetAuthor(string) newsBuilder
	SetText(string) newsBuilder
	Build() *News
}

type newsBuilder struct {
	news News
}

func GetNewsBuilder() *newsBuilder {
	return &newsBuilder{
		news: News{},
	}
}

func (n *newsBuilder) SetNewsId(id uint) *newsBuilder {
	n.news.NewsId = id
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

func (n *newsBuilder) Build() News {
	return n.news
}
