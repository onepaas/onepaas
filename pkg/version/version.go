package version

import (
	"bytes"
	"reflect"
	"runtime"
	"text/template"
)

var (
	// AppVersion represents OnePaaS version
	AppVersion string
	// GitCommitHash represents OnePaaS commit hash
	GitCommitHash string
	// BuildTime represents OnePaaS build time
	BuildTime string
)

type version struct {
	AppVersion    string `default:"unknown-app-version"`
	GoVersion     string
	GoOs          string
	GoArch        string
	GitCommitHash string `default:"unknown-git-commit-hash"`
	BuildTime     string `default:"unknown-git-commit-hash"`
}

var versionTemplate = `Version:           {{.AppVersion}}
Go version:        {{.GoVersion}}
Git commit:        {{.GitCommitHash}}
Built:             {{.BuildTime}}
OS/Arch:           {{.GoOs}}/{{.GoArch}}`

// NewVersion creates the version instance
func NewVersion(appVersion string, gitCommitHash string, buildTime string) Version {
	v := version{
		GoVersion: runtime.Version(),
		GoOs:      runtime.GOOS,
		GoArch:    runtime.GOARCH,
	}
	vType := reflect.TypeOf(v)

	if appVersion == "" {
		f, _ := vType.FieldByName("AppVersion")
		v.AppVersion = f.Tag.Get("default")
	}

	if gitCommitHash == "" {
		f, _ := vType.FieldByName("GitCommitHash")
		v.GitCommitHash = f.Tag.Get("default")
	}

	if buildTime == "" {
		f, _ := vType.FieldByName("BuildTime")
		v.BuildTime = f.Tag.Get("default")
	}

	return &v
}

// Render renders version template
func (v *version) Render(tmpl ...string) (string, error) {
	var tmplBytes bytes.Buffer

	if len(tmpl) > 0 {
		versionTemplate = tmpl[0]
	}

	t := template.Must(template.New("version").Parse(versionTemplate))
	err := t.Execute(&tmplBytes, v)
	if err != nil {
		return "", err
	}

	return tmplBytes.String(), nil
}
