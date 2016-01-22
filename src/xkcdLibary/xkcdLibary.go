package xkcdLibary

import (
	"encoding/json"
	"net/http"
	"strings"
	"strconv"
	"fmt"
	"io/ioutil"
)

type Comic struct {
	Num int
	Url string
	Year string
	Title string 
}

type ComicList struct {
	Comics []Comic
}

func GetTranscript(url string) (string,error) {
	var transcript struct{ Transcript string}
	resp, err := http.Get(url);
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return "error", err
	}

	if err := json.NewDecoder(resp.Body).Decode(&transcript); err != nil {
		resp.Body.Close()
		return "error", err
	}
	resp.Body.Close()
	fmt.Println(transcript.Transcript)
	return transcript.Transcript,nil
}

func GetAllComicsIndexes() (*ComicList, error) {

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

	latest.Url = q;
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
		resp, err := http.Get(strings.Join(resquest,""))

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
		next.Url = strings.Join(resquest,"")
		comicList[i] = next
	}
	var x ComicList 
	x.Comics = comicList
	fmt.Println(x.Comics)
	return &x, nil
}

func GetComics() (*ComicList, error) {
	
	file,err := ioutil.ReadFile("xkcdIndex.json")
	if err != nil {
		return nil, err
	}
	
	var comics ComicList
	json.Unmarshal(file, &comics)
	return &comics, nil
}