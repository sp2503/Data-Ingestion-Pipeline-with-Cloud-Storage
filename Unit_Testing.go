package transform // this is my transform package  it handles converting/adding extra fields to data

import (
    "testing"          // this is Go’s built-in testing library  lets me write test cases
    "yourmodule/models" // importing the Post model structure so can use it in the test
)

// This is a test function to check if the Transform() function is working as expected
func TestTransform(t *testing.T) {
    // creating a sample input  just one post with basic fields
    sample := []models.Post{
        {
            UserID: 1,
            ID:     1,
            Title:  "test",
            Body:   "body",
        },
    }

    //  call my Transform function with the sample input
    // pass "test_source" as the source name expect it to add
    result := Transform(sample, "test_source")

    //  check if the Source field in the result actually became "test_source"
    // If not, show an error with what actually got
    if result[0].Source != "test_source" {
        t.Errorf("Expected source to be 'test_source', but got %s", result[0].Source)
    }
}


