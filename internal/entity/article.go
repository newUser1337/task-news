package entity

import (
	"encoding/json"
	"encoding/xml"
)

type ArticleList struct {
	XMLName        xml.Name            `xml:"NewListInformation"`
	ClubName       string              `xml:"ClubName"`
	ClubWebsiteURL string              `xml:"ClubWebsiteURL"`
	ArticleList    NewsletterNewsItems `xml:"NewsletterNewsItems"`
}

type NewsletterNewsItems struct {
	XMLName  xml.Name              `xml:"NewsletterNewsItems"`
	Articles []*NewsletterNewsItem `xml:"NewsletterNewsItem"`
}

type NewsletterNewsItem struct {
	XMLName           xml.Name `xml:"NewsletterNewsItem"`
	ArticleURL        string   `xml:"ArticleURL" bson:"articleUrl" json:"articleURL"`
	NewsArticleID     int64    `xml:"NewsArticleID" bson:"_id" json:"newsArticleID"`
	PublishDate       string   `xml:"PublishDate" bson:"publishDate" json:"publishDate"`
	Taxonomies        string   `xml:"Taxonomies" bson:"taxonomies" json:"taxonomies"`
	TeaserText        string   `xml:"TeaserText" bson:"teaserText" json:"teaserText"`
	ThumbnailImageURL string   `xml:"ThumbnailImageURL" bson:"thumbnailImageURL" json:"thumbnailImageURL"`
	Title             string   `xml:"Title" bson:"title" json:"title"`
	OptaMatchId       string   `xml:"OptaMatchId" bson:"optaMatchId" json:"optaMatchId"`
	LastUpdateDate    string   `xml:"LastUpdateDate" bson:"lastUpdateDate" json:"lastUpdateDate"`
	IsPublished       string   `xml:"IsPublished" bson:"isPublished" json:"isPublished"`
}

func (a NewsletterNewsItem) MarshalJSON() ([]byte, error) {
	newsletterNewsItem := struct {
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
	}{
		ArticleURL:        a.ArticleURL,
		NewsArticleID:     a.NewsArticleID,
		PublishDate:       a.PublishDate,
		Taxonomies:        a.Taxonomies,
		TeaserText:        a.TeaserText,
		ThumbnailImageURL: a.ThumbnailImageURL,
		Title:             a.Title,
		OptaMatchId:       a.OptaMatchId,
		LastUpdateDate:    a.LastUpdateDate,
		IsPublished:       a.IsPublished,
	}
	return json.Marshal(&newsletterNewsItem)
}

type Article struct {
	XMLName        xml.Name    `xml:"NewsArticleInformation"`
	ClubName       string      `xml:"ClubName"`
	ClubWebsiteURL string      `xml:"ClubWebsiteURL"`
	NewsArticle    NewsArticle `xml:"NewsArticle"`
}

type NewsArticle struct {
	XMLName           xml.Name `xml:"NewsArticle"`
	ArticleURL        string   `xml:"ArticleURL" bson:"articleUrl" json:"articleURL"`
	NewsArticleID     int64    `xml:"NewsArticleID" bson:"_id" json:"newsArticleID"`
	PublishDate       string   `xml:"PublishDate" bson:"publishDate" json:"publishDate"`
	Taxonomies        string   `xml:"Taxonomies" bson:"taxonomies" json:"taxonomies"`
	TeaserText        string   `xml:"TeaserText" bson:"teaserText" json:"teaserText"`
	Subtitle          string   `xml:"Subtitle" bson:"subtitle" json:"subtitle"`
	ThumbnailImageURL string   `xml:"ThumbnailImageURL" bson:"thumbnailImageURL" json:"thumbnailImageURL"`
	Title             string   `xml:"Title" bson:"title" json:"title"`
	BodyText          string   `xml:"BodyText" bson:"bodyText" json:"bodyText"`
	GalleryImageURLs  string   `xml:"GalleryImageURLs" bson:"galleryImageURLs" json:"galleryImageURLs"`
	VideoURL          string   `xml:"VideoURL" bson:"videoURL" json:"videoURL"`
	OptaMatchId       string   `xml:"OptaMatchId" bson:"optaMatchId" json:"optaMatchId"`
	LastUpdateDate    string   `xml:"LastUpdateDate" bson:"lastUpdateDate" json:"lastUpdateDate"`
	IsPublished       string   `xml:"IsPublished" bson:"isPublished" json:"isPublished"`
}

func (a NewsArticle) MarshalJSON() ([]byte, error) {
	newsArticle := struct {
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
	}{
		ArticleURL:        a.ArticleURL,
		NewsArticleID:     a.NewsArticleID,
		PublishDate:       a.PublishDate,
		Taxonomies:        a.Taxonomies,
		TeaserText:        a.TeaserText,
		Subtitle:          a.Subtitle,
		ThumbnailImageURL: a.ThumbnailImageURL,
		Title:             a.Title,
		BodyText:          a.BodyText,
		GalleryImageURLs:  a.GalleryImageURLs,
		VideoURL:          a.VideoURL,
		OptaMatchId:       a.OptaMatchId,
		LastUpdateDate:    a.LastUpdateDate,
		IsPublished:       a.IsPublished,
	}
	return json.Marshal(&newsArticle)
}
