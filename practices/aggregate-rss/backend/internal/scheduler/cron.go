package scheduler

import (
	"log"
	"time"

	"github.com/dungtc/aggregate-rss/backend/internal/collector"
	"github.com/go-co-op/gocron/v2"
)

func InitCron() (gocron.Scheduler, error) {
	s, err := gocron.NewScheduler()
	if err != nil {
		return nil, err
	}

	// Schedule every 1 hour
	_, err = s.NewJob(
		gocron.DurationJob(1*time.Hour),
		gocron.NewTask(collector.CollectAllFeeds),
	)
	if err != nil {
		return nil, err
	}

	s.Start()
	log.Println("Started cron scheduler. Next run in 1 hour.")

	// Run once immediately on startup
	go collector.CollectAllFeeds()

	return s, nil
}
