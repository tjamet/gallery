package config


type ConfigGetter interface {
	Get(name string) string
}