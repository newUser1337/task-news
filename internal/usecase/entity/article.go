package entity

type (
	NewsletterNewsItemResponse struct {
		ArticleURL        string `json:"articleURL"`
		NewsArticleID     int64  `json:"newsArticleID"`
		PublishDate       string `json:"publishDate"`
		Taxonomies        string `json:"taxonomies"`
		TeaserText        string `json:"teaserText"`
		ThumbnailImageURL string `json:"thumbnailImageURL"`
		Title             string `json:"title"`
		OptaMatchId       string `json:"optaMatchId"`
		LastUpdateDate    string `json:"lastUpdateDate"`
		IsPublished       string `json:"isPublished"`
	}

	NewsArticleResponse struct {
		ArticleURL        string `json:"articleURL"`
		NewsArticleID     int64  `json:"newsArticleID"`
		PublishDate       string `json:"publishDate"`
		Taxonomies        string `json:"taxonomies"`
		TeaserText        string `json:"teaserText"`
		Subtitle          string `json:"subtitle"`
		ThumbnailImageURL string `json:"thumbnailImageURL"`
		Title             string `json:"title"`
		BodyText          string `json:"bodyText"`
		GalleryImageURLs  string `json:"galleryImageURLs"`
		VideoURL          string `json:"videoURL"`
		OptaMatchId       string `json:"optaMatchId"`
		LastUpdateDate    string `json:"lastUpdateDate"`
		IsPublished       string `json:"isPublished"`
	}
)
