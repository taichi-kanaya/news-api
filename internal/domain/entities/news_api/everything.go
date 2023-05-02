package news_api

/*
 * NewsAPIの記事保存用モデル
 * ref: https://newsapi.org/docs/endpoints/everything
 */
type Everything struct {
	Status       string    `json:"status"`
	TotalResults int       `json:"totalResults"`
	Articles     []Article `json:"articles"`
}

type Article struct {
	Source      ArticleSource `json:"source"`
	Author      string        `json:"author"`
	Title       string        `json:"title"`
	Description string        `json:"description"`
	URL         string        `json:"url"`
	URLToImage  string        `json:"urlToImage"`
	PublishedAt string        `json:"publishedAt"`
	Content     string        `json:"content"`
}

type ArticleSource struct {
	Id   interface{} `json:"id"`
	Name string      `json:"name"`
}
