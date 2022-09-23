package controllers

import (
	"encoding/json"
	"go-api-book/controllers/services"
	"go-api-book/models"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type ArticleController struct {
	service services.ArticleServicer
}

func NewArticleController(s services.ArticleServicer) *ArticleController {
	return &ArticleController{service: s}
}

// HelloHandler /hello のハンドラ
func (c *ArticleController) HelloHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello, world!\n")
}

// PostArticleHandler /article のハンドラ
func (c *ArticleController) PostArticleHandler(w http.ResponseWriter, r *http.Request) {
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
func (c *ArticleController) ArticleListHandler(w http.ResponseWriter, r *http.Request) {
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
func (c *ArticleController) ArticleDetailHandler(w http.ResponseWriter, r *http.Request) {
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
func (c *ArticleController) PostNiceHandler(w http.ResponseWriter, r *http.Request) {
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
