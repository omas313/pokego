package pokeapi

import (
	"net/http"
	"time"

	"github.com/omas313/pokego/internal/pokecache"
)

const baseURL = "https://pokeapi.co/api/v2"

type Client struct {
	cach       pokecache.Cache
	httpClient http.Client
}

func NewClient(cacheDuration time.Duration) Client {
	return Client{
		cach: pokecache.NewCache(cacheDuration),
		httpClient: http.Client{
			Timeout: time.Minute,
		},
	}
}
