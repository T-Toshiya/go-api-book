package services

import (
	"database/sql"
	"errors"
	"go-api-book/apperrors"
	"go-api-book/models"
	"go-api-book/repositories"
)

func (s *MyAppService) GetArticleService(articleID int) (models.Article, error) {
	article, err := repositories.SelectArticleDetail(s.db, articleID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = apperrors.NAData.Wrap(err, "no data")
			return models.Article{}, err
		}
		err = apperrors.GetDataFailed.Wrap(err, "fail to get data")
		return models.Article{}, err
	}

	commentList, err := repositories.SelectCommentList(s.db, articleID)
	if err != nil {
		err = apperrors.GetDataFailed.Wrap(err, "fail to get data")
		return models.Article{}, err
	}

	article.CommentList = append(article.CommentList, commentList...)

	return article, nil
}

func (s *MyAppService) PostArticleService(article models.Article) (models.Article, error) {
	newArticle, err := repositories.InsertArticle(s.db, article)
	if err != nil {
		err = apperrors.InsertDataFailed.Wrap(err, "fail to record data")
		return models.Article{}, err
	}
	return newArticle, nil
}

func (s *MyAppService) ListArticleService(page int) ([]models.Article, error) {
	articles, err := repositories.SelectArticleList(s.db, page)
	if err != nil {
		err = apperrors.GetDataFailed.Wrap(err, "failt to get data")
		return nil, err
	}

	if len(articles) == 0 {
		err := apperrors.NAData.Wrap(ErrNoData, "no data")
		return nil, err
	}

	return articles, nil
}

func (s *MyAppService) PostNiceService(article models.Article) (models.Article, error) {
	if err := repositories.UpdateNiceNum(s.db, article.ID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = apperrors.NoTargetData.Wrap(err, "no target data")
			return models.Article{}, err
		}
		err = apperrors.UpdateDataFailed.Wrap(err, "fail to update data")
		return models.Article{}, err
	}

	return models.Article{
		ID:          article.ID,
		Title:       article.Title,
		Contents:    article.Contents,
		UserName:    article.UserName,
		NiceNum:     article.NiceNum + 1,
		CommentList: article.CommentList,
		CreatedAt:   article.CreatedAt,
	}, nil
}
