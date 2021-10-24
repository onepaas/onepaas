package problem

import (
	"encoding/json"
	"net/http"
)

const (
	MediaTypeJSON = "application/problem+json"
	MediaTypeXML  = "application/problem+xml"
	DefaultType   = "about:blank"
)

// Option configures a Problem.
type Option func(s *Problem)

var MainMembers = map[string]struct{}{"type": struct{}{}, "title": struct{}{}, "status": struct{}{}, "detail": struct{}{}, "instance": struct{}{}}

// Problem represents problem details
type Problem struct {
	details map[string]interface{}
}

func (p *Problem) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &p.details)
}

func (p *Problem) MarshalJSON() ([]byte, error) {
	return json.Marshal(p.details)
}

func (p *Problem) Write(writer http.ResponseWriter) (int, error) {
	if status, ok := p.details["status"]; ok {
		writer.WriteHeader(status.(int))

		if bodyAllowedForStatus(status.(int)) {
			if contentType := writer.Header().Get("Content-Type"); len(contentType) == 0 {
				writer.Header().Set("Content-Type", MediaTypeJSON)
			}

			body, err := p.MarshalJSON()
			if err != nil {
				return 0, err
			}

			return writer.Write(body)
		}
	}

	return 0, nil
}

func NewProblem(opts ...Option) *Problem {
	p := &Problem{
		details: make(map[string]interface{}),
	}

	opts = append([]Option{WithType(DefaultType)}, opts...)

	// apply the list of options to Problem
	for _, opt := range opts {
		opt(p)
	}

	return p
}

func NewStatusProblem(status int, opts ...Option) *Problem {
	opts = append([]Option{WithStatus(status), WithTitle(http.StatusText(status))}, opts...)

	return NewProblem(opts...)
}

// WithType configures problem's type
func WithType(t string) Option {
	return func(p *Problem) {
		p.details["type"] = t
	}
}

// WithTitle configures problem's title
func WithTitle(title string) Option {
	return func(p *Problem) {
		p.details["title"] = title
	}
}

// WithStatus configures problem's status
func WithStatus(status int) Option {
	return func(p *Problem) {
		p.details["status"] = status
	}
}

// WithDetail configures problem's detail
func WithDetail(detail string) Option {
	return func(p *Problem) {
		p.details["detail"] = detail
	}
}

// WithInstance configures problem's instance
func WithInstance(instance string) Option {
	return func(p *Problem) {
		p.details["instance"] = instance
	}
}

// WithExtension configures problem's extension
func WithExtension(key string, value interface{}) Option {
	return func(p *Problem) {
		if _, ok := MainMembers[key]; !ok {
			p.details[key] = value
		}
	}
}

// bodyAllowedForStatus is a copy of http.bodyAllowedForStatus non-exported function.
func bodyAllowedForStatus(status int) bool {
	switch {
	case status >= 100 && status <= 199:
		return false
	case status == http.StatusNoContent:
		return false
	case status == http.StatusNotModified:
		return false
	}

	return true
}
