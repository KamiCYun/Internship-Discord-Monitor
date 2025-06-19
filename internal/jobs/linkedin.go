package jobs

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func MonitorLinkedin(delay uint32, keywords string) ([]Job, error) {
	const baseURL = "https://www.linkedin.com/jobs-guest/jobs/api/seeMoreJobPostings/search"

	jobs := []Job{}

	client := &http.Client{}

	for {
		// Construct query
		params := url.Values{
			"location": {"United States"},
			"keywords": {keywords},
			"f_TPR":    {"r3600"},
			"start":    {strconv.Itoa(len(jobs) + 1)},
		}
		fullURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())

		// Build request
		req, err := http.NewRequest("GET", fullURL, nil)
		if err != nil {
			return nil, fmt.Errorf("creating request: %w", err)
		}
		setLinkedinHeaders(req)

		// Send request
		resp, err := client.Do(req)
		if err != nil {
			return nil, fmt.Errorf("sending request: %w", err)
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("reading response: %w", err)
		}

		// Parse jobs
		prevCount := len(jobs)
		parseLinkedinJobs(string(body), &jobs)

		// no more jobs
		if len(jobs) == prevCount || len(jobs) >= 1000 {
			log.Printf("Scraped %d total jobs.", len(jobs))
			break
		}

		log.Printf("Scraped %d jobs so far.", len(jobs))
		time.Sleep(time.Duration(delay) * time.Millisecond)
	}

	return jobs, nil
}

func parseLinkedinJobs(html string, jobs *[]Job) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		log.Printf("Linkedin parse error: %v", err)
		return
	}

	doc.Find("div.base-search-card").Each(func(i int, s *goquery.Selection) {
		title := strings.TrimSpace(s.Find("h3.base-search-card__title").Text())

		company := strings.TrimSpace(s.Find("h4.base-search-card__subtitle").Text())
		location := strings.TrimSpace(s.Find("span.job-search-card__location").Text())

		link, ok := s.Find("a.base-card__full-link").Attr("href")
		if !ok {
			link = ""
		}

		dateStr, ok := s.Find("time.job-search-card__listdate--new").Attr("datetime")
		if !ok {
			dateStr = ""
		}

		*jobs = append(*jobs, Job{
			Title:    title,
			Company:  company,
			Location: location,
			Link:     link,
			Time:     dateStr,
		})
	})
}

func setLinkedinHeaders(req *http.Request) {
	req.Header.Set("authority", "www.linkedin.com")
	req.Header.Set("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8")
	req.Header.Set("accept-language", "en-US,en;q=0.9")
	req.Header.Set("cache-control", "max-age=0")
	req.Header.Set("upgrade-insecure-requests", "1")
	req.Header.Set("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36")
}
