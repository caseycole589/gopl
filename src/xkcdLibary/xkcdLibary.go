package xkcdLibary

import (
	"encoding/json"
	"net/http"
	"strings"
	"fmt"
	"strconv"
)


type Comic struct {
	Num int
	Url string
	Year string
	Title string 
}

func GetAllComicsIndexes() (*[]Comic, error) {

	q := "http://xkcd.com/info.0.json"
	resp, err := http.Get(q);

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, err
	}

	var latest Comic
	if err := json.NewDecoder(resp.Body).Decode(&latest); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	
	comicList := make([]Comic,latest.Num) 
	comicList[0] = latest
	
	resquest := make([]string,3)
	resquest[0] = "http://xkcd.com/"
	resquest[2] = "/info.0.json"
	for i := 1; i < len(comicList); i++ {
		if i == 404 {
			continue
		}
		resquest[1] = strconv.Itoa(i)
		resp, err := http.Get(strings.Join(resquest,""));

		if resp.StatusCode != http.StatusOK {
			resp.Body.Close()
			return nil, err
		}
		var next Comic
		if err := json.NewDecoder(resp.Body).Decode(&next); err != nil {
			resp.Body.Close()
			return nil, err
		}
		resp.Body.Close()
		comicList[i] = next
	}
	return &comicList, nil
}