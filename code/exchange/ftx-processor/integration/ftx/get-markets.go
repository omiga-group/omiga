package ftx

import (
	"io/ioutil"

	"github.com/omiga-group/omiga/code/exchange/ftx-processor/integration/entities"
)

func (c Client) GetMarkets() error {
	res, err := c.http.Get(baseAPIURL + "markets")
	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	markets, err := entities.MarketsFromResponse(body)
	if err != nil {
		return err
	}

	_ = markets

	return nil
}
