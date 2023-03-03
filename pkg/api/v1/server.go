package v1

// Server represents a server
type Server struct {
	Metadata Metadata   `json:"metadata"`
	Spec     ServerSpec `json:"spec"`
}

// ServerSpec represents the server specifications
type ServerSpec struct {
	// Type is server type (k8s, VM, etc.)
	Type string `json:"type" validate:"required,eq=k8s"`

	// Properties contains server properties
	Properties map[string]string `json:"properties" validate:"required"`
}
