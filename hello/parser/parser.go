package parser

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Parser struct {
}

func (pars *Parser) get(uri string) Response {
	response, err := http.Get("http://pokeapi.co/api/v2/pokedex/kanto/")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject Response
	json.Unmarshal(responseData, &responseObject)

	fmt.Println(responseObject.Status)
	fmt.Println(len(responseObject.Items))

	return responseObject
}

type Response struct {
	Status int    `json:"httpStatus"`
	Items  []Item `json:"items"`
}

type Item struct {
	Type string `json:"type"`
}

func main() {

	parser := Parser{}

	articlesUri, exists := os.LookupEnv("ARTICLES_URI")

	if exists {
		parser.get(articlesUri)

		//  fmt.Println(articlesData)
	}

}
