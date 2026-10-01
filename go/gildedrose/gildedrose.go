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

// GetOrDefaulter is implemented by map-like types that can return a fallback
// value for missing keys.
type GetOrDefaulter[K comparable, V any] interface {
	GetOrDefault(key K, defaultValue V) V
}

// Map extends the builtin map with GetOrDefault.
type Map[K comparable, V any] map[K]V

func (m Map[K, V]) GetOrDefault(key K, defaultValue V) V {
	if v, found := m[key]; found {
		return v
	}
	return defaultValue
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

var qualityHandlerMap = Map[string, func(item *Item)]{
	agedBrie:        incrementQualityByOne,
	conjured:        decrementQualityByTwo,
	backstagePasses: updateBackstagePassQuality,
	sulfuras:        itemNoOp,
}

var postSellInHandlerMap = Map[string, func(item *Item)]{
	agedBrie:        incrementQualityByOne,
	conjured:        decrementQualityByTwo,
	backstagePasses: zeroQuality,
	sulfuras:        itemNoOp,
}

func decrementSellInByOne(item *Item) {
	item.SellIn--
}

var sellInHandlerMap = Map[string, func(item *Item)]{
	sulfuras: itemNoOp,
}

var clampMap = Map[string, func(item *Item)]{
	sulfuras: itemNoOp,
}

func UpdateQuality(items []*Item) {
	for _, item := range items {
		category := itemCategory(item.Name)

		qualityHandlerMap.GetOrDefault(category, decrementQualityByOne)(item)

		sellInHandlerMap.GetOrDefault(category, decrementSellInByOne)(item)

		if item.SellIn < 0 {
			postSellInHandlerMap.GetOrDefault(category, decrementQualityByOne)(item)
		}

		clampMap.GetOrDefault(category, clampQualityZeroToFifty)(item)
	}

}
