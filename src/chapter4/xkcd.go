//issues prints a table of git hub issue matching search term
package main

import	(
	"fmt"
	"log"
	"encoding/json"
	"xkcdLibary"
	"os"
	"io/ioutil"
)
func main() {

	if !Exists("xkcdIndex.json") {
		buildIndex("xkcdIndex.json")
	}
	c, err := xkcdLibary.GetComics()
	if err != nil {
		log.Fatal("Error: %s",err)
	}
	for _,e := range os.Args[1:] {
		
	}
	for i, e := range c.Comics {
		if i != 404 {
			xkcdLibary.GetTranscript(e.Url)
		}
	}
}


func buildIndex(name string) {
	result,err := xkcdLibary.GetAllComicsIndexes();
	if err != nil {
		log.Fatal(err);
	}
	//pretty print json to file 
	data, err := json.MarshalIndent(*result,"","	")
	if  err != nil {
		log.Fatal("JSON marshalling failed: %s",err)
	}

	err = ioutil.WriteFile(name,data,0664)
	if err != nil {
		log.Fatal("Failed to WriteFile: %s",err)
	}
}

// Exists reports whether the named file or directory exists.
func Exists(name string) bool {
    if _, err := os.Stat(name); err != nil {
    	if os.IsNotExist(err) {
            return false
        }
    }
    return true
}