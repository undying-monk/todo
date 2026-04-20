package api

import (
	"log"
	"net/http"
	"strconv"

	"github.com/dungtc/aggregate-rss/backend/internal/database"
	"github.com/dungtc/aggregate-rss/backend/internal/models"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// CORS Setup
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	api := r.Group("/api")
	{
		api.GET("/articles", GetArticles)
		api.GET("/articles/search", SearchArticles)
	}

	return r
}

func GetArticles(c *gin.Context) {
	cursor := c.Query("cursor")
	category := c.Query("category")
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "200"))

	if limit < 1 || limit > 200 {
		limit = 200
	}

	log.Println("GetArticles", cursor, category, limit)

	var articles []models.Article
	query := database.DB.Order("id desc").Limit(limit)

	if category != "" {
		query = query.Where("category = ?", category)
	}

	if cursor != "" {
		query = query.Where("id < ?", cursor)
	}

	result := query.Find(&articles)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	var nextCursor string
	if len(articles) == limit {
		nextCursor = articles[len(articles)-1].ID
	}

	c.JSON(http.StatusOK, gin.H{
		"data":     articles,
		"cursor":   nextCursor,
		"limit":    limit,
		"category": category,
	})
}

func SearchArticles(c *gin.Context) {
	q := c.Query("q")
	cursor := c.Query("cursor")
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "200"))

	if limit < 1 || limit > 200 {
		limit = 200
	}

	var articles []models.Article
	query := database.DB.Order("id desc")

	if q != "" {
		searchStr := "%" + q + "%"
		query = query.Where("title ILIKE ? OR content_snippet ILIKE ?", searchStr, searchStr)
	}
	if cursor != "" {
		query = query.Where("id < ?", cursor)
	}

	result := query.Limit(limit).Find(&articles)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	var nextCursor string
	if len(articles) == limit {
		nextCursor = articles[len(articles)-1].ID
	}

	c.JSON(http.StatusOK, gin.H{
		"data":   articles,
		"cursor": nextCursor,
		"limit":  limit,
		"query":  q,
	})
}
