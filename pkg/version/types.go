package version

// Version represents an object that can show your application's version
type Version interface {
	Render(tmpl ...string) (string, error)
}
