package main

func main() {
	postIDs := getHNPostIDs(bestStoriesUrl)

	for i := 0; i <= 5 && i < len(postIDs); i++ {
		postDetails := getDetailsForPost(postIDs[i])

		displayPost(postDetails)
	}

}
