package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/kamicyun/internship-discord-monitor/internal/config"
	"github.com/kamicyun/internship-discord-monitor/internal/jobs"
)

func main() {
	go serveHealthCheck()

	log.Println("Loading config...")
	cfg := config.Load()

	ticker := time.NewTicker(1 * time.Hour)
	defer ticker.Stop()

	for {
		log.Println("Starting LinkedIn scraping cycle...")
		var allJobs []jobs.Job
		seen := make(map[string]struct{})

		for _, keyword := range cfg.Keywords {
			log.Printf("Scraping keyword: %s", keyword)
			newJobs, err := jobs.MonitorLinkedin(cfg.Delay, keyword)
			if err != nil {
				log.Printf("Error scraping LinkedIn for keyword '%s': %v", keyword, err)
				continue
			}
			allJobs = appendUniqueJobs(allJobs, newJobs, seen)
		}

		if err := jobs.SendDiscordEmbeds(cfg.LinkedinWH, allJobs); err != nil {
			log.Printf("Error sending Discord webhook: %v", err)
		} else {
			log.Printf("Successfully sent %d Unqiue LinkedIn jobs", len(allJobs))
		}

		log.Println("Cycle complete. Sleeping for 1 hour.")
		<-ticker.C
	}
}

func serveHealthCheck() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Bot is running.")
	})
	port := getPort()
	log.Printf("Starting HTTP server on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("HTTP server failed: %v", err)
	}
}

func getPort() string {
	if port := os.Getenv("PORT"); port != "" {
		return port
	}
	return "8080"
}

func appendUniqueJobs(base []jobs.Job, toAdd []jobs.Job, seen map[string]struct{}) []jobs.Job {
	for _, job := range toAdd {
		parsedURL, err := url.Parse(job.Link)
		if err != nil {
			continue // skip invalid URLs
		}
		parsedURL.RawQuery = "" // strip query params
		cleanURL := parsedURL.String()

		if _, exists := seen[cleanURL]; !exists {
			job.Link = cleanURL // optionally sanitize original link too
			base = append(base, job)
			seen[cleanURL] = struct{}{}
		}
	}
	return base
}
