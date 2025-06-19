package jobs

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type EmbedField struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Inline bool   `json:"inline"`
}

type EmbedFooter struct {
	Text    string `json:"text,omitempty"`
	IconURL string `json:"icon_url,omitempty"` // Optional
}

type Embed struct {
	Title       string       `json:"title"`
	Description string       `json:"description,omitempty"`
	URL         string       `json:"url,omitempty"`
	Color       int          `json:"color,omitempty"`
	Fields      []EmbedField `json:"fields,omitempty"`
	Timestamp   string       `json:"timestamp,omitempty"`
	Footer      *EmbedFooter `json:"footer,omitempty"` // ← ADD THIS
}

type WebhookPayload struct {
	Content   string  `json:"content,omitempty"`
	Username  string  `json:"username,omitempty"`
	AvatarURL string  `json:"avatar_url,omitempty"`
	Embeds    []Embed `json:"embeds"`
}

const maxEmbedsPerMessage = 10
const embedColor = 0x000000
const rateLimitDelay = 4 * time.Second

func SendDiscordEmbeds(webhookURL string, jobs []Job) error {
	var (
		batch      []Embed
		shouldPing bool
	)

	for i, job := range jobs {
		// Prepare the embed
		embed := Embed{
			Title: job.Title,
			URL:   job.Link,
			Color: embedColor,
			Fields: []EmbedField{
				{Name: "Company", Value: job.Company, Inline: true},
				{Name: "Location", Value: job.Location, Inline: true},
				{Name: "Posted", Value: job.Time, Inline: true},
			},
			Timestamp: time.Now().UTC().Format(time.RFC3339),
			Footer: &EmbedFooter{
				Text: "CYUN v0.2",
			},
		}
		if !isBlacklisted(job.Company) {
			batch = append(batch, embed)
		}

		// Trigger ping if prestigious company found
		if !shouldPing && isPrestigious(job.Company) {
			shouldPing = true
		}

		// Send every maxEmbedsPerMessage or on final iteration
		isLast := i == len(jobs)-1
		if len(batch) == maxEmbedsPerMessage || isLast {
			if err := sendWebhook(webhookURL, batch, shouldPing); err != nil {
				return err
			}
			batch = nil
			shouldPing = false
			if !isLast {
				time.Sleep(rateLimitDelay)
			}
		}
	}

	return nil
}

func sendWebhook(url string, embeds []Embed, ping bool) error {
	payload := WebhookPayload{
		Username: "Job Monitor",
		Embeds:   embeds,
	}
	if ping {
		payload.Content = "@everyone"
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal webhook payload: %w", err)
	}

	resp, err := http.Post(url, "application/json", bytes.NewReader(body))
	if err != nil {
		return fmt.Errorf("failed to send webhook: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		return fmt.Errorf("webhook failed: status %d", resp.StatusCode)
	}

	return nil
}

func isPrestigious(company string) bool {
	_, ok := TopTechCompanies[strings.ToLower(company)]
	return ok
}

func isBlacklisted(company string) bool {
	_, ok := Blacklist[strings.ToLower(company)]
	return ok
}
