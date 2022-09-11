package handlers

import (
	"encoding/json"
	"go-api-book/models"
	"io"
	"net/http"
)

// HelloHandler /hello のハンドラ
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello, world!\n")
}

// PostArticleHandler /article のハンドラ
func PostArticleHandler(w http.ResponseWriter, r *http.Request) {
	var reqArticle models.Article
	if err := json.NewDecoder(r.Body).Decode(reqArticle); err != nil {
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
		return
	}

	article := reqArticle

	json.NewEncoder(w).Encode(article)
}

// ArticleListHandler /article/list のハンドラ
func ArticleListHandler(w http.ResponseWriter, r *http.Request) {
	articles := []models.Article{models.Article1, models.Article2}
	json.NewEncoder(w).Encode(articles)
}

// ArticleDetailHandler /article/1 のハンドラ
func ArticleDetailHandler(w http.ResponseWriter, r *http.Request) {
	article := models.Article1
	json.NewEncoder(w).Encode(article)
}

// PostNiceHandler /article/nice のハンドラ
func PostNiceHandler(w http.ResponseWriter, r *http.Request) {
	article := models.Article1
	json.NewEncoder(w).Encode(article)
}

// PostCommentHandler /comment のハンドラ
func PostCommentHandler(w http.ResponseWriter, r *http.Request) {
	comment := models.Comment1
	json.NewEncoder(w).Encode(comment)
}
