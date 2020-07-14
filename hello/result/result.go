package result

import (
	"../parser"
)

type Result struct {
	Articles         []parser.Item
	ContentMarketing []parser.Item
	Ads              []Ads
}

type Ads struct {
	Text string `default0:"No result for content marketing"`
}

func (res Result) PrepareResponse(articles []parser.Item, contentMarketing []parser.Item, page int) Result {

	//Huh, I feel sry for this part
	if len(articles) >= page*5 {
		firstKey := page * 5
		secondKey := firstKey + 5

		if len(articles) >= firstKey+5 {
			res.Articles = articles[firstKey:secondKey]
		} else {
			res.Articles = articles[firstKey:len(articles)]
		}
	}

	if len(contentMarketing) >= page {
		res.ContentMarketing = contentMarketing[page : page+1]
	} else {

	}

	return res
}
