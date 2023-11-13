package services

import (
	"fmt"

	"github.com/go-redis/cache"
	"github.com/go-redis/redis"
	"github.com/vmihailenco/msgpack"
)

// Cache variable cache
var Cache *cache.Codec

// OpenRedisConnection Open Redis Conection
func OpenRedisConnection(redisHost string, redisPort string) {

	connectionString := fmt.Sprintf("%s:%s", redisHost, redisPort)

	ring := redis.NewRing(&redis.RingOptions{
		Addrs: map[string]string{
			"server1": connectionString,
		},
	})

	Cache = &cache.Codec{
		Redis: ring,

		Marshal: func(v interface{}) ([]byte, error) {
			return msgpack.Marshal(v)
		},
		Unmarshal: func(b []byte, v interface{}) error {
			return msgpack.Unmarshal(b, v)
		},
	}
}
