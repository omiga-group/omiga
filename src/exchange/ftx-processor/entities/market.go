package entities

type Market struct {
	Name           string  `json:"name"`
	BaseCurrency   string  `json:"baseCurrency"`
	QuoteCurrency  string  `json:"quoteCurrency"`
	MinProvideSize float64 `json:"minProvideSize"`
	Ask            float64 `json:"ask"`
	Bid            int     `json:"bid"`
	Last           float64 `json:"last"`
	Price          float64 `json:"price"`
}
