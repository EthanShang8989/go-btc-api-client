package mempoolspace

import "fmt"

func (c *MempoolSpaceClient) GetAddressValidation(address string) (AddressValidationInfo, error) {
	var info AddressValidationInfo
	err := c.sendRequest(fmt.Sprintf("/v1/validate-address/%s", address), &info)
	return info, err
}
