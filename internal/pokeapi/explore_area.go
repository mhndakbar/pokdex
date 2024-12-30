package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ExploreArea(areaName string) (RespAreaInformation, error) {
	url := baseUrl + "location-area" + "/" + areaName

	if val, ok := c.cache.Get(url); ok {
		areaInfoResp := RespAreaInformation{}
		err := json.Unmarshal(val, &areaInfoResp)
		if err != nil {
			return RespAreaInformation{}, err
		}

		return areaInfoResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespAreaInformation{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespAreaInformation{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespAreaInformation{}, err
	}

	areaInfoResp := RespAreaInformation{}
	err = json.Unmarshal(data, &areaInfoResp)
	if err != nil {
		return RespAreaInformation{}, err
	}

	c.cache.Add(url, data)
	return areaInfoResp, nil
}
