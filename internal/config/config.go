package config

import (
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	DiscordToken    string
	PostgreSQLToken string
	Delay           uint32
	Keywords        []string
	LinkedinWH      string
	GlassdoorWH     string
}

func Load() *Config {
	_ = godotenv.Load(".env")

	// Parse Delay to uint32
	delayStr := os.Getenv("DELAY")
	delay32, err := strconv.ParseUint(delayStr, 10, 32)
	if err != nil {
		panic("Invalid DELAY .env value: " + err.Error())
	}

	envKeywords := os.Getenv("KEYWORDS")
	keywords := strings.Split(envKeywords, ",")

	return &Config{
		PostgreSQLToken: os.Getenv("POSTGRESQL_TOKEN"),
		Delay:           uint32(delay32),
		Keywords:        keywords,
		LinkedinWH:      os.Getenv("LINKEDINWH"),
		GlassdoorWH:     os.Getenv("GLASSDOORWH"),
	}
}
