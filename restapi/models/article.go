package models

type Article struct {
	Id      string
	Title   string
	Desc    string
	Content string
}

var Articles = []Article{
	Article{
		Id:      "1",
		Title:   "Hello",
		Desc:    "Article Description",
		Content: "Article Content",
	},
	Article{
		Id:      "2",
		Title:   "Hello 2",
		Desc:    "Article Description",
		Content: "Article Content",
	},
}
