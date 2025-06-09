package main

import (
	"fmt"

	"github.com/kamicyun/internship-discord-monitor/internal/jobs"
)

func main() {
	// cfg := config.Load()

	rs, _ := jobs.MonitorLinkedin(1000, "software Engineer")
	fmt.Print(rs)
}
