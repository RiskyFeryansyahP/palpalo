package valorant

type DetailStoreBundleDetailItem struct {
	ItemTypeID string `json:"ItemTypeID"`
	ItemID     string `json:"ItemID"`
	Quantity   int    `json:"Quantity"`
}

type DetailStoreBundleItem struct {
	Item            DetailStoreBundleDetailItem `json:"item"`
	BasePrice       int                         `json:"BasePrice"`
	CurrencyID      string                      `json:"CurrencyID"`
	DiscountPercent float64                     `json:"DiscountPercent"`
	DiscountedPrice float64                     `json:"DiscountedPrice"`
	IsPromoItem     bool                        `json:"IsPromoItem"`
}

type DetailStoreBundle struct {
	Items       []DetailStoreBundleItem `json:"Items"`
	ID          string                  `json:"ID"`
	DataAssetID string                  `json:"DataAssetID"`
	CurrencyID  string                  `json:"CurrencyID"`
}

type DetailStoreFeaturedBundle struct {
	Bundle                           DetailStoreBundle   `json:"Bundle"`
	Bundles                          []DetailStoreBundle `json:"Bundles"`
	BundleRemainingDurationInSeconds int                 `json:"BundleRemainingDurationInSeconds"`
}

type DetailStoreReward struct {
	ItemTypeID string `json:"ItemTypeID"`
	ItemID     string `json:"ItemID"`
	Quantity   int    `json:"Quantity"`
}

type DetailStoreSingleItemStoreOffer struct {
	OfferID          string              `json:"OfferID"`
	IsDirectPurchase bool                `json:"IsDirectPurchase"`
	StartDate        string              `json:"StartDate"`
	Cost             map[string]int      `json:"Cost"`
	Rewards          []DetailStoreReward `json:"Rewards"`
}

type DetailStoreSkinPanelLayout struct {
	SingleItemStoreOffers                      []DetailStoreSingleItemStoreOffer `json:"SingleItemStoreOffers"`
	SingleItemOffers                           []string                          `json:"SingleItemOffers"`
	SingleItemOffersRemainingDurationInSeconds int                               `json:"SingleItemOffersRemainingDurationInSeconds"`
}

type DetailStoreOffer struct {
	OfferID          string              `json:"OfferID"`
	IsDirectPurchase bool                `json:"IsDirectPurchase"`
	StartDate        string              `json:"StartDate"`
	Cost             map[string]int      `json:"Cost"`
	Rewards          []DetailStoreReward `json:"Rewards"`
}

type DetailStoreUpgradeCurrencyOffer struct {
	Offer            DetailStoreOffer `json:"Offer"`
	OfferID          string           `json:"OfferID"`
	StorefrontItemID string           `json:"StorefrontItemID"`
}

type DetailStoreUpgradeCurrencyStore struct {
	UpgradeCurrencyOffers []DetailStoreUpgradeCurrencyOffer `json:"UpgradeCurrencyOffers"`
}

type DetailStoreBonusStoreOffer struct {
	Offer           DetailStoreOffer `json:"Offer"`
	BonusOfferID    string           `json:"BonusOfferID"`
	DiscountPercent float64          `json:"DiscountPercent"`
	DiscountCosts   map[string]int   `json:"DiscountCosts"`
	IsSeen          bool             `json:"IsSeen"`
}

// DetailStoreBonusStore
// this is a night market use case
type DetailStoreBonusStore struct {
	BonusStoreOffers []DetailStoreBonusStoreOffer `json:"BonusStoreOffers"`
}

type DetailStoreResponse struct {
	FeaturedBundle       DetailStoreFeaturedBundle       `json:"FeaturedBundle"`
	SkinsPanelLayout     DetailStoreSkinPanelLayout      `json:"SkinsPanelLayout"`
	UpgradeCurrencyStore DetailStoreUpgradeCurrencyStore `json:"UpgradeCurrencyStore"`
	BonusStore           DetailStoreBonusStore           `json:"BonusStore,omitempty"`
}
