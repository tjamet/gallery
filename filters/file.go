package filters

type InfoProvider interface {
	IsDir() bool
}

// IsFile returns true if the file provided as parameter is
// a file, false otherwise
func IsFile(info InfoProvider) bool {
	return !info.IsDir()
}
