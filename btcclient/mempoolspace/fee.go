package mempoolspace

func (c *MempoolSpaceClient) GetRecommendFee() (FeeInfo, error) {
	var estimates FeeInfo
	err := c.sendRequest("/v1/fees/recommended", &estimates)
	return estimates, err
}
