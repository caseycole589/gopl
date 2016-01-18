//issues prints a table of git hub issue matching search term
package main

import	(
	"fmt"
	"log"
	"os"
	"github"
	"time"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	// fmt.Println(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	var months,year,moreThanYear github.IssuesSearchResult
	var monthTotal,yearTotal,moreThanYearTotal = 0,0,0
	fmt.Printf("%d issues:\n",result.TotalCount)
	for _, item := range result.Items {
		days := time.Since(item.CreatedAt).Hours() / 24
		if days <= 30 {
			months.Items = append(months.Items, item)
			months.TotalCount++
			monthTotal++
		} else if days  > 30 && days  < 365 {
			year.Items = append(year.Items, item)
			year.TotalCount++
			yearTotal++
		} else {
			moreThanYear.Items = append(moreThanYear.Items,item)
			moreThanYear.TotalCount++
			moreThanYearTotal++
		}
	}
	fmt.Println("less than a month")
	fmt.Println(months.TotalCount, " issues")
	for _,item := range months.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
	fmt.Println("within the last year")
	fmt.Println(year.TotalCount," issues")
	for _,item := range year.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
	fmt.Println("more than a year ago")
	fmt.Println(moreThanYear.TotalCount, " issues")
	for _,item := range moreThanYear.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
}