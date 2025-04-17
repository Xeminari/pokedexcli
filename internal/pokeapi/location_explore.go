package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
)

// List Pokemon per Location
func (c *Client) ExploreLocation(name string) (RespExploreLocation, error) {
	url := fmt.Sprintf("%s/location-area/%s", baseURL, name)

	if data, ok := c.cache.Get(url); ok {
		locationResp := RespExploreLocation{}
		err := json.Unmarshal(data, &locationResp)
		return locationResp, err
	}
	resp, err := c.httpClient.Get(url)
	if err != nil {
		return RespExploreLocation{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespExploreLocation{}, err
	}

	locationResp := RespExploreLocation{}
	err = json.Unmarshal(dat, &locationResp)
	if err != nil {
		return RespExploreLocation{}, err
	}
	c.cache.Add(url, dat)
	return locationResp, nil

}
