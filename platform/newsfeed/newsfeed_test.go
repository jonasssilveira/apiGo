package newsfeed

import "testing"

func TestRepo_Add(t *testing.T) {
	feed := New()
	feed.Add(&Item{})
	if len(feed.GetAll()) == 0{
		t.Error("Item was not added")
	}
}

func TestRepo_GetAll(t *testing.T) {
	feed := New()
	feed.Add(&Item{})
	if len(feed.GetAll())==0{
		t.Error("Item was not added")
	}
}