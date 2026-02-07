package utils

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCoordinatesFromAddress(t *testing.T) {

	client := &GaodeMapsClient{
		APIKey:  "bff42bf37382382b61e29c13b4964ad4",
		BaseURL: "https://restapi.amap.com",
	}

	lat, lng, err := client.GetCoordinatesFromAddress("北京市朝阳区阜通东大街6号")
	fmt.Println(lat, lng)
	assert.NoError(t, err)
}

func TestGetAddressFromCoordinates(t *testing.T) {

	client := &GaodeMapsClient{
		APIKey:  "bff42bf37382382b61e29c13b4964ad4",
		BaseURL: "https://restapi.amap.com",
	}

	addr, err := client.GetAddressFromCoordinates(39.99301, 116.47351)
	fmt.Println(addr)
	assert.NoError(t, err)
}
