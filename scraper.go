package main

import (
    "fmt"
    "net/http"
    "time"
    "encoding/json"
)


type names struct{
	
}

func CheckError(e error) {
	if e != nil{
		panic(e)
	}
}

func scrape(url string)  {
	
}

func user(w http.ResponseWriter, r *http.Request) {
	pid := r.URL.Query().Get("t")
	if pid == "" {
		http.Error(w, "Error has Occur", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Expires", "Mon, 26 Jul 1997 05:00:00 GMT")
	w.Header().Set("Cache-Control", "no-cache, must-revalidate")
	w.Header().Set("Content-type", "application/json")
	w.Header().Set("content-language", "en")
	fmt.Fprintf(w, "Hello, World \n %#v \n\n %s", r, pid)
}

func tag(w http.ResponseWriter, r *http.Request) {
	tid := r.URL.Query().Get("t")
	first := r.URL.Query().Get("l")

	fmt.Fprintf(w, "Tag, \n %v", r)
}

func tuser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "tag and user \n %#v", w)
}

func main() {
	start := time.Now().Unix()
	// var first int
	// first = 120
	http.HandleFunc("/user", user)
	http.HandleFunc("/tag", tag)
	http.HandleFunc("/tuser", tuser)
	http.ListenAndServe(":8090", nil)
	// fmt.Printf("Enter the amount of data >> ")
	// fmt.Scanf("%d", first)
	// url := fmt.Sprintf("https://www.instagram.com/graphql/query/?query_id=17882293912014529&tag_name=india&first=%d&count=1", first)
	// resp, err := http.Get(url)
	// CheckError(err)
	// fmt.Printf("%#v", resp)
	end := time.Now().Unix()-start
	fmt.Printf("\n\nTotal time taken:  %v", end)
}