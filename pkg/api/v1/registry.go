package v1

// Registry represents the Docker registry
type Registry struct {
	Metadata Metadata     `json:"metadata"`
	Spec     RegistrySpec `json:"spec"`
}

// RegistrySpec represents the registry specifications
type RegistrySpec struct {
	// URL is registry address
	URL string `json:"url" validate:"required,url"`

	// Username is registry username
	Username string `json:"username" validate:"required"`

	// Secret is registry secret
	Secret string `json:"secret" validate:"required"`
}

// RegistryList represents the list of registries
type RegistryList struct {
	// Items is the list of registries
	Items []Registry `json:"items"`
}
