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
	article := models.Article1
	jsonData, err := json.Marshal(article)
	if err != nil {
		http.Error(w, "fail to encode json\n", http.StatusInternalServerError)
		return
	}
	w.Write(jsonData)
}

// ArticleListHandler /article/list のハンドラ
func ArticleListHandler(w http.ResponseWriter, r *http.Request) {
	articles := []models.Article{models.Article1, models.Article2}
	jsonData, err := json.Marshal(articles)
	if err != nil {
		http.Error(w, "fail to encode json\n", http.StatusInternalServerError)
		return
	}
	w.Write(jsonData)
}

// ArticleDetailHandler /article/1 のハンドラ
func ArticleDetailHandler(w http.ResponseWriter, r *http.Request) {
	article := models.Article1
	jsonData, err := json.Marshal(article)
	if err != nil {
		http.Error(w, "fail to encode json\n", http.StatusInternalServerError)
		return
	}
	w.Write(jsonData)
}

// PostNiceHandler /article/nice のハンドラ
func PostNiceHandler(w http.ResponseWriter, r *http.Request) {
	article := models.Article1
	jsonData, err := json.Marshal(article)
	if err != nil {
		http.Error(w, "fail to encode json\n", http.StatusInternalServerError)
		return
	}
	w.Write(jsonData)
}

// PostCommentHandler /comment のハンドラ
func PostCommentHandler(w http.ResponseWriter, r *http.Request) {
	comment := models.Comment1
	jsonData, err := json.Marshal(comment)
	if err != nil {
		http.Error(w, "fail to encode json\n", http.StatusInternalServerError)
		return
	}
	w.Write(jsonData)
}
