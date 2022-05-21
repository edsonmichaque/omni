package config

type Config struct {
	Address string `json:"address,omitempty"`
	Port    int    `json:"port,omitempty"`
	Cache   Cache  `json:"cache,omitempty"`
	Logger  Logger `json:"logger,omitempty"`
}

type Cache struct {
	Redis     Redis     `json:"redis,omitempty"`
	Memcached Memcached `json:"memcached,omitempty"`
}

type Redis struct {
	Address string `json:"address,omitempty"`
	Port    int    `json:"port,omitempty"`
}

type Memcached struct {
	Address string `json:"address,omitempty"`
	Port    int    `json:"port,omitempty"`
}

type Logger struct {
	Zap    Zap    `json:"zap,omitempty"`
	Logrus Logrus `json:"logrus,omitempty"`
}

type Zap struct{}

type Logrus struct{}
