package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/spf13/viper"
)

type GaodeMapsClient struct {
	APIKey  string
	BaseURL string
}

func NewGaodeMapsClient() *GaodeMapsClient {
	return &GaodeMapsClient{
		APIKey:  viper.GetString("gaode_maps.api_key"),
		BaseURL: viper.GetString("gaode_maps.base_url"),
	}
}

type geocodeResponse struct {
	Status   string `json:"status"`
	Geocodes []struct {
		Location string `json:"location"`
	} `json:"geocodes"`
}

type regeocodeResponse struct {
	Status    string `json:"status"`
	Regeocode struct {
		FormattedAddress string `json:"formatted_address"`
	} `json:"regeocode"`
}

func (c *GaodeMapsClient) GetCoordinatesFromAddress(address string) (float64, float64, error) {
	u, _ := url.Parse(c.BaseURL + "/v3/geocode/geo")
	q := u.Query()
	q.Set("key", c.APIKey)
	q.Set("address", address)
	u.RawQuery = q.Encode()

	resp, err := http.Get(u.String())
	if err != nil {
		return 0, 0, err
	}
	defer resp.Body.Close()

	var res geocodeResponse
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return 0, 0, err
	}

	if res.Status != "1" || len(res.Geocodes) == 0 {
		return 0, 0, fmt.Errorf("failed to geocode address")
	}

	var lat, lng float64
	_, err = fmt.Sscanf(res.Geocodes[0].Location, "%f,%f", &lng, &lat)
	return lat, lng, err
}

func (c *GaodeMapsClient) GetAddressFromCoordinates(lat, lng float64) (string, error) {
	u, _ := url.Parse(c.BaseURL + "/v3/geocode/regeo")
	q := u.Query()
	q.Set("key", c.APIKey)
	q.Set("location", fmt.Sprintf("%f,%f", lng, lat))
	u.RawQuery = q.Encode()

	resp, err := http.Get(u.String())
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var res regeocodeResponse
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return "", err
	}

	if res.Status != "1" {
		return "", fmt.Errorf("failed to regeocode coordinates")
	}

	return res.Regeocode.FormattedAddress, nil
}
