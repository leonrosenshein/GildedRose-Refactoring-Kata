package gildedrose_test

import (
	"testing"

	"github.com/emilybache/gildedrose-refactoring-kata/gildedrose"
)

func Test_UpdateQuality(t *testing.T) {
	type testItem struct {
		*gildedrose.Item
		nextQuality int
		nextSellIn  int
	}
	type testCase []testItem

	tests := map[string]testCase{
		"normal": {
			{Item: &gildedrose.Item{Name: "foo", SellIn: 5, Quality: 10}, nextQuality: 9, nextSellIn: 4},
			{Item: &gildedrose.Item{Name: "foo1", SellIn: 0, Quality: 10}, nextQuality: 8, nextSellIn: -1},
			{Item: &gildedrose.Item{Name: "foo2", SellIn: 0, Quality: 0}, nextQuality: 0, nextSellIn: -1},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			items := make([]*gildedrose.Item, len(test))
			for i, item := range test {
				items[i] = item.Item
			}

			gildedrose.UpdateQuality(items)

			for i, item := range test {
				if item.Quality != item.nextQuality {
					t.Errorf("item %d (%s): Quality: expected %d but got %d", i, item.Name, item.nextQuality, item.Quality)
				}
				if item.SellIn != item.nextSellIn {
					t.Errorf("item %d (%s): SellIn: expected %d but got %d", i, item.Name, item.nextSellIn, item.SellIn)
				}
			}
		})
	}
}
