package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Parser struct {
	Status   int      `json:"httpStatus"`
	Response Response `json:"response"`
}

func (pars *Parser) get(uri string) Parser {
	response, err := http.Get(uri)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject Parser
	json.Unmarshal(responseData, &responseObject)

	fmt.Println(responseObject.Status)
	fmt.Println(len(responseObject.Response.Items))

	for i := 0; i < len(responseObject.Response.Items); i++ {
		fmt.Println(responseObject.Response.Items[i].Url)

	}

	return responseObject
}

type Response struct {
	Items []Item `json:"items"`
}

type Item struct {
	Type         string `json:"type"`
	HarvesterId  string `json:"harvesterId"`
	CerebroScore int    `json:"cerebro-score"`
	Url          string `json:"url"`
	Title        string `json:"title"`
	CleanImage   string `json:"cleanImage"`
}

func main() {

	parser := Parser{}
	parser.get("https://storage.googleapis.com/aller-structure-task/articles.json")

	//articlesUri, exists := os.LookupEnv("https://storage.googleapis.com/aller-structure-task/articles.json")
	//
	//if exists {
	//    parser.get(articlesUri)
	//
	//    //  fmt.Println(articlesData)
	//}

}
