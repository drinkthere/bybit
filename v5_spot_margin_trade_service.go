package bybit

import (
	"encoding/json"
	"net/url"
)

// V5SpotMarginTradeServiceI :
type V5SpotMarginTradeServiceI interface {
	GetTradeState() (*V5GetTradeStatResponse, error)
	SetLeverage(param V5SpotSetLeverageParam) (*CommonV5Response, error)
	ChangeSwitchMode(param V5SpotChangeSwitchModeParam) (*V5SpotChangeSwitchModeResponse, error)
}

// V5SpotMarginTradeService :
type V5SpotMarginTradeService struct {
	client *Client
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
