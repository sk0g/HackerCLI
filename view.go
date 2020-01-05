package main

import (
	"fmt"
)

func displayPost(postDetails Data) {
	fmt.Printf("Score: %v | Title: %v | Comments: %v \n",
		postDetails.Score, postDetails.Title, postDetails.CommentCount)
}
