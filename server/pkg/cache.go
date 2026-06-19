package pkg

import "github.com/dgraph-io/ristretto"

const (
	MAX_REFERENCES = 10
)

func NewCache() *ristretto.Cache {
	cache, err := ristretto.NewCache(&ristretto.Config{
		NumCounters: 1e7,
		MaxCost:     1 << 30,
		BufferItems: 64,
	})

	if err != nil {
		panic(err)
	}

	return cache
}
