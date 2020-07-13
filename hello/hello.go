package main

import (
	"./parser"
	"fmt"
)

func main() {

	parser := parser.Parser{}

	articlesData := parser.GET("https://storage.googleapis.com/aller-structure-task/articles.json")

	fmt.Println("Articles Status", articlesData.Status)

	marketingData := parser.GET("https://storage.googleapis.com/aller-structure-task/contentmarketing.json")

	fmt.Println("Marketing Status", marketingData.Status)

}
