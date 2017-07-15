/*
* This is example for Dynamic Database. Means, if you change content of text file after 
* server start. It will affect the data which is render via api. For Static database
* Use Static Database example code
*/


package main

import (
    "fmt"
    "net/http"
    "encoding/json"
    "io/ioutil"
    "strings"
)

/*
* Checking for the Error
*/

func CheckError(e error) {
	if e != nil {
		panic(e)
	}
}

type JsonData struct {
	Word map[string]string
}

var mycontent = new(JsonData)  //Use to store data from text file

/*
* Use to handle getAllWords request and only respond to GET request
*/

func getAllWords(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Only GET Request", http.StatusMethodNotAllowed) // Sending method error on NON GET REQUEST
		return
	}
	w.Header().Set("Expires", "Mon, 26 Jul 1997 05:00:00 GMT")
	w.Header().Set("Cache-Control", "no-cache, must-revalidate")
	w.Header().Set("Content-type", "application/json")
	w.Header().Set("content-language", "en")

	/*
	* I have added because you want to make a restful api, so now you can use 
	* javascript to request it from browser and it will respond to your request.
	* 
	* If you want to use this for own website comment this below line
	*/

	w.Header().Set("Access-Control-Allow-Origin", "*") 

	chain := make(chan map[string]string)
	
	go func() {
		contents,err := ioutil.ReadFile("sampleWords.txt");
		CheckError(err)
		das := strings.Split(string(contents), "\n")
		mycontent.Word = make(map[string]string)
		for index, value := range das {
			str := fmt.Sprintf("%s%d", "Word", index+1)
			mycontent.Word[str] = strings.TrimSuffix(value, "\r")
		}

		chain <- mycontent.Word
	}()

	data := <- chain
	json, _ := json.MarshalIndent(data, "", "  ")
	fmt.Fprintf(w, "%s", json)
}

/*
* Use to handle saveAllWords request and only respond to POST request
*/

func saveAllWords(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Only POST Request", http.StatusMethodNotAllowed) //Sending method error on NON POST REQUEST
		return
	}
	go func() {
		contents,errin := ioutil.ReadFile("sampleWords.txt");
		CheckError(errin)
		err := ioutil.WriteFile("allWords.txt", contents, 0644)
		CheckError(err)
	}()
	fmt.Fprintf(w, "File allWords.txt has been created and saved data from sampleWords.txt")
}

func main() {
    http.HandleFunc("/getAllWords", getAllWords) //Register getAllWords handler
    http.HandleFunc("/saveAllWords", saveAllWords) //Register saveAllWords handler
    http.ListenAndServe(":8090", nil)
}