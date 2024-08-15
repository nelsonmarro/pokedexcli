package client

import (
	"net/http"
	"time"

	"github.com/nelsonmarro/pokedexcli/internal/pokecache"
)

var Base_url = "https://pokeapi.co/api/v2/"

// Client -
type Client struct {
	cache *pokecache.Cache
	http  http.Client
}

// NewClient -
func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(cacheInterval),
		http: http.Client{
			Timeout: timeout,
		},
	}
}
