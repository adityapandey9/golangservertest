/*
* This is example for Static Database. Means, even if you change content of text file after 
* server start. It will not affect the data which is render via api. For Dynamic database
* Use Dynamic Database example code
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
	WordIndex, Word []string
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

	/*
	* I can easily solve this problem and it will faster then this one, but you said to use 
	* struct. So, I am using this. It will ease, fast and better if we use only map instead
	* of struct.
	*/
	data := map[string]string {}

	for index, value := range mycontent.Word {
		data[mycontent.WordIndex[index]] = strings.TrimSuffix(value, "\r")
	}

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
 		value := strings.Join(mycontent.Word, "");
		err := ioutil.WriteFile("allWords.txt", []byte(value), 0644)
		CheckError(err)
	}()
	fmt.Fprintf(w, "File allWords.txt has been created and saved data from sampleWords.txt")
}

func main() {
	go func(){
		contents,err := ioutil.ReadFile("sampleWords.txt");
		CheckError(err)
		das := strings.Split(string(contents), "\n")
		for index, value := range das {
			str := fmt.Sprintf("%s%d", "Word", index+1)
			mycontent.WordIndex = append(mycontent.WordIndex, str)
			mycontent.Word = append(mycontent.Word, value)
		}
	}()
    http.HandleFunc("/getAllWords", getAllWords) //Register getAllWords handler
    http.HandleFunc("/saveAllWords", saveAllWords) //Register saveAllWords handler
    http.ListenAndServe(":8090", nil)
}