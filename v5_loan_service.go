package bybit

import "github.com/google/go-querystring/query"

// V5LoanServiceI :
type V5LoanServiceI interface {
	GetLoanableData(param V5GetLoanableDataParam) (*V5GetLoanableDataResponse, error)
	GetCollateralData(param V5GetCollateralDataParam) (*V5GetCollateralDataResponse, error)
}

// V5GetCollateralDataParam :
type V5GetCollateralDataParam struct {
	VipLevel string `json:"vipLevel,omitempty"`
	Currency string `json:"currency,omitempty"`
}

// V5GetCollateralDataResponse :
type V5GetCollateralDataResponse struct {
	CommonV5Response `json:",inline"`
	Result           V5GetCollateralDataResult `json:"result"`
}

// V5GetCollateralDataResult :
type V5GetCollateralDataResult struct {
	VipCoinList []V5GetCollateralVipCoinListItem `json:"vipCoinList"`
}

// V5GetCollateralVipCoinListItem :
type V5GetCollateralVipCoinListItem struct {
	List     []V5GetCollateralVipCoinListInnerItem `json:"list"`
	VipLevel string                                `json:"vipLevel"`
}

// V5GetCollateralVipCoinListInnerItem :
type V5GetCollateralVipCoinListInnerItem struct {
	Currency           string `json:"currency"`
	CollateralAccuracy int    `json:"collateralAccuracy"`
	InitialLTV         string `json:"initialLTV"`
	MarginCallLTV      string `json:"marginCallLTV"`
	MaxLimit           string `json:"maxLimit"`
	LiquidationLTV     string `json:"liquidationLTV"`
}

func (s *V5LoanService) GetCollateralData(param V5GetCollateralDataParam) (*V5GetCollateralDataResponse, error) {
	var res V5GetCollateralDataResponse

	queryString, err := query.Values(param)
	if err != nil {
		return nil, err
	}

	if err := s.client.getPublicly("/v5/crypto-loan/collateral-data", queryString, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// V5GetLoanableDataParam :
type V5GetLoanableDataParam struct {
	VipLevel string `json:"vipLevel,omitempty"`
	Currency string `json:"currency,omitempty"`
}

// V5GetLoanableDataResponse :
type V5GetLoanableDataResponse struct {
	CommonV5Response `json:",inline"`
	Result           V5GetLoanableDataResult `json:"result"`
}

// V5GetLoanableDataResult :
type V5GetLoanableDataResult struct {
	VipCoinList []V5GetLoanableVipCoinListItem `json:"vipCoinList"`
}

// V5GetLoanableVipCoinListItem :
type V5GetLoanableVipCoinListItem struct {
	List     []V5GetLoanableVipCoinListInnerItem `json:"list"`
	VipLevel string                              `json:"vipLevel"`
}

// V5GetLoanableVipCoinListInnerItem :
type V5GetLoanableVipCoinListInnerItem struct {
	Currency                   string `json:"currency"`
	BorrowingAccuracy          int    `json:"borrowingAccuracy"`
	FlexibleHourlyInterestRate string `json:"flexibleHourlyInterestRate"`
	MaxBorrowingAmount         string `json:"maxBorrowingAmount"`
	MinBorrowingAmount         string `json:"minBorrowingAmount"`
}

// V5LoanService :
type V5LoanService struct {
	client *Client
}

func (s *V5LoanService) GetLoanableData(param V5GetLoanableDataParam) (*V5GetLoanableDataResponse, error) {
	var res V5GetLoanableDataResponse

	queryString, err := query.Values(param)
	if err != nil {
		return nil, err
	}

	if err := s.client.getPublicly("/v5/crypto-loan/loanable-data", queryString, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
