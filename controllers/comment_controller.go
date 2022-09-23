package controllers

import (
	"encoding/json"
	"go-api-book/apperrors"
	"go-api-book/controllers/services"
	"go-api-book/models"
	"net/http"
)

type CommentController struct {
	service services.CommentServicer
}

func NewCommentController(s services.CommentServicer) *CommentController {
	return &CommentController{service: s}
}

// PostCommentHandler /comment のハンドラ
func (c *CommentController) PostCommentHandler(w http.ResponseWriter, r *http.Request) {
	var reqComment models.Comment
	if err := json.NewDecoder(r.Body).Decode(&reqComment); err != nil {
		err = apperrors.ReqBodyDecodeFailed.Wrap(err, "bad request body")
		apperrors.ErrorHandler(w, r, err)
	}

	comment, err := c.service.PostCommentService(reqComment)
	if err != nil {
		apperrors.ErrorHandler(w, r, err)
		return
	}
	json.NewEncoder(w).Encode(comment)
}
