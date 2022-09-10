package handlers

import (
	"fmt"
	"io"
	"net/http"
)

// HelloHandler /hello のハンドラ
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		io.WriteString(w, "Hello, world!\n")
	} else {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
	}
}

// PostArticleHandler /article のハンドラ
func PostArticleHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		io.WriteString(w, "Posting Article...\n")
	} else {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
	}
}

// ArticleListHandler /article/list のハンドラ
func ArticleListHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		io.WriteString(w, "Article List\n")
	} else {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
	}
}

// ArticleDetailHandler /article/1 のハンドラ
func ArticleDetailHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		articleID := 1
		resString := fmt.Sprintf("Article No.%d\n", articleID)
		io.WriteString(w, resString)
	} else {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
	}
}

// PostNiceHandler /article/nice のハンドラ
func PostNiceHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		io.WriteString(w, "Posting Nice...\n")
	} else {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
	}
}

// PostCommentHandler /comment のハンドラ
func PostCommentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		io.WriteString(w, "Posting Comment...\n")
	} else {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
	}
}
