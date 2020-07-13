package main

import (
	"fmt"
    "github.com/joho/godotenv"
	"os"
	"net/http"
	"io/ioutil"
    "encoding/json"
)



func init() {
    // loads values from .env into the system
    if err := godotenv.Load(); err != nil {
        fmt.Println("No .env file found")
    }
}




/*  Parser (maybe move to subfolder  */

type Parser struct {
    structure struct{
        Items []struct {
            Name              string
            Count             int
            Is_required       bool
            Is_moderator_only bool
            Has_synonyms      bool
        }
    }
}

func (pars *Parser) get(uri string) string {
    var data struct {
        Items []struct {
            Name              string
            Count             int
            Is_required       bool
            Is_moderator_only bool
            Has_synonyms      bool
        }
    }


    resp, err := http.Get(uri)

    if err != nil {
        fmt.Println(err)
    }

    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)

    if err != nil {
       fmt.Println(err)
    }

    dec := json.NewDecoder(resp.Body)
    dec.Decode(&data)


    return string(body)
}





/* */



func main() {

    parser := Parser{}


    articlesUri, exists := os.LookupEnv("ARTICLES_URI")

    if exists {
          parser.get(articlesUri)

      //  fmt.Println(articlesData)
    }


}




