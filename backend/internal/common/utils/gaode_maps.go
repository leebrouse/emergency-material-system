package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/spf13/viper"
)

// GaodeMapsClient 高德地图客户端
type GaodeMapsClient struct {
	APIKey  string
	BaseURL string
}

// NewGaodeMapsClient 创建高德地图客户端
func NewGaodeMapsClient() *GaodeMapsClient {
	return &GaodeMapsClient{
		APIKey:  viper.GetString("gaode_maps.api_key"),
		BaseURL: viper.GetString("gaode_maps.base_url"),
	}
}

// GaodeGeocodingResponse 高德地理编码响应
type GaodeGeocodingResponse struct {
	Status   string `json:"status"`
	Info     string `json:"info"`
	Geocodes []struct {
		Location string `json:"location"` // "lng,lat"
	} `json:"geocodes"`
}

// GaodeReverseGeocodingResponse 高德逆地理编码响应
type GaodeReverseGeocodingResponse struct {
	Status    string `json:"status"`
	Info      string `json:"info"`
	Regeocode struct {
		FormattedAddress string `json:"formatted_address"`
	} `json:"regeocode"`
}

// GetCoordinatesFromAddress 根据地址获取经纬度 (地理编码)
func (c *GaodeMapsClient) GetCoordinatesFromAddress(address string) (lat, lng float64, err error) {
	if c.APIKey == "" {
		return 0, 0, fmt.Errorf("gaode maps API Key is not configured")
	}

	apiURL := fmt.Sprintf("%s/v3/geocode/geo?address=%s&key=%s",
		c.BaseURL, url.QueryEscape(address), c.APIKey)

	resp, err := http.Get(apiURL)
	if err != nil {
		return 0, 0, err
	}
	defer resp.Body.Close()

	var result GaodeGeocodingResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return 0, 0, err
	}

	if result.Status != "1" {
		return 0, 0, fmt.Errorf("gaode api error: %s", result.Info)
	}

	if len(result.Geocodes) == 0 {
		return 0, 0, fmt.Errorf("no results found for address: %s", address)
	}

	// 解析 "lng,lat"
	locParts := strings.Split(result.Geocodes[0].Location, ",")
	if len(locParts) != 2 {
		return 0, 0, fmt.Errorf("invalid location format from gaode api")
	}

	fmt.Sscanf(locParts[0], "%f", &lng)
	fmt.Sscanf(locParts[1], "%f", &lat)

	return lat, lng, nil
}

// GetAddressFromCoordinates 根据经纬度获取地址 (逆地理编码)
func (c *GaodeMapsClient) GetAddressFromCoordinates(lat, lng float64) (string, error) {
	if c.APIKey == "" {
		return "", fmt.Errorf("gaode maps API Key is not configured")
	}

	// 高德逆地理编码坐标格式为 "lng,lat"
	apiURL := fmt.Sprintf("%s/v3/geocode/regeo?location=%f,%f&key=%s",
		c.BaseURL, lng, lat, c.APIKey)

	resp, err := http.Get(apiURL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var result GaodeReverseGeocodingResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	if result.Status != "1" {
		return "", fmt.Errorf("gaode api error: %s", result.Info)
	}

	return result.Regeocode.FormattedAddress, nil
}
