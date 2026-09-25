package gildedrose_test

import (
	"testing"

	"github.com/emilybache/gildedrose-refactoring-kata/gildedrose"
)

func Test_Foo(t *testing.T) {
	expectdName := "foo"
	var items = []*gildedrose.Item{
		{expectdName, 0, 0},
	}

	gildedrose.UpdateQuality(items)

	if items[0].Name != expectdName {
		t.Errorf("Name: Expected %s but got %s ", expectdName, items[0].Name)
	}
}
