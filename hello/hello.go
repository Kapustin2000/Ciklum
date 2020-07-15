package main

import (
	"./parser"
	"./result"
	"fmt"
)

func main() {
	//Change this const value if you want to get next articles
	const PAGE = 1
	//Change this const value if you want to get more articles per call
	const LIMIT = 5

	parser := parser.Parser{}

	articlesData := parser.GET("https://storage.googleapis.com/aller-structure-task/articles.json").Response.Items

	marketingData := parser.GET("https://storage.googleapis.com/aller-structure-task/contentmarketing.json").Response.Items

	result := result.Result{}

	response := result.PrepareResponse(articlesData, marketingData, PAGE, LIMIT)

	fmt.Println(response)

}
