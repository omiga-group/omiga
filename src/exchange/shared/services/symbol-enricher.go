package services

import "fmt"

type SymbolEnricher interface {
	GetCoinPair(symbol string) (string, string, string, string, error)
}

type symbolEnricher struct {
}

func NewSymbolEnricher() (SymbolEnricher, error) {
	return &symbolEnricher{}, nil
}

func (se *symbolEnricher) GetCoinPair(symbol string) (string, string, string, string, error) {
	if len(symbol) != 6 {
		return "", "", "", "", fmt.Errorf("invalid symbol %s. Length must be 6", symbol)
	}

	baseCode := symbol[:3]
	baseName, err := se.getCoinName(baseCode)
	if err != nil {
		return "", "", "", "", err
	}

	counterCode := symbol[3:]
	counterName, err := se.getCoinName(counterCode)
	if err != nil {
		return "", "", "", "", err
	}

	return baseCode,
		baseName,
		counterCode,
		counterName,
		nil
}

func (se *symbolEnricher) getCoinName(code string) (string, error) {
	switch code {
	case "USD":
		return "US Dollar", nil
	case "XBT":
		return "XBit", nil
	case "ADA":
		return "Cardano", nil
	case "LTC":
		return "Litecoin", nil
	case "BTC":
		return "Bitcoin", nil
	default:
		return "", fmt.Errorf("coin code %s is not recognized", code)
	}
}
