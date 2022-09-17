package search

import "testing"

func TestSearchByKeyword(t *testing.T) {
	repoList := SearchByKeyword([]string{"one", "two", "three"})
	if repoList[0] != "myrepository" {
		t.Fail()
	}
}
