package collector

import (
	"log"
	"time"

	"github.com/dungtc/aggregate-rss/backend/internal/database"
	"github.com/dungtc/aggregate-rss/backend/internal/models"
	"github.com/mmcdole/gofeed"
	"gorm.io/gorm/clause"
	"regexp"
	"strings"
)

var sourceMetadata = map[string]struct {
	URL     string
	Favicon string
}{
	"VNExpress": {URL: "https://vnexpress.net/rss/tin-moi-nhat.rss", Favicon: "https://vne-static.zadn.vn/vnews/v1/favicon.ico"},
	"TuoiTre":   {URL: "https://tuoitre.vn/rss/tin-moi-nhat.rss", Favicon: "https://static.tuoitre.vn/tuoitre/favicon.ico"},
	"ThanhNien": {URL: "https://thanhnien.vn/rss/home.rss", Favicon: "https://thanhnien.vn/favicon.ico"},
}

func extractThumbnail(item *gofeed.Item) string {
	// 1. Check Image field
	if item.Image != nil && item.Image.URL != "" {
		return item.Image.URL
	}

	// 2. Check Enclosures
	for _, enc := range item.Enclosures {
		if strings.HasPrefix(enc.Type, "image/") {
			return enc.URL
		}
	}

	// 3. Try to extract from Description (standard for many VN RSS feeds)
	re := regexp.MustCompile(`<img[^>]+src="([^">]+)"`)
	matches := re.FindStringSubmatch(item.Description)
	if len(matches) > 1 {
		return matches[1]
	}

	return ""
}

func mapToInternalCategory(rssCategories []string) string {
	if len(rssCategories) == 0 {
		return "General"
	}
	
	cat := strings.ToLower(rssCategories[0])
	switch {
	case strings.Contains(cat, "thế giới") || strings.Contains(cat, "world"):
		return "World"
	case strings.Contains(cat, "thời sự") || strings.Contains(cat, "chính trị") || strings.Contains(cat, "politics"):
		return "Politics"
	case strings.Contains(cat, "thể thao") || strings.Contains(cat, "sports"):
		return "Sports"
	case strings.Contains(cat, "công nghệ") || strings.Contains(cat, "technology") || strings.Contains(cat, "tech"):
		return "Technology"
	case strings.Contains(cat, "giải trí") || strings.Contains(cat, "entertainment") || strings.Contains(cat, "văn hóa"):
		return "Entertainment"
	case strings.Contains(cat, "kinh doanh") || strings.Contains(cat, "business"):
		return "Business"
	default:
		return rssCategories[0]
	}
}

func CollectAllFeeds() {
	log.Println("Starting RSS feed collection...")
	fp := gofeed.NewParser()

	for source, metadata := range sourceMetadata {
		feed, err := fp.ParseURL(metadata.URL)
		if err != nil {
			log.Printf("Error parsing feed %s (%s): %v\n", source, metadata.URL, err)
			continue
		}

		var newArticles []models.Article
		for _, item := range feed.Items {
			snippet := item.Description

			pubDate := time.Now()
			if item.PublishedParsed != nil {
				pubDate = *item.PublishedParsed
			} else if item.UpdatedParsed != nil {
				pubDate = *item.UpdatedParsed
			}

			category := mapToInternalCategory(item.Categories)

			article := models.Article{
				Title:          item.Title,
				Link:           item.Link,
				Source:         source,
				Favicon:        metadata.Favicon,
				PublishedAt:    pubDate,
				ContentSnippet: snippet,
				Category:       category,
				ThumbnailURL:   extractThumbnail(item),
			}
			newArticles = append(newArticles, article)
		}

		if len(newArticles) > 0 {
			result := database.DB.Clauses(clause.OnConflict{
				Columns:   []clause.Column{{Name: "link"}},
				DoNothing: true,
			}).Create(&newArticles)
			
			if result.Error != nil {
				log.Printf("Error saving articles for %s: %v\n", source, result.Error)
			} else {
				log.Printf("Processed feed from %s.\n", source)
			}
		}
	}
	log.Println("Finished RSS feed collection.")
}
