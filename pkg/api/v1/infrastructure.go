package v1

// Infrastructure represents an infra-structure
type Infrastructure struct {
	Metadata Metadata           `json:"metadata"`
	Spec     InfrastructureSpec `json:"spec"`
}

// InfrastructureSpec represents the infra-structure specifications
type InfrastructureSpec struct {
	// Type is infra-structure  type (k8s, VM, etc.)
	Type string `json:"type" validate:"required,eq=k8s"`

	// Properties contains infra-structure properties
	Properties map[string]string `json:"properties" validate:"required"`
}
