package gumi

const (
	VERSION_MAJOR = 0
	VERSION_MINOR = 0
	VERSION_PATCH = 1
)

func Version() (major, minor, patch int) {
	return VERSION_MAJOR, VERSION_MINOR, VERSION_PATCH
}