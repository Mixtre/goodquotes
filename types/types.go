package types

type Quote struct {
	TextQuote string
	Author    string
	Tags      []string
	Likes     uint64
}

type ScrapedData struct {
	Name   string
	Pages  uint64
	Quotes []Quote
}
