package collector

import (
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/dungtc/aggregate-rss/backend/internal/database"
	"github.com/dungtc/aggregate-rss/backend/internal/models"
	"github.com/mmcdole/gofeed"
	"gorm.io/gorm/clause"
)

var sourceMetadata = map[string]struct {
	IndexURL string
	Selector string
	Favicon  string
	BaseURL  string
}{
	"VNExpress": {
		IndexURL: "https://vnexpress.net/rss",
		Selector: "div.wrap-list-rss ul.list-rss a",
		Favicon:  "https://vne-static.zadn.vn/vnews/v1/favicon.ico",
		BaseURL:  "https://vnexpress.net",
	},
	"TuoiTre": {
		IndexURL: "https://tuoitre.vn/rss.htm",
		Selector: "div.content ul.list-rss a",
		Favicon:  "https://static.tuoitre.vn/tuoitre/favicon.ico",
		BaseURL:  "https://tuoitre.vn",
	},
	"ThanhNien": {
		IndexURL: "https://thanhnien.vn/rss.html",
		Selector: "ul.cate-content li.item a:first-child",
		Favicon:  "https://thanhnien.vn/favicon.ico",
		BaseURL:  "https://thanhnien.vn",
	},
}

func getRSSLinks(indexURL, selector, baseURL string) ([]string, error) {
	res, err := http.Get(indexURL)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, err
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}

	var links []string
	seen := make(map[string]bool)

	doc.Find(selector).Each(func(i int, s *goquery.Selection) {
		href, exists := s.Attr("href")
		if exists {
			// Handle relative URLs
			u, err := url.Parse(href)
			if err != nil {
				return
			}
			if !u.IsAbs() {
				ref, _ := url.Parse(baseURL)
				href = ref.ResolveReference(u).String()
			}

			if !seen[href] && (strings.HasSuffix(href, ".rss") || strings.HasSuffix(href, ".html") || strings.HasSuffix(href, ".htm")) {
				links = append(links, href)
				seen[href] = true
			}
		}
	})

	return links, nil
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
	log.Println("Starting comprehensive RSS feed collection...")
	fp := gofeed.NewParser()

	for source, metadata := range sourceMetadata {
		log.Printf("Discovering RSS feeds for %s...\n", source)

		// Fetch the latest published_at for this source to support incremental collection
		var latestArticle models.Article
		database.DB.Where("source = ?", source).Order("published_at desc").First(&latestArticle)
		latestTime := latestArticle.PublishedAt
		if !latestTime.IsZero() {
			log.Printf("Incremental collection for %s: checking articles published after %v\n", source, latestTime)
		}

		rssLinks, err := getRSSLinks(metadata.IndexURL, metadata.Selector, metadata.BaseURL)
		if err != nil {
			log.Printf("Error discovering feeds for %s: %v\n", source, err)
			continue
		}

		log.Printf("Found %d feeds for %s. Starting parsing...\n", len(rssLinks), source)

		for _, rssURL := range rssLinks {
			feed, err := fp.ParseURL(rssURL)
			if err != nil {
				continue
			}

			var newArticles []models.Article
			for _, item := range feed.Items {
				pubDate := time.Now()
				if item.PublishedParsed != nil {
					pubDate = *item.PublishedParsed
				} else if item.UpdatedParsed != nil {
					pubDate = *item.UpdatedParsed
				}

				// Incremental check: Skip if article is older than or equal to the latest one in DB
				if !latestTime.IsZero() && !pubDate.After(latestTime) {
					continue
				}

				snippet := item.Description
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
					log.Printf("Error saving articles from %s: %v\n", rssURL, result.Error)
				}
			}
		}
		log.Printf("Finished processing all discovered feeds for %s.\n", source)
	}
	log.Println("Finished comprehensive RSS feed collection.")
}
