package main

import (
	"go-api-book/handlers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlers.HelloHandler)
	http.HandleFunc("/article", handlers.PostArticleHandler)
	http.HandleFunc("/article/list", handlers.ArticleListHandler)
	http.HandleFunc("/article/1", handlers.ArticleDetailHandler)
	http.HandleFunc("/article/nice", handlers.PostNiceHandler)
	http.HandleFunc("/article/comment", handlers.PostCommentHandler)

	log.Println("server start at port 3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
