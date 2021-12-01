package config

var Cfg config

type config struct {
	ENV string `env:"ENV"`
}
