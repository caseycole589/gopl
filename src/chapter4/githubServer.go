package main
import (
	// "os"
	"log"
	"html/template"
	"github"
	"net/http"	
	"fmt"
	"io/ioutil"
	// "strings"
)

type Page struct {
	Title string
	Body []byte
}
func main() {
	http.HandleFunc("/",home)
	http.HandleFunc("/repo",handler)
	log.Fatal(http.ListenAndServe("localhost:8000",nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	
    // for k, v := range r.Form {
    //     fmt.Println("key:", k)
    //     fmt.Println("val:", strings.Join(v, ""))
    // }
	result,err := github.SearchIssues(r.Form["repo"])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(r.FormValue("choice"))
  	t:= template.Must(template.ParseFiles(r.FormValue("choice") + ".html"))
  	t.Execute(w, result)
	
  
 }
func home(w http.ResponseWriter, r *http.Request) {
	t:= template.Must(template.ParseFiles("home.html"))
    t.Execute(w,nil)
}
//load html template from current working directory
func loadPage (title string) (*Page, error) {
	filename := title 
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return &Page{Title: title, Body: body}, nil
}