package testdata

import "go-api-book/models"

var ArticleTestData = []models.Article{
	{
		ID:       1,
		Title:    "firstPost",
		Contents: "This is my first blog",
		UserName: "TOC",
		NiceNum:  2,
	},
	{
		ID:       2,
		Title:    "2nd",
		Contents: "Second blog post",
		UserName: "TOC",
		NiceNum:  4,
	},
}
