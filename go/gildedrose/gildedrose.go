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

func UpdateQuality(items []*Item) {
	for _, item := range items {

		if slices.Contains(specialItems, item.Name) {
			if item.Quality < 50 {
				item.Quality = item.Quality + 1
				if item.Name == backstagePasses {
					if item.SellIn < 11 {
						if item.Quality < 50 {
							item.Quality = item.Quality + 1
						}
					}
					if item.SellIn < 6 {
						if item.Quality < 50 {
							item.Quality = item.Quality + 1
						}
					}
				}
			}
		} else {
			if item.Quality > 0 {
				item.Quality = item.Quality - 1
			}
		}

		if item.Name != sulfuras {
			item.SellIn = item.SellIn - 1
		}

		if item.SellIn < 0 {
			if item.Name != agedBrie {
				if item.Name != backstagePasses {
					if item.Quality > 0 {
						if item.Name != sulfuras {
							item.Quality = item.Quality - 1
						}
					}
				} else {
					item.Quality = item.Quality - item.Quality
				}
			} else {
				if item.Quality < 50 {
					item.Quality = item.Quality + 1
				}
			}
		}
	}

}
