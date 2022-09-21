package controllers

import (
	"encoding/json"
	"go-api-book/models"
	"go-api-book/services"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type MyAppController struct {
	service *services.MyAppService
}

func NewMyAppController(s *services.MyAppService) *MyAppController {
	return &MyAppController{service: s}
}

// HelloHandler /hello のハンドラ
func (c *MyAppController) HelloHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello, world!\n")
}

// PostArticleHandler /article のハンドラ
func (c *MyAppController) PostArticleHandler(w http.ResponseWriter, r *http.Request) {
	var reqArticle models.Article
	if err := json.NewDecoder(r.Body).Decode(&reqArticle); err != nil {
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
		return
	}

	article, err := c.service.PostArticleService(reqArticle)
	if err != nil {
		http.Error(w, "fail internal exec\n", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(article)
}

// ArticleListHandler /article/list のハンドラ
func (c *MyAppController) ArticleListHandler(w http.ResponseWriter, r *http.Request) {
	queryMap := r.URL.Query()
	// クエリパラメータ page を取得
	var page int
	if p, ok := queryMap["page"]; ok && len(p) > 0 {
		var err error
		page, err = strconv.Atoi(p[0])
		if err != nil {
			http.Error(w, "Invalid query parameter", http.StatusBadRequest)
			return
		}
	} else {
		page = 1
	}

	articles, err := c.service.ListArticleService(page)
	if err != nil {
		http.Error(w, "fail internal exec\n", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(articles)
}

// ArticleDetailHandler /article/1 のハンドラ
func (c *MyAppController) ArticleDetailHandler(w http.ResponseWriter, r *http.Request) {
	articleID, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Invalid query parameter", http.StatusBadRequest)
		return
	}

	article, err := c.service.GetArticleService(articleID)
	if err != nil {
		http.Error(w, "fail internal exec\n", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(article)
}

// PostNiceHandler /article/nice のハンドラ
func (c *MyAppController) PostNiceHandler(w http.ResponseWriter, r *http.Request) {
	var reqArticle models.Article
	if err := json.NewDecoder(r.Body).Decode(&reqArticle); err != nil {
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
		return
	}
	article, err := c.service.PostNiceService(reqArticle)
	if err != nil {
		http.Error(w, "fail internal exec\n", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(article)
}

// PostCommentHandler /comment のハンドラ
func (c *MyAppController) PostCommentHandler(w http.ResponseWriter, r *http.Request) {
	var reqComment models.Comment
	if err := json.NewDecoder(r.Body).Decode(&reqComment); err != nil {
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
	}

	comment, err := c.service.PostCommentService(reqComment)
	if err != nil {
		http.Error(w, "fail internal exec\n", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(comment)
}
