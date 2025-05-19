package main // Go program that runs needs to start in package "main"

import (
    "fmt"                  //  for printing messages and formatting strings
    "yourmodule/fetch"     // custom package to get data (probably from an API)
    "yourmodule/storage"   // handles saving stuff like files
    "yourmodule/transform" // clean or modify the data before saving
)

func main() {
    // Step 1: Fetch data using the function from fetch package
    posts, err := fetch.FetchPosts()
    if err != nil {
        // panic basically crashes the program and shows an error
        panic(fmt.Errorf("failed to fetch posts: %w", err))
    }

    // Step 2: Transform the posts as adding metadata like timestamp and source
    transformed := transform.Transform(posts, "placeholder_api")

    // Step 3: Saveing the transformed data into a JSON file
    err = storage.SaveToFile(transformed, "ingested_data.json")
    if err != nil {
        panic(fmt.Errorf("failed to save: %w", err))
    }

    // If everything went okay just printing a success message
    fmt.Println("Data ingestion completed successfully.")
}
