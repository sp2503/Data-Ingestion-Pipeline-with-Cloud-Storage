// Starting with the main package. Every Go program needs a main package to run.
package main

// importing some tools from Go’s standard library.
// These help with things like making web requests, reading responses, working with JSON, printing stuff, and handling files and time.
import (
	"encoding/json" // helps convert JSON into Go objects and back
	"fmt"           // used for printing to the console
	"io"            // lets me read the data from the HTTP response
	"net/http"      // this is needed to make HTTP requests
	"os"            // I’ll use this to write the final data into a file
	"time"          // for getting the current date & time (useful for tracking when data was fetched)
)

// This is a structure to match the shape of the data& expecting from the API.
//  checked the API  it returns JSON with userId, id, title, and body for each post.
type Post struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

// This function makes an API call to get the posts.
// returns a list of Post objects, or an error if something breaks.
func fetchPosts() ([]Post, error) {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts") // calling the API here
	if err != nil {
		return nil, err // if the call fails, just return the error
	}
	defer resp.Body.Close() // this makes sure don’t forget to close the response later

	// reading the data got from the API
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err // again, return the error if something goes wrong
	}

	// Converting the JSON response into Go structs
	var posts []Post
	err = json.Unmarshal(body, &posts) // unmarshal = decode JSON into Go objects
	if err != nil {
		return nil, err
	}

	return posts, nil // if everything went well, return the list of posts
}

// This is the main function  the starting point of the program.
func main() {
	posts, err := fetchPosts() // calling the function to get the data
	if err != nil {
		panic(err) // if there was an error,just stop everything and show the error
	}

	// printing the titles of each post to check if the data was fetched correctly
	for _, post := range posts {
		fmt.Println(post.Title)
	}
}

// this part is about transforming the data  want to add extra fields like timestamp and source info

// This is a new struct that takes everything from Post, but also adds 2 extra fields.
type TransformedPost struct {
	Post
	IngestedAt string `json:"ingested_at"` //  when got the data
	Source     string `json:"source"`      // where the data came from (like tagging the API)
}

// This function loops through each post and adds the ingested_at and source fields
func TransformPosts(posts []Post) []TransformedPost {
	var result []TransformedPost

	// getting the current time in a format that looks good in JSON
	now := time.Now().UTC().Format(time.RFC3339)

	// Going through each post one by one and creating a new TransformedPost
	for _, p := range posts {
		result = append(result, TransformedPost{
			Post:       p,
			IngestedAt: now,
			Source:     "placeholder_api", // hardcoded this since the data is coming from that API
		})
	}
	return result // returning the transformed list
}

// This function saves the transformed posts into a file on my computer
func StorePosts(posts []TransformedPost) error {
	file, err := os.Create("output.json") // trying to create a file named output.json
	if err != nil {
		return err // if file creation fails, return the error
	}
	defer file.Close() // making sure the file closes properly

	encoder := json.NewEncoder(file) // this lets  write JSON into a file
	encoder.SetIndent("", "  ")      // this makes the output look nice (pretty printed!)
	return encoder.Encode(posts)     // finally writing the data to the file
}
