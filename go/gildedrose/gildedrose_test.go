package gildedrose_test

import (
	"testing"

	"github.com/emilybache/gildedrose-refactoring-kata/gildedrose"
)

const (
	agedBrie        = "Aged Brie"
	backstagePasses = "Backstage passes to a TAFKAL80ETC concert"
	sulfuras        = "Sulfuras, Hand of Ragnaros"
)

func Test_UpdateQuality(t *testing.T) {
	type testItem struct {
		*gildedrose.Item
		nextSellIn  int
		nextQuality int
	}
	type testCase []testItem

	tests := map[string]testCase{
		"normal": {
			{Item: &gildedrose.Item{Name: "foo", SellIn: 5, Quality: 10}, nextSellIn: 4, nextQuality: 9},
			{Item: &gildedrose.Item{Name: "foo1", SellIn: 0, Quality: 10}, nextSellIn: -1, nextQuality: 8},
			{Item: &gildedrose.Item{Name: "foo2", SellIn: 0, Quality: 0}, nextSellIn: -1, nextQuality: 0},
		},
		"aged brie": {
			{Item: &gildedrose.Item{Name: agedBrie, SellIn: 5, Quality: 10}, nextSellIn: 4, nextQuality: 11},
			{Item: &gildedrose.Item{Name: agedBrie, SellIn: 0, Quality: 10}, nextSellIn: -1, nextQuality: 12},
			{Item: &gildedrose.Item{Name: agedBrie, SellIn: 5, Quality: 50}, nextSellIn: 4, nextQuality: 50},
			{Item: &gildedrose.Item{Name: agedBrie, SellIn: 0, Quality: 49}, nextSellIn: -1, nextQuality: 50},
		},
		"sulfuras": {
			{Item: &gildedrose.Item{Name: sulfuras, SellIn: 5, Quality: 50}, nextSellIn: 5, nextQuality: 50},
			{Item: &gildedrose.Item{Name: sulfuras, SellIn: -1, Quality: 10}, nextSellIn: -1, nextQuality: 10},
		},
		"backstage passes": {
			{Item: &gildedrose.Item{Name: backstagePasses, SellIn: 11, Quality: 20}, nextSellIn: 10, nextQuality: 21},
			{Item: &gildedrose.Item{Name: backstagePasses, SellIn: 10, Quality: 20}, nextSellIn: 9, nextQuality: 22},
			{Item: &gildedrose.Item{Name: backstagePasses, SellIn: 6, Quality: 20}, nextSellIn: 5, nextQuality: 22},
			{Item: &gildedrose.Item{Name: backstagePasses, SellIn: 5, Quality: 20}, nextSellIn: 4, nextQuality: 23},
			{Item: &gildedrose.Item{Name: backstagePasses, SellIn: 1, Quality: 20}, nextSellIn: 0, nextQuality: 23},
			{Item: &gildedrose.Item{Name: backstagePasses, SellIn: 0, Quality: 20}, nextSellIn: -1, nextQuality: 0},
			{Item: &gildedrose.Item{Name: backstagePasses, SellIn: 10, Quality: 49}, nextSellIn: 9, nextQuality: 50},
			{Item: &gildedrose.Item{Name: backstagePasses, SellIn: 5, Quality: 48}, nextSellIn: 4, nextQuality: 50},
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
				if item.SellIn != item.nextSellIn {
					t.Errorf("item %d (%s): SellIn: expected %d but got %d", i, item.Name, item.nextSellIn, item.SellIn)
				}
				if item.Quality != item.nextQuality {
					t.Errorf("item %d (%s): Quality: expected %d but got %d", i, item.Name, item.nextQuality, item.Quality)
				}
			}
		})
	}
}
