package bybit

import (
	"encoding/json"
	"github.com/google/go-querystring/query"
	"net/url"
)

// V5SpotMarginTradeServiceI :
type V5SpotMarginTradeServiceI interface {
	GetTradeState() (*V5GetTradeStatResponse, error)
	SetLeverage(param V5SpotSetLeverageParam) (*CommonV5Response, error)
	ChangeSwitchMode(param V5SpotChangeSwitchModeParam) (*V5SpotChangeSwitchModeResponse, error)
	GetTradeData(param V5GetTradeDataParam) (*V5GetTradeDataResponse, error)
}

// V5SpotMarginTradeService :
type V5SpotMarginTradeService struct {
	client *Client
}

type V5GetTradeDataParam struct {
	VipLevel string `json:"vipLevel"`
	Currency string `json:"currency"`
}

type V5GetTradeDataResponse struct {
	CommonV5Response `json:",inline"`
	Result           V5GetTradeDataResult `json:"result"`
}

// V5GetTradeDataResult :
type V5GetTradeDataResult struct {
	VipCoinList []V5GetTradeDataVipCoinListItem `json:"vipCoinList"`
}

// V5GetTradeDataVipCoinListItem :
type V5GetTradeDataVipCoinListItem struct {
	List     []V5GetTradeDataVipCoinListInnerItem `json:"list"`
	VipLevel string                               `json:"vipLevel"`
}

// V5GetTradeDataVipCoinListInnerItem :
type V5GetTradeDataVipCoinListInnerItem struct {
	Borrowable         bool   `json:"borrowable"`
	CollateralRatio    string `json:"collateralRatio"`
	Currency           string `json:"currency"`
	HourlyBorrowRate   string `json:"hourlyBorrowRate"`
	LiquidationOrder   string `json:"liquidationOrder"`
	MarginCollateral   bool   `json:"marginCollateral"`
	MaxBorrowingAmount string `json:"maxBorrowingAmount"`
}

func (s V5SpotMarginTradeService) GetTradeData(param V5GetTradeDataParam) (*V5GetTradeDataResponse, error) {
	res := V5GetTradeDataResponse{}

	queryString, err := query.Values(param)
	if err != nil {
		return nil, err
	}

	if err := s.client.getPublicly("/v5/spot-margin-trade/data", queryString, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

type V5GetTradeStatResponse struct {
	CommonV5Response  `json:",inline"`
	SpotLeverage      string `json:"spotLeverage"`
	SpotMarginMode    string `json:"spotMarginMode"`
	EffectiveLeverage string `json:"effectiveLeverage"`
}

// GetTradeState :
func (s V5SpotMarginTradeService) GetTradeState() (*V5GetTradeStatResponse, error) {
	var (
		res   V5GetTradeStatResponse
		query = make(url.Values)
	)

	if err := s.client.getV5Privately("/v5/spot-margin-trade/state", query, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// V5SpotSetLeverageParam :
type V5SpotSetLeverageParam struct {
	Leverage string `json:"leverage"`
}

func (s V5SpotMarginTradeService) SetLeverage(param V5SpotSetLeverageParam) (*CommonV5Response, error) {
	res := CommonV5Response{}

	body, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}

	if err := s.client.postV5JSON("/v5/spot-margin-trade/set-leverage", body, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

type V5SpotChangeSwitchModeParam struct {
	SpotMarginMode string `json:"spotMarginMode"`
}

type V5SpotChangeSwitchModeResponse struct {
	CommonV5Response `json:",inline"`
	SpotMarginMode   string `json:"spotMarginMode"`
}

func (s V5SpotMarginTradeService) ChangeSwitchMode(param V5SpotChangeSwitchModeParam) (*V5SpotChangeSwitchModeResponse, error) {
	res := V5SpotChangeSwitchModeResponse{}

	body, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}

	if err := s.client.postV5JSON("/v5/spot-margin-trade/switch-mode", body, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
