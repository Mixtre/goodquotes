package goodquotes

import (
	"strconv"
	"strings"

	. "github.com/Mixtre/goodquotes/types"

	"github.com/gocolly/colly/v2"
)

func Quotes(query string, page int) ScrapedData {
	scraper := colly.NewCollector()

	var scrapedData ScrapedData

	scraper.OnHTML(".mainContentContainer .mainContent .mainContentFloat", func(resp *colly.HTMLElement) {
		scrapedData.Name = resp.ChildText("h1")

		PageDiv := resp.ChildText("div.leftContainer div:nth-child(33) div")
		if len(PageDiv) > 0 {
			Pages := strings.Fields(PageDiv)
			if len(Pages) > 0 {
				Pages := Pages[len(Pages)-3]
				scrapedData.Pages, _ = strconv.ParseUint(Pages, 10, 64)
			} else {
				scrapedData.Pages = 0
			}
		}

		resp.ForEach(".leftContainer .quote", func(_ int, e *colly.HTMLElement) {
			quoteText := e.ChildText("div.quoteText")
			likes := e.ChildText("div.quoteFooter div.right a")
			likeCount, _ := strconv.ParseUint(strings.Fields(likes)[0], 10, 64)
			scrapedData.Quotes = append(scrapedData.Quotes, Quote{
				TextQuote: strings.TrimSpace(strings.Split(quoteText, "―")[0]),
				Author:    strings.TrimSpace(strings.Split(quoteText, "―")[1]),
				Tags:      e.ChildTexts("div.quoteFooter .greyText a"),
				Likes:     likeCount,
			})
		})
	})

	query = strings.ReplaceAll(query, " ", "-")
	if page > 1 {
		scraper.Visit("https://www.goodreads.com/quotes/tag/" + query + "?page=" + strconv.Itoa(page))
	} else {
		scraper.Visit("https://www.goodreads.com/quotes/tag/" + query + "?page=1")
	}
	return scrapedData
}
