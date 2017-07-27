package config

type Test struct {
	defaultValue string
}

func (t *Test) Get(name string) (string, error) {
	return t.defaultValue, nil
}
