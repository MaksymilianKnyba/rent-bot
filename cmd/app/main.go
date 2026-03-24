package main
import (
	"log"
	"time"

	"rent-bot/internal/config"
	"rent-bot/internal/engine"
)

func main() {
	log.Println("Starting rent bot...")

	cfg, err := config.LoadConfig("config/config.yaml")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	log.Println("Config loaded successfully")

	eng := engine.New(cfg)

	log.Println("Starting loop...")

for {
    start := time.Now()

    log.Println("Running engine...")

    err := eng.Run()
    if err != nil {
        log.Println("error:", err)
    }

    log.Printf("Cycle completed in %s", time.Since(start))

    time.Sleep(time.Duration(cfg.IntervalSeconds) * time.Second)
}
}