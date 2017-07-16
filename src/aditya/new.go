package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	r "gopkg.in/gorethink/gorethink.v3"
)

// CheckError for checking error and run panic if there is any
func CheckError(e error) {
	if e != nil {
		panic(e)
	}
}

var session *r.Session
var url string

func getConnected() {
	url = os.Getenv("RETHINKDB_URL")

	if url == "" {
		url = "localhost:28015"
	}

	session, _ = r.Connect(r.ConnectOpts{
		Address: url,
	})
}

func rootHandle(w http.ResponseWriter, rs *http.Request) {
	if rs.Method != "GET" {
		fmt.Fprintln(w, "Only Get Method", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "How to do it,  \n Get : use \n http://localhost:8090/get?db={name}&table={table_name} 	\n Also You can use where_name and where_age parameter in GET Request\n Post : use \n http://localhost:8090/get?db={name}&table={table_name}&personname={enter_name}&personage={age}\n Delete : use \n http://localhost:8090/delete?db={name}&table={table_name}&where_name={name} OR http://localhost:8090/delete?db={name}&table={table_name}&where_age={age}\n Update : use \n http://localhost:8090/delete?db={name}&table={table_name}&where_name={name}&name={new_name}&age={new_age} \n You can Also use where_age filter in Update method")
}

func getHandle(w http.ResponseWriter, rs *http.Request) {
	if rs.Method != "GET" {
		fmt.Fprintln(w, "Only Get Method", http.StatusBadRequest)
		return
	}
	dbname := rs.URL.Query().Get("db")
	tablename := rs.URL.Query().Get("table")
	pname := rs.URL.Query().Get("where_name")
	page := rs.URL.Query().Get("where_age")

	Data := make(map[string]interface{})
	if pname != "" {
		Data["name"] = pname
	}

	if page != "" {
		Data["age"] = page
	}

	//  _ ,err := r.DB(dbname).Run(session)
	// if err != nil {
	// 	fmt.Fprintf(w, "%s Database does not exist %#v", dbname, err)
	// 	return
	// }
	table, err := r.DB(dbname).Table(tablename).Filter(Data).Run(session)
	if err != nil {
		fmt.Fprintf(w, "%s Table Does not Exist", tablename)
		return
	}
	defer table.Close()
	var hero []interface{}
	error := table.All(&hero)
	if error != nil {
		fmt.Fprintf(w, "Error Occur %#v", error)
		return
	}
	w.Header().Set("Expires", "Mon, 26 Jul 1997 05:00:00 GMT")
	w.Header().Set("Cache-Control", "no-cache, must-revalidate")
	w.Header().Set("Content-type", "application/json")
	w.Header().Set("content-language", "en")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json, _ := json.MarshalIndent(hero, "", "  ")
	fmt.Fprintf(w, "%s", json)
}

func postHandle(w http.ResponseWriter, rs *http.Request) {
	if rs.Method != "GET" {
		fmt.Fprintln(w, "Error 400 Only Post Method", http.StatusBadRequest)
		return
	}
	dbname := rs.URL.Query().Get("db")
	tablename := rs.URL.Query().Get("table")
	pname := rs.URL.Query().Get("personname")
	page := rs.URL.Query().Get("personage")
	Data := make(map[string]interface{})

	if pname != "" {
		Data["name"] = pname
	}

	if page != "" {
		Data["age"] = page
	}
	table, err := r.DB(dbname).Table(tablename).Insert(Data).Run(session)
	if err != nil {
		fmt.Fprintf(w, "%s Table Does not Exist", tablename)
		return
	}
	defer table.Close()
	var hero []interface{}
	error := table.All(&hero)
	if error != nil {
		fmt.Fprintf(w, "Error Occur %#v", error)
		return
	}
	w.Header().Set("Expires", "Mon, 26 Jul 1997 05:00:00 GMT")
	w.Header().Set("Cache-Control", "no-cache, must-revalidate")
	w.Header().Set("Content-type", "application/json")
	w.Header().Set("content-language", "en")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json, _ := json.MarshalIndent(hero, "", "  ")
	fmt.Fprintf(w, "%s", json)
}

func deleteHandle(w http.ResponseWriter, rs *http.Request) {
	if rs.Method != "GET" {
		fmt.Fprintln(w, "Only Delete Method", http.StatusBadRequest)
		return
	}
	dbname := rs.URL.Query().Get("db")
	tablename := rs.URL.Query().Get("table")
	pname := rs.URL.Query().Get("where_name")
	page := rs.URL.Query().Get("where_age")
	Data := make(map[string]interface{})

	if pname != "" {
		Data["name"] = pname
	}

	if page != "" {
		Data["age"] = page
	}
	table, err := r.DB(dbname).Table(tablename).Filter(Data).Delete().Run(session)
	if err != nil {
		fmt.Fprintf(w, "%s Table Does not Exist", tablename)
		return
	}
	defer table.Close()
	var hero []interface{}
	error := table.All(&hero)
	if error != nil {
		fmt.Fprintf(w, "Error Occur %#v", error)
		return
	}
	w.Header().Set("Expires", "Mon, 26 Jul 1997 05:00:00 GMT")
	w.Header().Set("Cache-Control", "no-cache, must-revalidate")
	w.Header().Set("Content-type", "application/json")
	w.Header().Set("content-language", "en")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json, _ := json.MarshalIndent(hero, "", "  ")
	fmt.Fprintf(w, "%s", json)
}

func updateHandle(w http.ResponseWriter, rs *http.Request) {
	if rs.Method != "GET" {
		fmt.Fprintf(w, "Only Get Request")
		return
	}
	dbname := rs.URL.Query().Get("db")
	tablename := rs.URL.Query().Get("table")
	pname := rs.URL.Query().Get("where_name")
	page := rs.URL.Query().Get("where_age")
	name := rs.URL.Query().Get("name")
	age := rs.URL.Query().Get("age")

	Fdata := make(map[string]interface{})
	Idata := make(map[string]interface{})
	if pname != "" {
		Fdata["name"] = pname
	}
	if page != "" {
		Fdata["age"] = page
	}
	if name != "" {
		Idata["name"] = name
	}
	if age != "" {
		Idata["age"] = age
	}
	table, err := r.DB(dbname).Table(tablename).Filter(Fdata).Update(Idata).Run(session)
	if err != nil {
		fmt.Fprintf(w, "%s Table Does not Exist", tablename)
		return
	}
	defer table.Close()
	var hero []interface{}
	error := table.All(&hero)
	if error != nil {
		fmt.Fprintf(w, "Error Occur %#v", error)
		return
	}
	w.Header().Set("Expires", "Mon, 26 Jul 1997 05:00:00 GMT")
	w.Header().Set("Cache-Control", "no-cache, must-revalidate")
	w.Header().Set("Content-type", "application/json")
	w.Header().Set("content-language", "en")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json, _ := json.MarshalIndent(hero, "", "  ")
	fmt.Fprintf(w, "%s", json)

}

func main() {
	getConnected()
	http.HandleFunc("/", rootHandle)
	http.HandleFunc("/get", getHandle)
	http.HandleFunc("/post", postHandle)
	http.HandleFunc("/delete", deleteHandle)
	http.HandleFunc("/update", updateHandle)
	fmt.Printf("Server 127.0.0.1 at port 8090 has been started\n")
	http.ListenAndServe(":8090", nil)
}
