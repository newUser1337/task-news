package entity

import (
	"encoding/xml"
)

type (
	ArticleListExternal struct {
		XMLName        xml.Name                     `xml:"NewListInformation"`
		ClubName       string                       `xml:"ClubName"`
		ClubWebsiteURL string                       `xml:"ClubWebsiteURL"`
		ArticleList    *NewsletterNewsItemsExternal `xml:"NewsletterNewsItems"`
	}

	NewsletterNewsItemsExternal struct {
		XMLName  xml.Name                      `xml:"NewsletterNewsItems"`
		Articles []*NewsletterNewsItemExternal `xml:"NewsletterNewsItem"`
	}

	NewsletterNewsItemExternal struct {
		XMLName           xml.Name `xml:"NewsletterNewsItem"`
		ArticleURL        string   `xml:"ArticleURL"`
		NewsArticleID     int64    `xml:"NewsArticleID"`
		PublishDate       string   `xml:"PublishDate"`
		Taxonomies        string   `xml:"Taxonomies"`
		TeaserText        string   `xml:"TeaserText"`
		ThumbnailImageURL string   `xml:"ThumbnailImageURL"`
		Title             string   `xml:"Title"`
		OptaMatchId       string   `xml:"OptaMatchId"`
		LastUpdateDate    string   `xml:"LastUpdateDate"`
		IsPublished       string   `xml:"IsPublished"`
	}

	ArticleExternal struct {
		XMLName        xml.Name             `xml:"NewsArticleInformation"`
		ClubName       string               `xml:"ClubName"`
		ClubWebsiteURL string               `xml:"ClubWebsiteURL"`
		NewsArticle    *NewsArticleExternal `xml:"NewsArticle"`
	}

	NewsArticleExternal struct {
		XMLName           xml.Name `xml:"NewsArticle"`
		ArticleURL        string   `xml:"ArticleURL"`
		NewsArticleID     int64    `xml:"NewsArticleID"`
		PublishDate       string   `xml:"PublishDate"`
		Taxonomies        string   `xml:"Taxonomies"`
		TeaserText        string   `xml:"TeaserText"`
		Subtitle          string   `xml:"Subtitle"`
		ThumbnailImageURL string   `xml:"ThumbnailImageURL"`
		Title             string   `xml:"Title"`
		BodyText          string   `xml:"BodyText"`
		GalleryImageURLs  string   `xml:"GalleryImageURLs"`
		VideoURL          string   `xml:"VideoURL"`
		OptaMatchId       string   `xml:"OptaMatchId"`
		LastUpdateDate    string   `xml:"LastUpdateDate"`
		IsPublished       string   `xml:"IsPublished"`
	}
)

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
