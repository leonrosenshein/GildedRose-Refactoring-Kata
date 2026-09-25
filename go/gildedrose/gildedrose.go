package gildedrose

const (
	agedBrie        = "Aged Brie"
	backstagePasses = "Backstage passes to a TAFKAL80ETC concert"
	sulfuras        = "Sulfuras, Hand of Ragnaros"
	conjured        = "Conjured Mana Cake"
)

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

		qualityOp, found := qualityHandlerMap[item.Name]
		if !found {
			qualityOp = decrementQualityByOne
		}
		qualityOp(item)

		sellInOp, found := sellInHandlerMap[item.Name]
		if !found {
			sellInOp = decrementSellInByOne
		}
		sellInOp(item)

		if item.SellIn < 0 {
			postSellInQualityOp, found := postSellInHandlerMap[item.Name]
			if !found {
				postSellInQualityOp = decrementQualityByOne
			}
			postSellInQualityOp(item)
		}

		clampOp, found := clampMap[item.Name]
		if !found {
			clampOp = clampQualityZeroToFifty
		}
		clampOp(item)
	}

}
