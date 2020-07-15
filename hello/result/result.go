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
	Text string
}

func (res Result) PrepareResponse(articles []parser.Item, contentMarketing []parser.Item, page int, limit int) Result {
	articlesLength := len(articles)
	contentMarketingLength := len(contentMarketing)

	if articlesLength >= page*limit {
		firstKey := page * limit
		secondKey := firstKey + limit

		/*
		   I have created this check for case when we have only 13 articles,
		   but user is asking range from 10 to 15. So to escape a error and I will return from 10 to 13 ( 3 items ).
		   But that looks awful, because of if/else variable scopes in Go, maybe I could get advice from you how setup it better
		*/

		if articlesLength >= firstKey+limit {
			res.Articles = articles[firstKey:secondKey]
		} else {
			res.Articles = articles[firstKey:articlesLength]
		}
	} else {
		res.Ads = append(res.Ads, Ads{"No articles available"})
	}

	if contentMarketingLength >= page {
		res.ContentMarketing = contentMarketing[page : page+1]
	} else {
		res.Ads = append(res.Ads, Ads{"No content marketing available"})
	}

	return res
}
