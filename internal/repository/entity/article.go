package entity

type (
	NewsletterNewsItemDB struct {
		ArticleURL        string `bson:"articleUrl"`
		NewsArticleID     int64  `bson:"_id"`
		PublishDate       string `bson:"publishDate"`
		Taxonomies        string `bson:"taxonomies"`
		TeaserText        string `bson:"teaserText"`
		ThumbnailImageURL string `bson:"thumbnailImageURL"`
		Title             string `bson:"title"`
		OptaMatchId       string `bson:"optaMatchId"`
		LastUpdateDate    string `bson:"lastUpdateDate"`
		IsPublished       string `bson:"isPublished"`
	}

	NewsArticleDB struct {
		ArticleURL        string `bson:"articleUrl"`
		NewsArticleID     int64  `bson:"_id"`
		PublishDate       string `bson:"publishDate"`
		Taxonomies        string `bson:"taxonomies"`
		TeaserText        string `bson:"teaserText"`
		Subtitle          string `bson:"subtitle"`
		ThumbnailImageURL string `bson:"thumbnailImageURL"`
		Title             string `bson:"title"`
		BodyText          string `bson:"bodyText"`
		GalleryImageURLs  string `bson:"galleryImageURLs"`
		VideoURL          string `bson:"videoURL"`
		OptaMatchId       string `bson:"optaMatchId"`
		LastUpdateDate    string `bson:"lastUpdateDate"`
		IsPublished       string `bson:"isPublished"`
	}
)
