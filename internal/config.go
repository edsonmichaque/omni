package internal

type Config struct {
	Address string      `json:"address,omitempty"`
	Port    int         `json:"port,omitempty"`
	Cache   CacheConfig `json:"cache,omitempty"`
}

type CacheConfig struct {
	Redis     RedisConfig     `json:"redis,omitempty"`
	Memcached MemcachedConfig `json:"memcached,omitempty"`
}

type RedisConfig struct {
	Address string `json:"address,omitempty"`
	Port    int    `json:"port,omitempty"`
}

type MemcachedConfig struct {
	Address string `json:"address,omitempty"`
	Port    int    `json:"port,omitempty"`
}
