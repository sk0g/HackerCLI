package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

const (
	newStoriesUrl  = "https://hacker-news.firebaseio.com/v0/newstories.json"
	topStoriesUrl  = "https://hacker-news.firebaseio.com/v0/topstories.json"
	bestStoriesUrl = "https://hacker-news.firebaseio.com/v0/beststories.json"

	itemDetailsUrl = "https://hacker-news.firebaseio.com/v0/item/$1.json"
)

// get HN post IDs, given a URL to query for the array of post ID ints
func getHNPostIDs(url string) (postIDs []int) {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal("Error querying for posts: ", err)
		return
	}

	responseBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal("Error reading posts: ", err)
	}

	err = json.Unmarshal(responseBody, &postIDs)
	if err != nil {
		log.Fatal("Error unmarshalling JSON: ", err)
	}

	return
}

func getDetailsForPost(postID int) (postData Data) {
	url := strings.Replace(itemDetailsUrl, "$1", strconv.Itoa(postID), 1)

	res, err := http.Get(url)
	if err != nil {
		log.Fatal("Error querying for post: ", err)
		return
	}

	responseBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal("Error reading post details: ", err)
	}

	err = json.Unmarshal(responseBody, &postData)
	if err != nil {
		log.Fatal("Error unmarshalling JSON: ", err)
	}

	return
}
