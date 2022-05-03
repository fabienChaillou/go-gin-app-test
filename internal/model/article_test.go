package model

import "testing"

// Test the function that fetches all articles
func TestGetAllArticles(t *testing.T) {
	list := GetAllArticles()

	// Check that the length of the list of articles returned is the
	// same as the length of the global variable holding the list
	if len(list) != len(articleList) {
		t.Fail()
	}

	// Check that each member is identical
	for i, v := range list {
		if v.Content != articleList[i].Content ||
			v.ID != articleList[i].ID ||
			v.Title != articleList[i].Title {

			t.Fail()
			break
		}
	}
}

// Test the function that fetche an Article by its ID
func TestGetArticleByID(t *testing.T) {
	a, err := GetArticleByID(1)

	if err != nil || a.ID != 1 || a.Title != "Article 1" || a.Content != "Article 1 body" {
		t.Fail()
	}
}
