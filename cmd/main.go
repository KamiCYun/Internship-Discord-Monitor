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

	// Immediately run once on startup
	go runLinkedInCycle(cfg)
	go runGlassdoorCycle(cfg)

	// Tickers
	linkedinTicker := time.NewTicker(1 * time.Hour)
	glassdoorTicker := time.NewTicker(24 * time.Hour)
	defer linkedinTicker.Stop()
	defer glassdoorTicker.Stop()

	for {
		select {
		case <-linkedinTicker.C:
			go runLinkedInCycle(cfg)
		case <-glassdoorTicker.C:
			go runGlassdoorCycle(cfg)
		}
	}
}

func runLinkedInCycle(cfg *config.Config) {
	log.Println("Starting LinkedIn scraping cycle...")
	var allJobs []jobs.Job
	seen := make(map[string]struct{})

	for _, keyword := range cfg.Keywords {
		log.Printf("Scraping LinkedIn for keyword: %s", keyword)
		newJobs, err := jobs.MonitorLinkedin(cfg.Delay, keyword)
		if err != nil {
			log.Printf("Error scraping LinkedIn for keyword '%s': %v", keyword, err)
			continue
		}
		allJobs = appendUniqueJobs(allJobs, newJobs, seen)
	}

	if err := jobs.SendDiscordEmbeds(cfg.LinkedinWH, allJobs); err != nil {
		log.Printf("Error sending LinkedIn webhook: %v", err)
	} else {
		log.Printf("Successfully sent %d unique LinkedIn jobs", len(allJobs))
	}
}

func runGlassdoorCycle(cfg *config.Config) {
	log.Println("Starting Glassdoor scraping cycle...")
	var allJobs []jobs.Job
	seen := make(map[string]struct{})

	for _, keyword := range cfg.Keywords {
		log.Printf("Scraping Glassdoor for keyword: %s", keyword)
		newJobs, err := jobs.MonitorGlassdoor(cfg.Delay, keyword)
		if err != nil {
			log.Printf("Error scraping Glassdoor for keyword '%s': %v", keyword, err)
			continue
		}
		allJobs = appendUniqueJobs(allJobs, newJobs, seen)
	}

	if err := jobs.SendDiscordEmbeds(cfg.GlassdoorWH, allJobs); err != nil {
		log.Printf("Error sending Glassdoor webhook: %v", err)
	} else {
		log.Printf("Successfully sent %d Glassdoor jobs", len(allJobs))
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
			base = append(base, job)
			seen[cleanURL] = struct{}{}
		}
	}
	return base
}
