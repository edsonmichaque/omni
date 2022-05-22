package config

type Omnid struct {
	Server *struct {
		Address string `koanf:"address"`
		Port    int    `koanf:"port"`
		TLS     *struct {
			CertificatePath string `koanf:"certificate_path"`
			PrivateKeyPath  string `koanf:"private_key_path"`
		} `koanf:"tls"`
	} `koanf:"server"`

	Cache struct {
		Redis *struct {
			Address string `koanf:"address"`
			Port    int    `koanf:"port"`
		} `koanf:"redis"`

		Memcached *struct {
			Address string `koanf:"address"`
			Port    int    `koanf:"port"`
		} `koanf:"memcached"`
	} `koanf:"cache"`

	Log struct {
		Zap    interface{} `koanf:"zap"`
		Logrus interface{} `koanf:"logrus"`
	} `koanf:"log"`

	Secrets struct {
		Vault struct {
			Address string `koanf:"address"`
			Token   string `koanf:"token"`
		} `koanf:"vault"`

		Database struct {
			EncryptionKey string `koanf:"encryption_key"`
			Table         string `koanf:"table"`
		} `koanf:"aws"`
	} `koanf:"secrets"`

	Database struct {
		SQLite struct {
			Path   string `koanf:"path"`
			Memory bool   `koanf:"memory"`
		} `koanf:"sqlite"`

		Postgres struct {
			Host     string `koanf:"host"`
			Port     string `koanf:"port"`
			User     string `koanf:"user"`
			Password string `koanf:"password"`
			Name     string `koanf:"name"`
		} `koanf:"postgres"`
	} `koanf:"database"`
}
