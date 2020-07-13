package main

import (
	"./parser"
	"fmt"
)

type Result struct {
	Articles         []parser.Item
	ContentMarketing []parser.Item
}

func main() {

	parser := parser.Parser{}

	articlesData := parser.GET("https://storage.googleapis.com/aller-structure-task/articles.json")

	fmt.Println("Articles Status", articlesData.Status)

	marketingData := parser.GET("https://storage.googleapis.com/aller-structure-task/contentmarketing.json")

	fmt.Println("Marketing Status", marketingData.Status)

	//fmt.Println(marketingData.Response.Items[6])

	result := Result{}

	result.Articles = articlesData.Response.Items[1:4]

}
