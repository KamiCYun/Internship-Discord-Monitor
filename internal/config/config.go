package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	DiscordToken    string
	PostgreSQLToken string
	Delay           uint32
	Keywords        string
}

func Load() *Config {
	_ = godotenv.Load(".env")

	// Parse Delay to uint32
	delayStr := os.Getenv("DELAY")
	delay32, err := strconv.ParseUint(delayStr, 10, 32)
	if err != nil {
		panic("Invalid DELAY .env value: " + err.Error())
	}

	return &Config{
		DiscordToken:    os.Getenv("DISCORD_TOKEN"),
		PostgreSQLToken: os.Getenv("POSTGRESQL_TOKEN"),
		Delay:           uint32(delay32),
		Keywords:        os.Getenv("KEYWORDS"),
	}
}
