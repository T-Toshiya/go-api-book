package controllers

import (
	"encoding/json"
	"go-api-book/apperrors"
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
		err = apperrors.ReqBodyDecodeFailed.Wrap(err, "bad request body")
		apperrors.ErrorHandler(w, r, err)
		return
	}

	article, err := c.service.PostArticleService(reqArticle)
	if err != nil {
		apperrors.ErrorHandler(w, r, err)
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
			err = apperrors.BadParam.Wrap(err, "queryparam must be number")
			apperrors.ErrorHandler(w, r, err)
			return
		}
	} else {
		page = 1
	}

	articles, err := c.service.ListArticleService(page)
	if err != nil {
		apperrors.ErrorHandler(w, r, err)
		return
	}
	json.NewEncoder(w).Encode(articles)
}

// ArticleDetailHandler /article/1 のハンドラ
func (c *ArticleController) ArticleDetailHandler(w http.ResponseWriter, r *http.Request) {
	articleID, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		err = apperrors.BadParam.Wrap(err, "queryparam must be number")
		apperrors.ErrorHandler(w, r, err)
		return
	}

	article, err := c.service.GetArticleService(articleID)
	if err != nil {
		apperrors.ErrorHandler(w, r, err)
		return
	}

	json.NewEncoder(w).Encode(article)
}

// PostNiceHandler /article/nice のハンドラ
func (c *ArticleController) PostNiceHandler(w http.ResponseWriter, r *http.Request) {
	var reqArticle models.Article
	if err := json.NewDecoder(r.Body).Decode(&reqArticle); err != nil {
		err = apperrors.ReqBodyDecodeFailed.Wrap(err, "bad request body")
		apperrors.ErrorHandler(w, r, err)
		return
	}
	article, err := c.service.PostNiceService(reqArticle)
	if err != nil {
		apperrors.ErrorHandler(w, r, err)
		return
	}
	json.NewEncoder(w).Encode(article)
}
