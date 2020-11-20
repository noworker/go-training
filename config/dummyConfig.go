package config

func NewDummyConfig() Config {
	conf := Config{App: AppConfig{KeyPath: "/home/ryo/matsuri/go-training/"}}
	return conf
}
