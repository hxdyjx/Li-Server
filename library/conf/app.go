package conf

type App struct {
	Web WebConfig `toml:"web"`
}

type WebConfig struct {
	Port    int    `toml:"port"`
	Address string `toml:"address"`
}
