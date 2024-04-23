<p align="center">
  <img src="https://emojipedia-us.s3.dualstack.us-west-1.amazonaws.com/thumbs/160/apple/271/books_1f4da.png" width="100" alt="GoodQuotes Logo">
</p>

<h1 align="center"><b>📚 GoodQuotes 📚</b></h1>

<p align="center">A Go package for scraping quotes from Goodreads.</p>

## **Installation**

To install **GoodQuotes**, you can use `go get`. Run the following command:

```bash
go get -u github.com/Mixtre/goodquotes
```

## **Usage**

### Scraping Quotes

You can use the `Quotes` function to scrape quotes for a specific query.

```go
package main

import (
	"fmt"

	"github.com/Mixtre/goodquotes"
)

func main() {
	// Example: Scraping quotes for a specific query with default page (first page)
	query := "inspiration"
	data := goodquotes.Quotes(query, 0)

	// Printing the scraped data
	for i, quote := range data.Quotes {
		fmt.Println("Quote", i+1, ":")
		fmt.Println("Text:", quote.TextQuote)
		fmt.Println("Author:", quote.Author)
		fmt.Println("Tags:", quote.Tags)
		fmt.Println("Likes:", quote.Likes)
		fmt.Println()
	}
}
```

### Checking for More Pages

You can check if there are more pages available for a query by examining the `data.Pages` field. If there are additional pages, you can make subsequent requests to scrape quotes from those pages.

```go
package main

import (
	"fmt"

	"github.com/Mixtre/goodquotes"
)

func main() {
	// Example: Checking for more pages and scraping quotes from additional pages
	query := "inspiration"
	data := goodquotes.Quotes(query, 0)

	fmt.Println("Total Pages: ",data.Pages)

	// Checking if there are more pages
	if data.Pages > 1 {
		// Scraping quotes from page 2
		page := 2
		data = goodquotes.Quotes(query, page)

		// Printing the scraped data from page 2
		for i, quote := range data.Quotes {
			fmt.Println("Quote", i+1, ":")
			fmt.Println("Text:", quote.TextQuote)
			fmt.Println("Author:", quote.Author)
			fmt.Println("Tags:", quote.Tags)
			fmt.Println("Likes:", quote.Likes)
			fmt.Println()
		}
	} else {
		fmt.Println("No more pages available.")
	}
}
```
