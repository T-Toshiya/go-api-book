package handlers

import (
	"fmt"
	"io"
	"net/http"
)

// HelloHandler /hello のハンドラ
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello, world!\n")
}

// PostArticleHandler /article のハンドラ
func PostArticleHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Posting Article...\n")
}

// ArticleListHandler /article/list のハンドラ
func ArticleListHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Article List\n")
}

// ArticleDetailHandler /article/1 のハンドラ
func ArticleDetailHandler(w http.ResponseWriter, r *http.Request) {
	articleID := 1
	resString := fmt.Sprintf("Article No.%d\n", articleID)
	io.WriteString(w, resString)
}

// PostNiceHandler /article/nice のハンドラ
func PostNiceHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Posting Nice...\n")
}

// PostCommentHandler /comment のハンドラ
func PostCommentHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Posting Comment...\n")
}
