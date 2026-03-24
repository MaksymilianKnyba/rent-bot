package scraper

import (
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"

)

func FetchOlx() error {
	url := "https://www.olx.pl/nieruchomosci/mieszkania/wynajem/wejherowo/?search%5Bfilter_float_price%3Afrom%5D=1000&search%5Bfilter_float_price%3Ato%5D=2000"
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return err
	}

	doc.Find("div[data-cy='l-card']").Each(func(i int, s *goquery.Selection) {
	title := s.Find("h6").Text()
	price := s.Find("p[data-testid='ad-price']").Text()
	link, _ := s.Find("a").Attr("href")

	log.Println("📦 OFFER:")
	log.Println("Title:", title)
	log.Println("Price:", price)
	log.Println("Link:", link)
	log.Println("-----")
})

return nil

}