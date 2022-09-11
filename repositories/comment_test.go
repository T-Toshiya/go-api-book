package repositories_test

import (
	"go-api-book/models"
	"go-api-book/repositories"
	"testing"
)

func TestInsertComment(t *testing.T) {
	comment := models.Comment{
		ArticleID: 1,
		Message:   "test comment",
	}

	expectedCommentNum := 3
	newComment, err := repositories.InsertComment(testDB, comment)
	if err != nil {
		t.Error(err)
	}
	if newComment.CommentID != expectedCommentNum {
		t.Errorf("new article id is expected %d but got %d\n", expectedCommentNum, newComment.CommentID)
	}

	t.Cleanup(func() {
		const sqlStr = `
			delete from comments
			where comment_id = ?;
		`
		testDB.Exec(sqlStr, newComment.CommentID)
	})
}

func TestSelectCommentList(t *testing.T) {
	expectedNum := 2
	const articleID = 1

	got, err := repositories.SelectCommentList(testDB, articleID)

	if err != nil {
		t.Error(err)
	}

	if num := len(got); num != expectedNum {
		t.Errorf("want %d but got %d comments\n", expectedNum, num)
	}
}
