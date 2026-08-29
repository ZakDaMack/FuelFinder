package govapi

type NodeFuelPrice struct {
	NodeID            string      `json:"node_id"`
	PublicPhoneNumber *string     `json:"public_phone_number"`
	TradingName       *string     `json:"trading_name"`
	FuelPrices        []FuelPrice `json:"fuel_prices"`
}

type FuelPrice struct {
	FuelType                      string  `json:"fuel_type"`
	Price                         float64 `json:"price"`
	PriceLastUpdated              string  `json:"price_last_updated"`               // ISO 8601
	PriceChangeEffectiveTimestamp string  `json:"price_change_effective_timestamp"` // ISO 8601
}

type StationNode struct {
	NodeID                      string       `json:"node_id"`
	PublicPhoneNumber           *string      `json:"public_phone_number"`
	TradingName                 string       `json:"trading_name"`
	IsSameTradingAndBrandName   bool         `json:"is_same_trading_and_brand_name"`
	BrandName                   string       `json:"brand_name"`
	TemporaryClosure            bool         `json:"temporary_closure"`
	PermanentClosure            bool         `json:"permanent_closure"`
	PermanentClosureDate        *string      `json:"permanent_closure_date"`
	IsMotorwayServiceStation    bool         `json:"is_motorway_service_station"`
	IsSupermarketServiceStation bool         `json:"is_supermarket_service_station"`
	Location                    Location     `json:"location"`
	Amenities                   []string     `json:"amenities"`
	OpeningTimes                OpeningTimes `json:"opening_times"`
	FuelTypes                   []string     `json:"fuel_types"`
}

type Location struct {
	AddressLine1 string  `json:"address_line_1"`
	AddressLine2 *string `json:"address_line_2"`
	City         string  `json:"city"`
	Country      string  `json:"country"`
	County       *string `json:"county"`
	Postcode     string  `json:"postcode"`
	Latitude     float64 `json:"latitude"`
	Longitude    float64 `json:"longitude"`
}

type OpeningTimes struct {
	UsualDays   UsualDays   `json:"usual_days"`
	BankHoliday BankHoliday `json:"bank_holiday"`
}

type UsualDays struct {
	Monday    OpeningTime `json:"monday"`
	Tuesday   OpeningTime `json:"tuesday"`
	Wednesday OpeningTime `json:"wednesday"`
	Thursday  OpeningTime `json:"thursday"`
	Friday    OpeningTime `json:"friday"`
	Saturday  OpeningTime `json:"saturday"`
	Sunday    OpeningTime `json:"sunday"`
}

type OpeningTime struct {
	Open      string `json:"open"`
	Close     string `json:"close"`
	Is24Hours bool   `json:"is_24_hours"`
}

type BankHoliday struct {
	Type      string `json:"type"`
	OpenTime  string `json:"open_time"`
	CloseTime string `json:"close_time"`
	Is24Hours bool   `json:"is_24_hours"`
}
