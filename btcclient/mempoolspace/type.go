package mempoolspace

type AddressValidationInfo struct {
	IsValid      bool   `json:"isvalid"`
	Address      string `json:"address"`
	ScriptPubKey string `json:"scriptPubKey"`
	IsScript     bool   `json:"isscript"`
	IsWitness    bool   `json:"iswitness"`
}

type FeeInfo struct {
	HalfHourFee float64 `json:"halfHourFee"`
	HourFee     float64 `json:"hourFee"`
	EconomyFee  float64 `json:"economyFee"`
	MinimumFee  float64 `json:"minimumFee"`
	FastestFee  float64 `json:"fastestFee"`
}

type DifficultyAdjustment struct {
	ProgressPercent       float64 `json:"progressPercent"`
	DifficultyChange      float64 `json:"difficultyChange"`
	EstimatedRetargetDate int64   `json:"estimatedRetargetDate"`
	RemainingBlocks       int     `json:"remainingBlocks"`
	RemainingTime         int     `json:"remainingTime"`
	PreviousRetarget      float64 `json:"previousRetarget"`
	NextRetargetHeight    int     `json:"nextRetargetHeight"`
	TimeAvg               int     `json:"timeAvg"`
	AdjustedTimeAvg       int     `json:"adjustedTimeAvg"`
	TimeOffset            int     `json:"timeOffset"`
}

type Price struct {
	Time int64   `json:"time"`
	USD  float64 `json:"USD"`
	EUR  float64 `json:"EUR"`
	GBP  float64 `json:"GBP"`
	CAD  float64 `json:"CAD"`
	CHF  float64 `json:"CHF"`
	AUD  float64 `json:"AUD"`
	JPY  float64 `json:"JPY"`
}

type HistoricalPrice struct {
	Prices       []Price            `json:"prices"`
	ExchangeRate map[string]float64 `json:"exchangeRates"`
}
