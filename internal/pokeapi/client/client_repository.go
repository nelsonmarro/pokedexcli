package client

import (
	"fmt"
	"io"
	"net/http"
)

func (c *Client) Get(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error while getting the data: %w", err)
	}

	response, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error while getting the data: %w", err)
	}
	defer response.Body.Close()

	// Leer el cuerpo de la respuesta
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("error while getting the data: %w", err)
	}
	if response.StatusCode > 299 {
		return nil, fmt.Errorf(
			"response failed with status code: %d and\nbody: %s",
			response.StatusCode,
			body,
		)
	}

	c.cache.Add(url, body)
	return body, nil
}

func (c *Client) TryGetFromCache(
	url string,
) ([]byte, bool) {
	data, ok := c.cache.Get(url)
	if !ok {
		return nil, false
	}
	return data, true
}
