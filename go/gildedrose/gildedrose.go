package gildedrose

import "regexp"

const (
	agedBrie        = "agedBrie"
	backstagePasses = "backstagePasses"
	sulfuras        = "sulfuras"
	conjured        = "conjured"
)

var itemCategoryMatchers = []struct {
	pattern  *regexp.Regexp
	category string
}{
	{regexp.MustCompile(`^Aged Brie$`), agedBrie},
	{regexp.MustCompile(`^Backstage passes `), backstagePasses},
	{regexp.MustCompile(`^Sulfuras, Hand of Ragnaros$`), sulfuras},
	{regexp.MustCompile(`^Conjured `), conjured},
}

func itemCategory(name string) string {
	for _, m := range itemCategoryMatchers {
		if m.pattern.MatchString(name) {
			return m.category
		}
	}
	return ""
}

type Item struct {
	Name            string
	SellIn, Quality int
}

func itemNoOp(item *Item) {
}

func incrementQualityByOne(item *Item) {
	item.Quality++
}

func zeroQuality(item *Item) {
	item.Quality = 0
}

func decrementQualityByOne(item *Item) {
	item.Quality--
}

func decrementQualityByTwo(item *Item) {
	item.Quality -= 2
}

func clampQualityZeroToFifty(item *Item) {
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
	conjured:        decrementQualityByTwo,
	backstagePasses: updateBackstagePassQuality,
	sulfuras:        itemNoOp,
}

var postSellInHandlerMap = map[string]func(item *Item){
	agedBrie:        incrementQualityByOne,
	conjured:        decrementQualityByTwo,
	backstagePasses: zeroQuality,
	sulfuras:        itemNoOp,
}

func decrementSellInByOne(item *Item) {
	item.SellIn--
}

var sellInHandlerMap = map[string]func(item *Item){
	sulfuras: itemNoOp,
}

var clampMap = map[string]func(item *Item){
	sulfuras: itemNoOp,
}

func UpdateQuality(items []*Item) {
	for _, item := range items {
		category := itemCategory(item.Name)

		qualityOp, found := qualityHandlerMap[category]
		if !found {
			qualityOp = decrementQualityByOne
		}
		qualityOp(item)

		sellInOp, found := sellInHandlerMap[category]
		if !found {
			sellInOp = decrementSellInByOne
		}
		sellInOp(item)

		if item.SellIn < 0 {
			postSellInQualityOp, found := postSellInHandlerMap[category]
			if !found {
				postSellInQualityOp = decrementQualityByOne
			}
			postSellInQualityOp(item)
		}

		clampOp, found := clampMap[category]
		if !found {
			clampOp = clampQualityZeroToFifty
		}
		clampOp(item)
	}

}
