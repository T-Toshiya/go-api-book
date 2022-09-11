package repositories

import (
	"database/sql"
	"go-api-book/models"

	_ "github.com/go-sql-driver/mysql"
)

const (
	articleNumPerPage = 5
)

func InsertArticle(db *sql.DB, article models.Article) (models.Article, error) {
	const sqlStr = `insert into articles (title, contents, username, nice, created_at) values (?, ?, ?, 0, now());`
	var newArticle models.Article

	newArticle.Title, newArticle.Contents, newArticle.UserName = article.Title, article.Contents, article.UserName

	result, err := db.Exec(sqlStr, article.Title, article.Contents, article.UserName)
	if err != nil {
		return models.Article{}, err
	}

	id, _ := result.LastInsertId()

	newArticle.ID = int(id)

	return newArticle, nil
}

func SelectArticleList(db *sql.DB, page int) ([]models.Article, error) {
	const sqlStr = `
		select article_id, title, contents, username, nice
		from articles
        limit ? offset ?;
    `

	rows, err := db.Query(sqlStr, articleNumPerPage, (page-1)*articleNumPerPage)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var articleArray []models.Article
	for rows.Next() {
		var article models.Article
		rows.Scan(&article.ID, &article.Title, &article.Contents,
			&article.UserName, &article.NiceNum)

		articleArray = append(articleArray, article)
	}

	return articleArray, nil
}

func SelectArticleDetail(db *sql.DB, articleID int) (models.Article, error) {
	const sqlStr = ` select *
        from articles
        where article_id = ?;
    `
	row := db.QueryRow(sqlStr, articleID)
	if err := row.Err(); err != nil {
		return models.Article{}, err
	}

	var article models.Article
	var createdTime sql.NullTime

	err := row.Scan(&article.ID, &article.Title, &article.Contents,
		&article.UserName, &article.NiceNum, &createdTime)

	if err != nil {
		return models.Article{}, err
	}

	if createdTime.Valid {
		article.CreatedAt = createdTime.Time
	}

	return article, nil
}

func UpdateNiceNum(db *sql.DB, articleID int) error {
	const sqlGetNice = ` select nice
        from articles
        where article_id = ?;
    `
	row := db.QueryRow(sqlGetNice, articleID)
	if err := row.Err(); err != nil {
		return err
	}

	var niceNum int
	err := row.Scan(&niceNum)
	if err != nil {
		return err
	}

	const sqlUpdateNice = `update articles set nice = ? where article_id = ?`
	_, err = db.Exec(sqlUpdateNice, niceNum+1, articleID)
	if err != nil {
		return err
	}

	return nil
}
