package controllers

import (
	"testing"
)

func TestPreparePrompt(t *testing.T) {
	// Write a test case for the function preparePrompt
	// The test case should check if the function returns the correct prompt
	// when brand is empty
	// The test case should check if the function returns the correct prompt
	// when brand is not empty

	// prompt := url.QueryEscape()
	// fmt.Println("URL encoded request is: ", prompt)

	got := preparePrompt("Save%2025%%20on%20selected%20kitchen%20appliances%20when%20using%20this%20ao%20discount%20code", 2, "Woman & Home")
	want := "Give me 2 alternate text for 'Save 25% on selected kitchen appliances when using this ao discount code' in the style of Woman & Home"
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}