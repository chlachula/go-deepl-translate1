/*
curl -k https://api-free.deepl.com/v2/translate -d auth_key=01234567-89ab-cdef-0123-456789abcdef:fx -d "text=Hello, <a href="https://wikipedia.org">world</a>!" -d "source_lang=EN" -d "target_lang=DE"

	"detected_source_language":"EN"
	"text":`Hallo Welt!

Der schnelle braune <a href="https://en.wikipedia.org/wiki/Fox">Fuchs</a> springt über den faulen Hund.`
*/
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

const my_DeepL_AuthKey = "01234567-89ab-cdef-0123-456789abcdef:fx" // get this at deepl.com
const endpoint = "https://api-free.deepl.com/v2/translate"
const textToTranslate = `Hello World! 
The quick brown <a href="https://en.wikipedia.org/wiki/Fox">fox</a> jumps over the lazy dog.`

type Tarray_type struct {
	Translation []T_type `json:"translations,omitempty"`
}
type T_type struct {
	DetectedSourceLanguage string `json:"detected_source_language,omitempty"`
	Text                   string `json:"text,omitempty"`
}

func main() {
	requestData := url.Values{}
	requestData.Set("auth_key", my_DeepL_AuthKey)
	requestData.Set("source_lang", "EN")
	requestData.Set("target_lang", "DE")
	requestData.Set("text", textToTranslate)
	r, err := http.NewRequest("POST", endpoint, strings.NewReader(requestData.Encode()))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("requestData:\n%v\n\n", requestData)

	//set headers to the request
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	fmt.Printf("request:\n%v\n\n", r)

	//send request and get the response
	client := &http.Client{}
	response, err := client.Do(r)
	if err != nil {
		log.Fatal(`error: `, err)
	}

	defer response.Body.Close()
	//Read the response body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var tArray Tarray_type
	json.Unmarshal(body, &tArray)

	for i, t := range tArray.Translation {
		fmt.Printf("response part %d:\n \"detected_source_language\":\"%s\"\n \"text\":`%s`\n", i, t.DetectedSourceLanguage, t.Text)
	}
}
