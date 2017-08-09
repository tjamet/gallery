package filters

type ErrorGetter interface {
	HasError() bool
}

func HasNoError(e ErrorGetter) bool {
	return !e.HasError()
}
