package gildedrose

import "slices"

const (
	agedBrie        = "Aged Brie"
	backstagePasses = "Backstage passes to a TAFKAL80ETC concert"
	sulfuras        = "Sulfuras, Hand of Ragnaros"
)

var specialItems = []string{agedBrie, backstagePasses, sulfuras}

type Item struct {
	Name            string
	SellIn, Quality int
}

func itemNoOp(item *Item) {
}

func incrementQualityByOne(item *Item) {
	item.Quality++
}

func decrementQualityByOne(item *Item) {
	item.Quality--
}

func clampQuality(item *Item) {
	item.Quality = max(0, min(50, item.Quality))
}

func updateBackstagePassQuality(item *Item) {
	item.Quality++
	if item.SellIn < 11 {
		item.Quality++
	}
	if item.SellIn < 6 {
		item.Quality++
	}
}

var qualityHandlerMap = map[string]func(item *Item){
	agedBrie:        incrementQualityByOne,
	backstagePasses: updateBackstagePassQuality,
	sulfuras:        itemNoOp,
}

func UpdateQuality(items []*Item) {
	for _, item := range items {

		qualityOp, found := qualityHandlerMap[item.Name]
		if !found {
			qualityOp = decrementQualityByOne
		}
		qualityOp(item)

		if item.Name != sulfuras {
			item.SellIn = item.SellIn - 1
		}

		if item.SellIn < 0 {
			if slices.Contains(specialItems, item.Name) {
				if item.Name == agedBrie {
					item.Quality = item.Quality + 1
				}
				if item.Name == backstagePasses {
					item.Quality = 0
				}
			} else {
				item.Quality = item.Quality - 1
			}
		}

		clampQuality(item)
	}

}
