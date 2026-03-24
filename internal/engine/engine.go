package engine

import (
	"log"

	"rent-bot/internal/config"
	"rent-bot/internal/scraper"
)

type Engine struct {
	cfg *config.Config
}

func New(cfg *config.Config)*Engine {
	return &Engine{
		cfg: cfg,
	}
}

func (e *Engine) Run() error {
	log.Println("Running engine...")
	err := scraper.FetchOlx()
	if err != nil {
		return err
	}
	return nil
}