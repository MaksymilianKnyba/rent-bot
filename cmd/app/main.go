package main
import (
	'log'
	'time'

	'rent-bot/internal/config'
	'rent-bot/internal/engine'
)

func main() {
	log.Println("Starting rent bot...")

	cfg, err := config.LoadConfig('config/config.yaml')
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	log.Println("Config loaded successfully")

	eng := engine.New(cfg)

	log.Println("Starting loop...")

	for {
		start := time.Now()

		err := eng.Run()
		if err != nil {
			log.Printf("Error running engine: %v", err)
		}

		log.Printf("Cycle completed in %v", time.Since(start))

		time.Sleep(time.Duration(cfg.PollInterval) * time.Second)
	}
}