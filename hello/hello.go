package main

import (
	"./parser"
	"./result"
	"fmt"
)

func main() {

	const PAGE = 500

	parser := parser.Parser{}

	articlesData := parser.GET("https://storage.googleapis.com/aller-structure-task/articles.json").Response.Items

	marketingData := parser.GET("https://storage.googleapis.com/aller-structure-task/contentmarketing.json").Response.Items

	result := result.Result{}

	response := result.PrepareResponse(articlesData, marketingData, PAGE)

	fmt.Println(response)

}
