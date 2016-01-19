//issues prints a table of git hub issue matching search term
package main

import	(
	"fmt"
	"log"
	
	"xkcdLibary"
)
func main() {
	result,err := xkcdLibary.GetAllComicsIndexes();
	if err != nil {
		log.Fatal(err);
	}

	fmt.Println(result)
	// for _, comic := range result.ComicsList {
	// 	fmt.Println(comic);
	// } 
}