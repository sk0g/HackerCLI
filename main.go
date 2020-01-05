package main

func main() {
	postIDs := getHNPostIDs(bestStoriesUrl)

	for i := 0; i <= 5 && i < len(postIDs); i++ {
		go func() {
			postDetails := getDetailsForPost(postIDs[i])

			// TODO: display post
		}()
	}
}
