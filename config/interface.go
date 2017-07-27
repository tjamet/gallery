package config


type Getter interface {
	Get(name string) (string, error)
}